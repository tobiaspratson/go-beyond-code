package main

import (
    "context"
    "fmt"
    "time"
)

func parentOperation(ctx context.Context) {
    fmt.Println("Starting parent operation")
    
    // Create child context with its own timeout
    childCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
    defer cancel()
    
    childOperation(childCtx)
    
    // Check if parent is still valid
    select {
    case <-ctx.Done():
        fmt.Printf("Parent operation cancelled: %v\n", ctx.Err())
    default:
        fmt.Println("Parent operation completed")
    }
}

func childOperation(ctx context.Context) {
    fmt.Println("Starting child operation")
    
    for i := 0; i < 5; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Child operation cancelled: %v\n", ctx.Err())
            return
        default:
            fmt.Printf("Child step %d\n", i+1)
            time.Sleep(300 * time.Millisecond)
        }
    }
    fmt.Println("Child operation completed")
}

func main() {
    // Parent context with 3-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    parentOperation(ctx)
    
    time.Sleep(4 * time.Second)
}