package main

import (
    "context"
    "fmt"
)

// Define custom types for context keys
type contextKey string

const (
    UserIDKey    contextKey = "userID"
    RequestIDKey contextKey = "requestID"
    TraceIDKey   contextKey = "traceID"
)

func setUserID(ctx context.Context, userID string) context.Context {
    return context.WithValue(ctx, UserIDKey, userID)
}

func getUserID(ctx context.Context) (string, bool) {
    userID, ok := ctx.Value(UserIDKey).(string)
    return userID, ok
}

func setRequestID(ctx context.Context, requestID string) context.Context {
    return context.WithValue(ctx, RequestIDKey, requestID)
}

func getRequestID(ctx context.Context) (string, bool) {
    requestID, ok := ctx.Value(RequestIDKey).(string)
    return requestID, ok
}

func main() {
    // Create context with values
    ctx := context.Background()
    ctx = setUserID(ctx, "12345")
    ctx = setRequestID(ctx, "req-001")
    
    // Retrieve values
    if userID, ok := getUserID(ctx); ok {
        fmt.Printf("User ID: %s\n", userID)
    }
    
    if requestID, ok := getRequestID(ctx); ok {
        fmt.Printf("Request ID: %s\n", requestID)
    }
}