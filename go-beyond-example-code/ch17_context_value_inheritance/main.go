package main

import (
    "context"
    "fmt"
)

func parentFunction(ctx context.Context) {
    fmt.Printf("Parent - User ID: %v\n", ctx.Value("userID"))
    fmt.Printf("Parent - Request ID: %v\n", ctx.Value("requestID"))
    
    // Add parent-specific value
    ctx = context.WithValue(ctx, "parentData", "parent-value")
    
    childFunction(ctx)
}

func childFunction(ctx context.Context) {
    fmt.Printf("Child - User ID: %v\n", ctx.Value("userID"))
    fmt.Printf("Child - Request ID: %v\n", ctx.Value("requestID"))
    fmt.Printf("Child - Parent Data: %v\n", ctx.Value("parentData"))
    
    // Add child-specific value
    ctx = context.WithValue(ctx, "childData", "child-value")
    
    grandchildFunction(ctx)
}

func grandchildFunction(ctx context.Context) {
    fmt.Printf("Grandchild - User ID: %v\n", ctx.Value("userID"))
    fmt.Printf("Grandchild - Request ID: %v\n", ctx.Value("requestID"))
    fmt.Printf("Grandchild - Parent Data: %v\n", ctx.Value("parentData"))
    fmt.Printf("Grandchild - Child Data: %v\n", ctx.Value("childData"))
}

func main() {
    // Create base context with values
    ctx := context.WithValue(context.Background(), "userID", "12345")
    ctx = context.WithValue(ctx, "requestID", "req-001")
    
    parentFunction(ctx)
}