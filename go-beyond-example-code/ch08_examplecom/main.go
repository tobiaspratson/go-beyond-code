package main

import (
    "errors"
    "fmt"
    "os"
)

func readUserData(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read user data: %w", err)
    }
    return string(data), nil
}

func processUserData(filename string) error {
    data, err := readUserData(filename)
    if err != nil {
        return fmt.Errorf("failed to process user data: %w", err)
    }
    
    // Simulate processing
    if len(data) == 0 {
        return errors.New("user data is empty")
    }
    
    fmt.Printf("Processed data: %s\n", data)
    return nil
}

func main() {
    err := processUserData("user.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        
        // Check if it's a specific error type
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("The file doesn't exist")
        }
        
        // Unwrap to get the original error
        unwrapped := errors.Unwrap(err)
        if unwrapped != nil {
            fmt.Printf("Original error: %v\n", unwrapped)
        }
    }
}