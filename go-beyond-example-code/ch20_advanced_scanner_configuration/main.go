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
    
    scanner := bufio.NewScanner(file)
    
    // Configure buffer size for large lines
    const maxCapacity = 1024 * 1024 // 1MB
    buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity)
    
    lineCount := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineCount++
        
        // Process only non-empty lines
        if len(line) > 0 {
            fmt.Printf("Line %d (%d chars): %s\n", lineCount, len(line), line)
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}