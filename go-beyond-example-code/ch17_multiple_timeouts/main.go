package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d: step %d\n", id, i+1)
            time.Sleep(300 * time.Millisecond)
        }
    }
    fmt.Printf("Worker %d completed\n", id)
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Start multiple workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }
    
    // Wait for timeout
    time.Sleep(2 * time.Second)
    fmt.Printf("All workers should be cancelled\n")
}