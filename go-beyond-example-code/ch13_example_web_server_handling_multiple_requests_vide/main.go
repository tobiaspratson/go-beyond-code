package main

import (
    "fmt"
    "time"
)

// Sequential processing
func processSequential() {
    start := time.Now()
    
    // Task 1: Simulate API call
    time.Sleep(1 * time.Second)
    fmt.Println("API call completed")
    
    // Task 2: Simulate database query
    time.Sleep(1 * time.Second)
    fmt.Println("Database query completed")
    
    // Task 3: Simulate file processing
    time.Sleep(1 * time.Second)
    fmt.Println("File processing completed")
    
    elapsed := time.Since(start)
    fmt.Printf("Sequential processing took: %v\n", elapsed)
}

// Concurrent processing
func processConcurrent() {
    start := time.Now()
    
    // Channel to receive completion signals
    done := make(chan bool)
    
    // Start all tasks concurrently
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("API call completed")
        done <- true
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("Database query completed")
        done <- true
    }()
    
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("File processing completed")
        done <- true
    }()
    
    // Wait for all tasks to complete
    for i := 0; i < 3; i++ {
        <-done
    }
    
    elapsed := time.Since(start)
    fmt.Printf("Concurrent processing took: %v\n", elapsed)
}

func main() {
    fmt.Println("=== Sequential Processing ===")
    processSequential()
    
    fmt.Println("\n=== Concurrent Processing ===")
    processConcurrent()
}