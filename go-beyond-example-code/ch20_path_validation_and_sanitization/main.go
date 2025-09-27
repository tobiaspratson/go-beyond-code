package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func sanitizePath(path string) string {
    // Clean the path to remove redundant separators and up-level references
    cleaned := filepath.Clean(path)
    
    // Additional sanitization for security
    // Remove any ".." components that might escape the intended directory
    components := strings.Split(cleaned, string(filepath.Separator))
    var safeComponents []string
    
    for _, component := range components {
        if component == ".." {
            // Remove the last component if it exists
            if len(safeComponents) > 0 {
                safeComponents = safeComponents[:len(safeComponents)-1]
            }
        } else if component != "." && component != "" {
            safeComponents = append(safeComponents, component)
        }
    }
    
    return filepath.Join(safeComponents...)
}

func validatePath(path string) error {
    // Check for suspicious patterns
    if strings.Contains(path, "..") {
        return fmt.Errorf("path contains parent directory reference: %s", path)
    }
    
    // Check for absolute paths in certain contexts
    if filepath.IsAbs(path) {
        // Add your security checks here
        fmt.Printf("Warning: absolute path detected: %s\n", path)
    }
    
    return nil
}

func main() {
    testPaths := []string{
        "/home/user/../user/documents/file.txt",
        "../../etc/passwd",
        "documents/file.txt",
        "/safe/path/file.txt",
        "unsafe/../../../etc/passwd",
    }
    
    for _, path := range testPaths {
        fmt.Printf("Original: %s\n", path)
        
        sanitized := sanitizePath(path)
        fmt.Printf("Sanitized: %s\n", sanitized)
        
        if err := validatePath(sanitized); err != nil {
            fmt.Printf("Validation error: %v\n", err)
        }
        
        fmt.Println()
    }
}