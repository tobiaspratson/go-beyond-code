package main

import (
    "context"
    "fmt"
    "time"
)

func apiHandler(ctx context.Context, userID string) error {
    fmt.Printf("API Handler: Processing user %s\n", userID)
    
    // Add user context
    userCtx := context.WithValue(ctx, "userID", userID)
    
    // Call business logic
    return businessLogic(userCtx)
}

func businessLogic(ctx context.Context) error {
    fmt.Printf("Business Logic: User %v\n", ctx.Value("userID"))
    
    // Add processing stage
    stageCtx := context.WithValue(ctx, "stage", "business-logic")
    
    // Call data access layer
    return dataAccess(stageCtx)
}

func dataAccess(ctx context.Context) error {
    fmt.Printf("Data Access: User %v, Stage %v\n", 
        ctx.Value("userID"), ctx.Value("stage"))
    
    // Simulate database operation
    select {
    case <-time.After(500 * time.Millisecond):
        fmt.Println("Database operation completed")
        return nil
    case <-ctx.Done():
        fmt.Printf("Database operation cancelled: %v\n", ctx.Err())
        return ctx.Err()
    }
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Start API handler
    err := apiHandler(ctx, "12345")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}