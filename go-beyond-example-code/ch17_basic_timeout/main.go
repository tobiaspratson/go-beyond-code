package main

import (
    "context"
    "fmt"
    "time"
)

func longRunningTask(ctx context.Context, name string) {
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Task %s cancelled: %v\n", name, ctx.Err())
            return
        default:
            fmt.Printf("Task %s: step %d\n", name, i+1)
            time.Sleep(500 * time.Millisecond)
        }
    }
    fmt.Printf("Task %s completed\n", name)
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Always call cancel to free resources
    
    // Start long running task
    go longRunningTask(ctx, "Worker1")
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}