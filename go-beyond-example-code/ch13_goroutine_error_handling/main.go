package main

import (
    "fmt"
    "time"
)

func riskyTask(id int, results chan<- string, errors chan<- error) {
    defer func() {
        if r := recover(); r != nil {
            errors <- fmt.Errorf("goroutine %d panicked: %v", id, r)
        }
    }()
    
    // Simulate work that might fail
    if id == 3 {
        panic("Simulated panic in worker 3")
    }
    
    time.Sleep(100 * time.Millisecond)
    results <- fmt.Sprintf("Task %d completed", id)
}

func main() {
    results := make(chan string, 5)
    errors := make(chan error, 5)
    
    // Start workers
    for i := 1; i <= 5; i++ {
        go riskyTask(i, results, errors)
    }
    
    // Collect results and errors
    for i := 0; i < 5; i++ {
        select {
        case result := <-results:
            fmt.Printf("Success: %s\n", result)
        case err := <-errors:
            fmt.Printf("Error: %s\n", err)
        }
    }
}