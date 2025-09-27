package main

import (
    "fmt"
    "time"
)

func operationWithRetry(operation func() (string, error), maxRetries int, timeout time.Duration) (string, error) {
    for attempt := 1; attempt <= maxRetries; attempt++ {
        resultChan := make(chan string, 1)
        errorChan := make(chan error, 1)
        
        // Start the operation
        go func() {
            result, err := operation()
            if err != nil {
                errorChan <- err
            } else {
                resultChan <- result
            }
        }()
        
        // Wait for result or timeout
        select {
        case result := <-resultChan:
            return result, nil
        case err := <-errorChan:
            if attempt == maxRetries {
                return "", fmt.Errorf("operation failed after %d attempts: %v", maxRetries, err)
            }
            fmt.Printf("Attempt %d failed: %v, retrying...\n", attempt, err)
            time.Sleep(time.Duration(attempt) * 100 * time.Millisecond) // Exponential backoff
        case <-time.After(timeout):
            if attempt == maxRetries {
                return "", fmt.Errorf("operation timed out after %d attempts", maxRetries)
            }
            fmt.Printf("Attempt %d timed out, retrying...\n", attempt)
            time.Sleep(time.Duration(attempt) * 100 * time.Millisecond)
        }
    }
    
    return "", fmt.Errorf("unexpected error")
}

func unreliableOperation() (string, error) {
    // Simulate an operation that sometimes fails
    if time.Now().UnixNano()%3 == 0 {
        return "Operation succeeded", nil
    }
    return "", fmt.Errorf("operation failed")
}

func main() {
    result, err := operationWithRetry(unreliableOperation, 3, 500*time.Millisecond)
    if err != nil {
        fmt.Printf("Final error: %v\n", err)
    } else {
        fmt.Printf("Final result: %s\n", result)
    }
}