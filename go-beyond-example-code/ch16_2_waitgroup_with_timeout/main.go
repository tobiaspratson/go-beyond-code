package main

import (
    "fmt"
    "sync"
    "time"
)

func workerWithTimeout(id int, wg *sync.WaitGroup, results chan<- int) {
    defer wg.Done()
    
    // Simulate work that might take a long time
    workDuration := time.Duration(id) * 200 * time.Millisecond
    time.Sleep(workDuration)
    
    results <- id * 10
}

func main() {
    var wg sync.WaitGroup
    results := make(chan int, 10)
    
    // Start workers
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go workerWithTimeout(i, &wg, results)
    }
    
    // Wait with timeout
    done := make(chan bool)
    go func() {
        wg.Wait()
        close(results)
        done <- true
    }()
    
    select {
    case <-done:
        fmt.Println("All workers completed")
        for result := range results {
            fmt.Printf("Result: %d\n", result)
        }
    case <-time.After(500 * time.Millisecond):
        fmt.Println("Timeout waiting for workers")
        // Note: Workers are still running in background
    }
}