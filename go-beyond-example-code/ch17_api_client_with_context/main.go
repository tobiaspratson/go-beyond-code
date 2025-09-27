package main

import (
    "context"
    "fmt"
    "time"
)

type APIClient struct {
    baseURL string
}

func (client *APIClient) Get(ctx context.Context, endpoint string) (string, error) {
    // Simulate API call
    select {
    case <-time.After(200 * time.Millisecond):
        return fmt.Sprintf("Response from %s%s", client.baseURL, endpoint), nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func (client *APIClient) Post(ctx context.Context, endpoint string, data string) error {
    // Simulate API call
    select {
    case <-time.After(300 * time.Millisecond):
        fmt.Printf("Posted to %s%s: %s\n", client.baseURL, endpoint, data)
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    client := &APIClient{baseURL: "https://api.example.com"}
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Make API calls
    response, err := client.Get(ctx, "/users")
    if err != nil {
        fmt.Printf("GET error: %v\n", err)
        return
    }
    fmt.Printf("GET response: %s\n", response)
    
    err = client.Post(ctx, "/users", "New user")
    if err != nil {
        fmt.Printf("POST error: %v\n", err)
        return
    }
}