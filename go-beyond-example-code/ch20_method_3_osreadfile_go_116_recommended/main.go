package main

import (
    "fmt"
    "os"
)

func main() {
    // Modern way to read entire file (Go 1.16+)
    content, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("File content: %s\n", string(content))
}