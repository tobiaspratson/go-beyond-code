package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d: step %d\n", id, i+1)
            time.Sleep(200 * time.Millisecond)
        }
    }
    fmt.Printf("Worker %d completed\n", id)
}

func coordinator(ctx context.Context) error {
    fmt.Println("Coordinator: Starting workers")
    
    var wg sync.WaitGroup
    
    // Start workers with context
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }
    
    // Wait for all workers or context cancellation
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        fmt.Println("All workers completed")
        return nil
    case <-ctx.Done():
        fmt.Printf("Coordinator cancelled: %v\n", ctx.Err())
        return ctx.Err()
    }
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Start coordinator
    err := coordinator(ctx)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}