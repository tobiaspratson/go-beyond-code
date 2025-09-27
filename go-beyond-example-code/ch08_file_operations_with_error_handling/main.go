package main

import (
    "fmt"
    "os"
)

func readFile(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file %s: %w", filename, err)
    }
    return string(data), nil
}

func writeFile(filename, content string) error {
    err := os.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("failed to write file %s: %w", filename, err)
    }
    return nil
}

func main() {
    // Try to read a file
    content, err := readFile("test.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        
        // Create the file if it doesn't exist
        err = writeFile("test.txt", "Hello, World!")
        if err != nil {
            fmt.Printf("Error creating file: %v\n", err)
            return
        }
        fmt.Println("File created successfully")
        
        // Try reading again
        content, err = readFile("test.txt")
        if err != nil {
            fmt.Printf("Error reading file again: %v\n", err)
            return
        }
    }
    
    fmt.Printf("File content: %s\n", content)
}