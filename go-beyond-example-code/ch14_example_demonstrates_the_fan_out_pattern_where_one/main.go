package main

import (
    "fmt"
    "time"
)

type CircuitState int

const (
    Closed CircuitState = iota
    Open
    HalfOpen
)

type CircuitBreaker struct {
    state         CircuitState
    failureCount  int
    maxFailures   int
    timeout       time.Duration
    lastFailTime  time.Time
    successCount  int
    resetTimeout  time.Duration
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:        Closed,
        maxFailures:  maxFailures,
        timeout:      timeout,
        resetTimeout: 5 * time.Second,
    }
}

func (cb *CircuitBreaker) Call(operation func() (string, error)) (string, error) {
    if cb.state == Open {
        if time.Since(cb.lastFailTime) > cb.resetTimeout {
            cb.state = HalfOpen
            cb.successCount = 0
        } else {
            return "", fmt.Errorf("circuit breaker is open")
        }
    }
    
    result, err := operation()
    
    if err != nil {
        cb.failureCount++
        if cb.failureCount >= cb.maxFailures {
            cb.state = Open
            cb.lastFailTime = time.Now()
        }
        return "", err
    }
    
    cb.successCount++
    if cb.state == HalfOpen && cb.successCount >= 2 {
        cb.state = Closed
        cb.failureCount = 0
    }
    
    return result, nil
}

func unreliableService() (string, error) {
    if time.Now().UnixNano()%3 == 0 {
        return "Service response", nil
    }
    return "", fmt.Errorf("service error")
}

func main() {
    cb := NewCircuitBreaker(3, 1*time.Second)
    
    for i := 0; i < 10; i++ {
        result, err := cb.Call(unreliableService)
        if err != nil {
            fmt.Printf("Attempt %d: %v\n", i+1, err)
        } else {
            fmt.Printf("Attempt %d: %s\n", i+1, result)
        }
        time.Sleep(500 * time.Millisecond)
    }
}