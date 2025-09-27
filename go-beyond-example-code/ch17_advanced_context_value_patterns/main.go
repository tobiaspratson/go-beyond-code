package main

import (
    "context"
    "fmt"
    "time"
)

// Request metadata structure
type RequestMetadata struct {
    UserID    string
    RequestID string
    StartTime time.Time
    TraceID   string
}

// Context key for request metadata
type requestMetadataKey struct{}

func withRequestMetadata(ctx context.Context, metadata RequestMetadata) context.Context {
    return context.WithValue(ctx, requestMetadataKey{}, metadata)
}

func getRequestMetadata(ctx context.Context) (RequestMetadata, bool) {
    metadata, ok := ctx.Value(requestMetadataKey{}).(RequestMetadata)
    return metadata, ok
}

func processRequest(ctx context.Context, data string) {
    // Get request metadata
    if metadata, ok := getRequestMetadata(ctx); ok {
        fmt.Printf("Processing request %s for user %s\n", metadata.RequestID, metadata.UserID)
        fmt.Printf("Request started at: %v\n", metadata.StartTime)
        fmt.Printf("Trace ID: %s\n", metadata.TraceID)
    }
    
    // Simulate processing
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Processed data: %s\n", data)
}

func main() {
    // Create context with request metadata
    metadata := RequestMetadata{
        UserID:    "12345",
        RequestID: "req-001",
        StartTime: time.Now(),
        TraceID:   "trace-abc-123",
    }
    
    ctx := withRequestMetadata(context.Background(), metadata)
    
    // Process request
    processRequest(ctx, "important data")
}