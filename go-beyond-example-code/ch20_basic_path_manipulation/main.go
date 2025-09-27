package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    // Example file path
    filePath := "/home/user/documents/example.txt"
    
    // Get directory part
    dir := filepath.Dir(filePath)
    fmt.Printf("Directory: %s\n", dir)
    
    // Get filename part
    filename := filepath.Base(filePath)
    fmt.Printf("Filename: %s\n", filename)
    
    // Get file extension
    ext := filepath.Ext(filePath)
    fmt.Printf("Extension: %s\n", ext)
    
    // Get filename without extension
    name := strings.TrimSuffix(filename, ext)
    fmt.Printf("Name without extension: %s\n", name)
    
    // Join paths (cross-platform)
    newPath := filepath.Join(dir, "newfile.txt")
    fmt.Printf("New path: %s\n", newPath)
    
    // Check if path is absolute
    fmt.Printf("Is absolute: %t\n", filepath.IsAbs(filePath))
    
    // Get relative path
    relPath, err := filepath.Rel("/home/user", filePath)
    if err != nil {
        fmt.Printf("Error getting relative path: %v\n", err)
    } else {
        fmt.Printf("Relative path: %s\n", relPath)
    }
}