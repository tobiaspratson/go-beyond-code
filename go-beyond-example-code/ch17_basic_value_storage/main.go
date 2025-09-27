package main

import (
    "context"
    "fmt"
)

func main() {
    // Create context with values
    ctx := context.WithValue(context.Background(), "userID", "12345")
    ctx = context.WithValue(ctx, "requestID", "req-001")
    ctx = context.WithValue(ctx, "traceID", "trace-abc")
    
    // Retrieve values
    if userID := ctx.Value("userID"); userID != nil {
        fmt.Printf("User ID: %v\n", userID)
    }
    
    if requestID := ctx.Value("requestID"); requestID != nil {
        fmt.Printf("Request ID: %v\n", requestID)
    }
    
    if traceID := ctx.Value("traceID"); traceID != nil {
        fmt.Printf("Trace ID: %v\n", traceID)
    }
    
    // Non-existent key
    if value := ctx.Value("nonExistent"); value != nil {
        fmt.Printf("Non-existent: %v\n", value)
    } else {
        fmt.Println("Non-existent key not found")
    }
}