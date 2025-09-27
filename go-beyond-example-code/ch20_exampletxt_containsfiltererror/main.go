package main

import (
    "bufio"
    "fmt"
    "os"
)

func processLargeFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    // Process in batches to avoid memory issues
    batchSize := 1000
    batch := make([]string, 0, batchSize)
    
    for scanner.Scan() {
        line := scanner.Text()
        batch = append(batch, line)
        
        if len(batch) >= batchSize {
            // Process batch
            processBatch(batch)
            batch = batch[:0] // Reset slice but keep capacity
        }
    }
    
    // Process remaining lines
    if len(batch) > 0 {
        processBatch(batch)
    }
    
    return scanner.Err()
}

func processBatch(lines []string) {
    fmt.Printf("Processing batch of %d lines\n", len(lines))
    // Process the batch of lines here
    // This approach keeps memory usage constant
}

func main() {
    err := processLargeFile("large_file.txt")
    if err != nil {
        fmt.Printf("Error processing file: %v\n", err)
    }
}