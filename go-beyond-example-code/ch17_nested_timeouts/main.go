package main

import (
    "context"
    "fmt"
    "time"
)

func outerOperation(ctx context.Context) {
    fmt.Println("Starting outer operation")
    
    // Create a shorter timeout for inner operation
    innerCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
    defer cancel()
    
    innerOperation(innerCtx)
    
    // Check if outer context is still valid
    select {
    case <-ctx.Done():
        fmt.Printf("Outer operation cancelled: %v\n", ctx.Err())
        return
    default:
        fmt.Println("Outer operation completed")
    }
}

func innerOperation(ctx context.Context) {
    fmt.Println("Starting inner operation")
    
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Inner operation completed")
    case <-ctx.Done():
        fmt.Printf("Inner operation cancelled: %v\n", ctx.Err())
    }
}

func main() {
    // Outer timeout: 2 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    outerOperation(ctx)
    
    time.Sleep(3 * time.Second)
}