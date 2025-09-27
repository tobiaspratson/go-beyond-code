package main

import (
    "fmt"
    "os"
)

func readFileContent(filename string) (string, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file %s: %w", filename, err)
    }
    return string(content), nil
}

func main() {
    content, err := readFileContent("example.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Content: %s\n", content)
}