package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Get context from request
    ctx := r.Context()
    
    // Add custom values to context
    ctx = context.WithValue(ctx, "userID", "12345")
    ctx = context.WithValue(ctx, "requestID", "req-001")
    
    // Process request with context
    err := processRequest(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Write([]byte("Request processed successfully"))
}

func processRequest(ctx context.Context) error {
    // Simulate work
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Request processed")
        return nil
    case <-ctx.Done():
        fmt.Printf("Request cancelled: %v\n", ctx.Err())
        return ctx.Err()
    }
}

func main() {
    http.HandleFunc("/", handler)
    
    // Start server
    go func() {
        http.ListenAndServe(":8080", nil)
    }()
    
    // Simulate client request
    time.Sleep(100 * time.Millisecond)
    
    // Make request
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    fmt.Printf("Response status: %s\n", resp.Status)
}