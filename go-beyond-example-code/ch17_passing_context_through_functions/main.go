package main

import (
    "context"
    "fmt"
    "time"
)

func processRequest(ctx context.Context, requestID string) error {
    fmt.Printf("Processing request %s\n", requestID)
    
    // Simulate some work
    select {
    case <-time.After(100 * time.Millisecond):
        fmt.Printf("Request %s processed successfully\n", requestID)
        return nil
    case <-ctx.Done():
        fmt.Printf("Request %s cancelled: %v\n", requestID, ctx.Err())
        return ctx.Err()
    }
}

func handleRequest(ctx context.Context, requestID string) error {
    fmt.Printf("Handling request %s\n", requestID)
    
    // Pass context to processing function
    return processRequest(ctx, requestID)
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
    defer cancel()
    
    // Handle request
    err := handleRequest(ctx, "req-001")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}