package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Combined workgroup and errorgroup
type WorkErrorgroup struct {
    wg      sync.WaitGroup
    errors  chan error
    results chan string
    ctx     context.Context
    cancel  context.CancelFunc
    mu      sync.Mutex
}

func NewWorkErrorgroup(ctx context.Context) *WorkErrorgroup {
    ctx, cancel := context.WithCancel(ctx)
    return &WorkErrorgroup{
        ctx:     ctx,
        cancel:  cancel,
        errors:  make(chan error, 10),
        results: make(chan string, 10),
    }
}

func (weg *WorkErrorgroup) AddWorker(fn func(context.Context) (string, error)) {
    weg.wg.Add(1)
    go func() {
        defer weg.wg.Done()
        result, err := fn(weg.ctx)
        if err != nil {
            weg.errors <- err
            weg.cancel() // Cancel context on error
        } else {
            weg.results <- result
        }
    }()
}

func (weg *WorkErrorgroup) Wait() ([]string, []error) {
    go func() {
        weg.wg.Wait()
        close(weg.errors)
        close(weg.results)
    }()
    
    var results []string
    var errors []error
    
    // Collect results and errors
    for {
        select {
        case result, ok := <-weg.results:
            if !ok {
                goto done
            }
            results = append(results, result)
        case err, ok := <-weg.errors:
            if !ok {
                goto done
            }
            errors = append(errors, err)
        }
    }
    
done:
    return results, errors
}

func (weg *WorkErrorgroup) Cancel() {
    weg.cancel()
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Create a work errorgroup
    weg := NewWorkErrorgroup(ctx)
    
    // Add workers that return results or errors
    weg.AddWorker(func(ctx context.Context) (string, error) {
        for i := 0; i < 5; i++ {
            select {
            case <-ctx.Done():
                return "", ctx.Err()
            default:
                fmt.Printf("Worker 1: step %d\n", i+1)
                time.Sleep(1 * time.Second)
            }
        }
        return "Worker 1 completed successfully", nil
    })
    
    weg.AddWorker(func(ctx context.Context) (string, error) {
        for i := 0; i < 3; i++ {
            select {
            case <-ctx.Done():
                return "", ctx.Err()
            default:
                fmt.Printf("Worker 2: step %d\n", i+1)
                time.Sleep(800 * time.Millisecond)
            }
        }
        return "Worker 2 completed successfully", nil
    })
    
    // Wait for all workers and collect results and errors
    results, errors := weg.Wait()
    
    fmt.Printf("Workers completed with %d results and %d errors:\n", len(results), len(errors))
    
    for i, result := range results {
        fmt.Printf("Result %d: %s\n", i+1, result)
    }
    
    for i, err := range errors {
        fmt.Printf("Error %d: %v\n", i+1, err)
    }
}