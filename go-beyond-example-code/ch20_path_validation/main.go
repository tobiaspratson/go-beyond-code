package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func validateFilePath(path string) error {
    // Check for path traversal attacks
    if strings.Contains(path, "..") {
        return fmt.Errorf("path contains parent directory reference: %s", path)
    }
    
    // Clean the path
    cleanPath := filepath.Clean(path)
    
    // Check if path is within allowed directory
    baseDir := "/safe/directory"
    absPath, err := filepath.Abs(cleanPath)
    if err != nil {
        return fmt.Errorf("failed to get absolute path: %w", err)
    }
    
    if !strings.HasPrefix(absPath, baseDir) {
        return fmt.Errorf("path outside allowed directory: %s", absPath)
    }
    
    return nil
}

func safeOpenFile(path string) (*os.File, error) {
    if err := validateFilePath(path); err != nil {
        return nil, err
    }
    
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    
    return file, nil
}

func main() {
    // Safe file access
    file, err := safeOpenFile("safe/file.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer file.Close()
    
    fmt.Println("File opened safely")
}