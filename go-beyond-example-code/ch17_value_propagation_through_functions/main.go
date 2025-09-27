package main

import (
    "context"
    "fmt"
)

func processUser(ctx context.Context, userID string) {
    fmt.Printf("Processing user: %s\n", userID)
    
    // Add processing metadata
    ctx = context.WithValue(ctx, "processingStage", "user-validation")
    
    validateUser(ctx)
}

func validateUser(ctx context.Context) {
    fmt.Printf("Validating user in stage: %v\n", ctx.Value("processingStage"))
    
    // Add validation metadata
    ctx = context.WithValue(ctx, "validationResult", "passed")
    
    authorizeUser(ctx)
}

func authorizeUser(ctx context.Context) {
    fmt.Printf("Authorizing user with result: %v\n", ctx.Value("validationResult"))
    
    // Final processing
    fmt.Printf("User processing completed in stage: %v\n", ctx.Value("processingStage"))
}

func main() {
    // Create context with user ID
    ctx := context.WithValue(context.Background(), "userID", "12345")
    
    processUser(ctx, "12345")
}