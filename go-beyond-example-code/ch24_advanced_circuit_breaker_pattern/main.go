package main

import (
    "fmt"
    "sync"
    "time"
)

// Circuit breaker states
type CircuitState int

const (
    StateClosed CircuitState = iota
    StateOpen
    StateHalfOpen
)

func (cs CircuitState) String() string {
    switch cs {
    case StateClosed:
        return "CLOSED"
    case StateOpen:
        return "OPEN"
    case StateHalfOpen:
        return "HALF_OPEN"
    default:
        return "UNKNOWN"
    }
}

// Enhanced circuit breaker with metrics
type CircuitBreaker struct {
    maxFailures    int
    timeout        time.Duration
    state          CircuitState
    failures       int
    successes      int
    lastFailure    time.Time
    lastSuccess    time.Time
    mu             sync.RWMutex
    callCount      int64
    errorCount     int64
    successCount   int64
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        maxFailures: maxFailures,
        timeout:    timeout,
        state:      StateClosed,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    cb.callCount++
    
    // Check if circuit is open
    if cb.state == StateOpen {
        if time.Since(cb.lastFailure) < cb.timeout {
            cb.errorCount++
            return fmt.Errorf("circuit breaker is open")
        }
        cb.state = StateHalfOpen
    }
    
    // Execute function
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.errorCount++
        cb.lastFailure = time.Now()
        
        if cb.failures >= cb.maxFailures {
            cb.state = StateOpen
        }
        return err
    }
    
    // Success - reset failures and update state
    cb.successes++
    cb.successCount++
    cb.lastSuccess = time.Now()
    cb.failures = 0
    
    if cb.state == StateHalfOpen {
        cb.state = StateClosed
    }
    
    return nil
}

func (cb *CircuitBreaker) State() CircuitState {
    cb.mu.RLock()
    defer cb.mu.RUnlock()
    return cb.state
}

func (cb *CircuitBreaker) Stats() map[string]interface{} {
    cb.mu.RLock()
    defer cb.mu.RUnlock()
    
    return map[string]interface{}{
        "state":        cb.state.String(),
        "failures":     cb.failures,
        "successes":    cb.successes,
        "call_count":   cb.callCount,
        "error_count":  cb.errorCount,
        "success_count": cb.successCount,
        "last_failure": cb.lastFailure,
        "last_success": cb.lastSuccess,
    }
}

func (cb *CircuitBreaker) Reset() {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    cb.state = StateClosed
    cb.failures = 0
    cb.successes = 0
    cb.callCount = 0
    cb.errorCount = 0
    cb.successCount = 0
}

func main() {
    // Create circuit breaker
    cb := NewCircuitBreaker(3, 5*time.Second)
    
    // Simulate failing service
    callCount := 0
    service := func() error {
        callCount++
        if callCount <= 5 {
            return fmt.Errorf("service error")
        }
        return nil
    }
    
    fmt.Println("=== Circuit Breaker Example ===")
    
    // Make calls
    for i := 0; i < 10; i++ {
        err := cb.Call(service)
        state := cb.State()
        
        fmt.Printf("Call %d: error=%v, state=%v\n", i+1, err, state)
        
        if state == StateOpen {
            fmt.Println("Circuit breaker is open, waiting...")
            time.Sleep(6 * time.Second)
        }
        
        time.Sleep(1 * time.Second)
    }
    
    // Show final statistics
    fmt.Println("\n=== Final Statistics ===")
    stats := cb.Stats()
    for key, value := range stats {
        fmt.Printf("%s: %v\n", key, value)
    }
}