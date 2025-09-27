package main

import (
    "fmt"
    "sync"
    "time"
)

type CircuitState int

const (
    Closed CircuitState = iota
    Open
    HalfOpen
)

type CircuitBreaker struct {
    state        CircuitState
    failures     int
    threshold    int
    timeout      time.Duration
    lastFailTime time.Time
    mutex        sync.RWMutex
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:     Closed,
        threshold: threshold,
        timeout:   timeout,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    
    if cb.state == Open {
        if time.Since(cb.lastFailTime) > cb.timeout {
            cb.state = HalfOpen
        } else {
            return fmt.Errorf("circuit breaker is open")
        }
    }
    
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        
        if cb.failures >= cb.threshold {
            cb.state = Open
        }
        return err
    }
    
    cb.failures = 0
    cb.state = Closed
    return nil
}

func simulateAPI(id int, cb *CircuitBreaker, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 5; i++ {
        err := cb.Call(func() error {
            // Simulate API call that fails sometimes
            if i == 2 {
                return fmt.Errorf("API error")
            }
            fmt.Printf("API call %d-%d succeeded\n", id, i)
            return nil
        })
        
        if err != nil {
            fmt.Printf("API call %d-%d failed: %v\n", id, i, err)
        }
        
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    cb := NewCircuitBreaker(2, 1*time.Second)
    var wg sync.WaitGroup
    
    // Start multiple goroutines making API calls
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go simulateAPI(i, cb, &wg)
    }
    
    wg.Wait()
}