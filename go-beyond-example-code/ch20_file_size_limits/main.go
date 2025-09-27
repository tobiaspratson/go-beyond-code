package main

import (
    "fmt"
    "os"
)

func readFileWithSizeLimit(filename string, maxSize int64) ([]byte, error) {
    // Check file size first
    info, err := os.Stat(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to get file info: %w", err)
    }
    
    if info.Size() > maxSize {
        return nil, fmt.Errorf("file too large: %d bytes (max: %d)", info.Size(), maxSize)
    }
    
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    
    return content, nil
}

func main() {
    content, err := readFileWithSizeLimit("example.txt", 1024*1024) // 1MB limit
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("File content (%d bytes): %s\n", len(content), string(content))
}