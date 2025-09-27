package main

import (
    "bufio"
    "fmt"
    "os"
)

func processLargeFileEfficiently(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
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
        
        // Process line without storing in memory
        processLine(line)
        
        // Optional: process in batches
        if lineCount%1000 == 0 {
            fmt.Printf("Processed %d lines\n", lineCount)
        }
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    
    fmt.Printf("Total lines processed: %d\n", lineCount)
    return nil
}

func processLine(line string) {
    // Process individual line
    // This keeps memory usage constant
}

func main() {
    err := processLargeFileEfficiently("large_file.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}