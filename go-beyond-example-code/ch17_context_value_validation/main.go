package main

import (
    "context"
    "fmt"
    "strings"
)

// Validation functions for context values
func validateUserID(userID string) error {
    if userID == "" {
        return fmt.Errorf("user ID cannot be empty")
    }
    if len(userID) < 3 {
        return fmt.Errorf("user ID must be at least 3 characters")
    }
    return nil
}

func validateRequestID(requestID string) error {
    if requestID == "" {
        return fmt.Errorf("request ID cannot be empty")
    }
    if !strings.HasPrefix(requestID, "req-") {
        return fmt.Errorf("request ID must start with 'req-'")
    }
    return nil
}

func setUserIDWithValidation(ctx context.Context, userID string) (context.Context, error) {
    if err := validateUserID(userID); err != nil {
        return ctx, err
    }
    return context.WithValue(ctx, "userID", userID), nil
}

func setRequestIDWithValidation(ctx context.Context, requestID string) (context.Context, error) {
    if err := validateRequestID(requestID); err != nil {
        return ctx, err
    }
    return context.WithValue(ctx, "requestID", requestID), nil
}

func main() {
    ctx := context.Background()
    
    // Set user ID with validation
    ctx, err := setUserIDWithValidation(ctx, "12345")
    if err != nil {
        fmt.Printf("Error setting user ID: %v\n", err)
        return
    }
    
    // Set request ID with validation
    ctx, err = setRequestIDWithValidation(ctx, "req-001")
    if err != nil {
        fmt.Printf("Error setting request ID: %v\n", err)
        return
    }
    
    // Try invalid request ID
    ctx, err = setRequestIDWithValidation(ctx, "invalid-id")
    if err != nil {
        fmt.Printf("Error setting invalid request ID: %v\n", err)
    }
    
    fmt.Printf("User ID: %v\n", ctx.Value("userID"))
    fmt.Printf("Request ID: %v\n", ctx.Value("requestID"))
}