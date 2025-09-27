package main

import (
    "context"
    "fmt"
    "time"
)

func validateInput(ctx context.Context, input string) error {
    fmt.Printf("Validating input: %s\n", input)
    
    // Simulate validation
    select {
    case <-time.After(100 * time.Millisecond):
        if input == "" {
            return fmt.Errorf("input cannot be empty")
        }
        fmt.Println("Input validation passed")
        return nil
    case <-ctx.Done():
        return fmt.Errorf("validation cancelled: %v", ctx.Err())
    }
}

func processData(ctx context.Context, data string) error {
    fmt.Printf("Processing data: %s\n", data)
    
    // Simulate processing
    select {
    case <-time.After(200 * time.Millisecond):
        fmt.Println("Data processing completed")
        return nil
    case <-ctx.Done():
        return fmt.Errorf("processing cancelled: %v", ctx.Err())
    }
}

func saveData(ctx context.Context, data string) error {
    fmt.Printf("Saving data: %s\n", data)
    
    // Simulate save operation
    select {
    case <-time.After(150 * time.Millisecond):
        fmt.Println("Data saved successfully")
        return nil
    case <-ctx.Done():
        return fmt.Errorf("save cancelled: %v", ctx.Err())
    }
}

func handleData(ctx context.Context, input string) error {
    // Step 1: Validate input
    if err := validateInput(ctx, input); err != nil {
        return fmt.Errorf("validation failed: %v", err)
    }
    
    // Step 2: Process data
    if err := processData(ctx, input); err != nil {
        return fmt.Errorf("processing failed: %v", err)
    }
    
    // Step 3: Save data
    if err := saveData(ctx, input); err != nil {
        return fmt.Errorf("save failed: %v", err)
    }
    
    return nil
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Handle data with context propagation
    err := handleData(ctx, "important data")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}