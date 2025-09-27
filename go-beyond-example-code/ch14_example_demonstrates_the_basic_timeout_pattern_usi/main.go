package main

import (
    "fmt"
    "time"
)

func operationWithTimeout(operation func() string, timeout time.Duration) (string, error) {
    resultChan := make(chan string, 1)
    errorChan := make(chan error, 1)
    
    // Start the operation
    go func() {
        defer func() {
            if r := recover(); r != nil {
                errorChan <- fmt.Errorf("operation panicked: %v", r)
            }
        }()
        
        result := operation()
        resultChan <- result
    }()
    
    // Wait for result or timeout
    select {
    case result := <-resultChan:
        return result, nil
    case err := <-errorChan:
        return "", err
    case <-time.After(timeout):
        return "", fmt.Errorf("operation timed out after %v", timeout)
    }
}

func slowOperation() string {
    time.Sleep(2 * time.Second)
    return "Operation completed"
}

func fastOperation() string {
    time.Sleep(100 * time.Millisecond)
    return "Fast operation completed"
}

func main() {
    // Test with slow operation (should timeout)
    result, err := operationWithTimeout(slowOperation, 1*time.Second)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %s\n", result)
    }
    
    // Test with fast operation (should succeed)
    result, err = operationWithTimeout(fastOperation, 1*time.Second)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %s\n", result)
    }
}