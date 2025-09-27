package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Generic workgroup with type safety
type GenericWorkgroup[T any] struct {
    wg      sync.WaitGroup
    results chan T
    errors  chan error
    ctx     context.Context
    cancel  context.CancelFunc
    mu      sync.Mutex
}

func NewGenericWorkgroup[T any](ctx context.Context) *GenericWorkgroup[T] {
    ctx, cancel := context.WithCancel(ctx)
    return &GenericWorkgroup[T]{
        ctx:     ctx,
        cancel:  cancel,
        results: make(chan T, 10),
        errors:  make(chan error, 10),
    }
}

func (gw *GenericWorkgroup[T]) AddWorker(fn func(context.Context) (T, error)) {
    gw.wg.Add(1)
    go func() {
        defer gw.wg.Done()
        result, err := fn(gw.ctx)
        if err != nil {
            gw.errors <- err
            gw.cancel()  // Cancel all other workers on error
        } else {
            gw.results <- result
        }
    }()
}

func (gw *GenericWorkgroup[T]) Wait() ([]T, []error) {
    go func() {
        gw.wg.Wait()
        close(gw.results)
        close(gw.errors)
    }()
    
    var results []T
    var errors []error
    
    for {
        select {
        case result, ok := <-gw.results:
            if !ok {
                goto done
            }
            results = append(results, result)
        case err, ok := <-gw.errors:
            if !ok {
                goto done
            }
            errors = append(errors, err)
        }
    }
    
done:
    return results, errors
}

func (gw *GenericWorkgroup[T]) Cancel() {
    gw.cancel()
}

// Worker function type
type WorkerFunc[T any] func(context.Context) (T, error)

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Create generic workgroup for string results
    gw := NewGenericWorkgroup[string](ctx)
    
    // Define worker functions
    workers := []WorkerFunc[string]{
        func(ctx context.Context) (string, error) {
            time.Sleep(1 * time.Second)
            return "Task 1 completed", nil
        },
        func(ctx context.Context) (string, error) {
            time.Sleep(2 * time.Second)
            return "Task 2 completed", nil
        },
        func(ctx context.Context) (string, error) {
            time.Sleep(3 * time.Second)
            return "Task 3 completed", nil
        },
    }
    
    // Add workers
    for _, worker := range workers {
        gw.AddWorker(worker)
    }
    
    // Wait for results
    results, errors := gw.Wait()
    
    fmt.Printf("Results: %v\n", results)
    fmt.Printf("Errors: %v\n", errors)
}