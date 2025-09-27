package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(200 * time.Millisecond)
        }
    }
}

func main() {
    // Create context with deadline
    deadline := time.Now().Add(2 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    
    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }
    
    // Wait for deadline
    time.Sleep(3 * time.Second)
    fmt.Printf("All workers should be cancelled by deadline\n")
}