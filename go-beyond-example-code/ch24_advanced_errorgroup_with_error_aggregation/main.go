package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Error types for categorization
type ErrorType int

const (
    NetworkError ErrorType = iota
    ValidationError
    TimeoutError
    UnknownError
)

type CategorizedError struct {
    Type    ErrorType
    Message string
    Err     error
}

func (ce CategorizedError) Error() string {
    return fmt.Sprintf("[%s] %s: %v", ce.Type, ce.Message, ce.Err)
}

// Advanced errorgroup
type AdvancedErrorgroup struct {
    wg       sync.WaitGroup
    errors   chan CategorizedError
    results  chan interface{}
    ctx      context.Context
    cancel   context.CancelFunc
    mu       sync.RWMutex
    errorMap map[ErrorType][]CategorizedError
}

func NewAdvancedErrorgroup(ctx context.Context) *AdvancedErrorgroup {
    ctx, cancel := context.WithCancel(ctx)
    return &AdvancedErrorgroup{
        ctx:      ctx,
        cancel:   cancel,
        errors:   make(chan CategorizedError, 10),
        results:  make(chan interface{}, 10),
        errorMap: make(map[ErrorType][]CategorizedError),
    }
}

func (aeg *AdvancedErrorgroup) AddWorker(fn func(context.Context) (interface{}, error)) {
    aeg.wg.Add(1)
    go func() {
        defer aeg.wg.Done()
        result, err := fn(aeg.ctx)
        if err != nil {
            // Categorize error based on context
            errorType := aeg.categorizeError(err)
            categorizedErr := CategorizedError{
                Type:    errorType,
                Message: "Worker error",
                Err:     err,
            }
            aeg.errors <- categorizedErr
        } else {
            aeg.results <- result
        }
    }()
}

func (aeg *AdvancedErrorgroup) categorizeError(err error) ErrorType {
    // Simple error categorization logic
    errStr := err.Error()
    if contains(errStr, "timeout") {
        return TimeoutError
    }
    if contains(errStr, "network") || contains(errStr, "connection") {
        return NetworkError
    }
    if contains(errStr, "validation") || contains(errStr, "invalid") {
        return ValidationError
    }
    return UnknownError
}

func (aeg *AdvancedErrorgroup) Wait() ([]interface{}, map[ErrorType][]CategorizedError) {
    go func() {
        aeg.wg.Wait()
        close(aeg.results)
        close(aeg.errors)
    }()
    
    var results []interface{}
    
    for {
        select {
        case result, ok := <-aeg.results:
            if !ok {
                goto done
            }
            results = append(results, result)
        case err, ok := <-aeg.errors:
            if !ok {
                goto done
            }
            aeg.mu.Lock()
            aeg.errorMap[err.Type] = append(aeg.errorMap[err.Type], err)
            aeg.mu.Unlock()
        }
    }
    
done:
    return results, aeg.errorMap
}

func (aeg *AdvancedErrorgroup) Cancel() {
    aeg.cancel()
}

func contains(s, substr string) bool {
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    aeg := NewAdvancedErrorgroup(ctx)
    
    // Add various workers with different error types
    aeg.AddWorker(func(ctx context.Context) (interface{}, error) {
        time.Sleep(1 * time.Second)
        return "Success 1", nil
    })
    
    aeg.AddWorker(func(ctx context.Context) (interface{}, error) {
        time.Sleep(500 * time.Millisecond)
        return nil, fmt.Errorf("network connection failed")
    })
    
    aeg.AddWorker(func(ctx context.Context) (interface{}, error) {
        time.Sleep(2 * time.Second)
        return nil, fmt.Errorf("validation error: invalid input")
    })
    
    aeg.AddWorker(func(ctx context.Context) (interface{}, error) {
        time.Sleep(4 * time.Second)  // This will timeout
        return "Success 2", nil
    })
    
    results, errorMap := aeg.Wait()
    
    fmt.Printf("Results: %v\n", results)
    fmt.Printf("Errors by type:\n")
    for errorType, errors := range errorMap {
        fmt.Printf("  %v: %d errors\n", errorType, len(errors))
        for _, err := range errors {
            fmt.Printf("    - %v\n", err)
        }
    }
}