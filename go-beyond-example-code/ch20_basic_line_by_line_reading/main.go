package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Create scanner with default line splitting
    scanner := bufio.NewScanner(file)
    lineNumber := 1
    
    // Read line by line
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("Line %d: %s\n", lineNumber, line)
        lineNumber++
    }
    
    // Check for scanning errors
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
    
    fmt.Printf("Total lines processed: %d\n", lineNumber-1)
}