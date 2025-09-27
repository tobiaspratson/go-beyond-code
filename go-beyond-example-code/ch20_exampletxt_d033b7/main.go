package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func processFile(filename string) error {
    // Check if file exists first
    if _, err := os.Stat(filename); err != nil {
        if os.IsNotExist(err) {
            return fmt.Errorf("file does not exist: %s", filename)
        }
        return fmt.Errorf("failed to check file status: %w", err)
    }
    
    // Open file with proper error handling
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            fmt.Printf("Warning: failed to close file %s: %v\n", filename, closeErr)
        }
    }()
    
    // Process file...
    return nil
}

func main() {
    err := processFile("nonexistent.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}