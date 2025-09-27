package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Progress information
type Progress struct {
    Completed int
    Total     int
    Percent   float64
    Message   string
}

// Progress tracking workgroup
type ProgressWorkgroup struct {
    wg        sync.WaitGroup
    results   chan interface{}
    errors    chan error
    progress  chan Progress
    ctx       context.Context
    cancel    context.CancelFunc
    mu        sync.RWMutex
    completed int
    total     int
}

func NewProgressWorkgroup(ctx context.Context, total int) *ProgressWorkgroup {
    ctx, cancel := context.WithCancel(ctx)
    return &ProgressWorkgroup{
        ctx:      ctx,
        cancel:   cancel,
        results:  make(chan interface{}, total),
        errors:   make(chan error, total),
        progress: make(chan Progress, total),
        total:    total,
    }
}

func (pwg *ProgressWorkgroup) AddWorker(fn func(context.Context) (interface{}, error)) {
    pwg.wg.Add(1)
    go func() {
        defer pwg.wg.Done()
        result, err := fn(pwg.ctx)
        
        pwg.mu.Lock()
        pwg.completed++
        progress := Progress{
            Completed: pwg.completed,
            Total:     pwg.total,
            Percent:   float64(pwg.completed) / float64(pwg.total) * 100,
            Message:   fmt.Sprintf("Completed %d/%d", pwg.completed, pwg.total),
        }
        pwg.mu.Unlock()
        
        pwg.progress <- progress
        
        if err != nil {
            pwg.errors <- err
        } else {
            pwg.results <- result
        }
    }()
}

func (pwg *ProgressWorkgroup) Wait() ([]interface{}, []error, []Progress) {
    go func() {
        pwg.wg.Wait()
        close(pwg.results)
        close(pwg.errors)
        close(pwg.progress)
    }()
    
    var results []interface{}
    var errors []error
    var progressUpdates []Progress
    
    for {
        select {
        case result, ok := <-pwg.results:
            if !ok {
                goto done
            }
            results = append(results, result)
        case err, ok := <-pwg.errors:
            if !ok {
                goto done
            }
            errors = append(errors, err)
        case progress, ok := <-pwg.progress:
            if !ok {
                goto done
            }
            progressUpdates = append(progressUpdates, progress)
            fmt.Printf("Progress: %.1f%% - %s\n", progress.Percent, progress.Message)
        }
    }
    
done:
    return results, errors, progressUpdates
}

func (pwg *ProgressWorkgroup) Cancel() {
    pwg.cancel()
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Create workgroup for 5 tasks
    pwg := NewProgressWorkgroup(ctx, 5)
    
    // Add workers with different completion times
    for i := 1; i <= 5; i++ {
        taskNum := i
        pwg.AddWorker(func(ctx context.Context) (interface{}, error) {
            duration := time.Duration(taskNum) * time.Second
            time.Sleep(duration)
            return fmt.Sprintf("Task %d completed", taskNum), nil
        })
    }
    
    results, errors, progress := pwg.Wait()
    
    fmt.Printf("\nFinal Results:\n")
    fmt.Printf("Completed: %d tasks\n", len(results))
    fmt.Printf("Errors: %d\n", len(errors))
    fmt.Printf("Progress updates: %d\n", len(progress))
    
    for _, result := range results {
        fmt.Printf("  - %v\n", result)
    }
}