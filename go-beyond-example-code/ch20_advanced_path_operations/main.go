package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func analyzePath(path string) {
    fmt.Printf("Analyzing path: %s\n", path)
    fmt.Printf("  Directory: %s\n", filepath.Dir(path))
    fmt.Printf("  Base name: %s\n", filepath.Base(path))
    fmt.Printf("  Extension: %s\n", filepath.Ext(path))
    fmt.Printf("  Is absolute: %t\n", filepath.IsAbs(path))
    
    // Split path components
    components := strings.Split(filepath.Clean(path), string(filepath.Separator))
    fmt.Printf("  Components: %v\n", components)
    
    // Get parent directory
    parent := filepath.Dir(path)
    if parent != path {
        fmt.Printf("  Parent directory: %s\n", parent)
    }
}

func main() {
    paths := []string{
        "/home/user/documents/file.txt",
        "relative/path/file.go",
        "C:\\Users\\Name\\Documents\\file.txt", // Windows style
        "file.txt",
        "/",
        ".",
        "..",
    }
    
    for _, path := range paths {
        analyzePath(path)
        fmt.Println()
    }
}