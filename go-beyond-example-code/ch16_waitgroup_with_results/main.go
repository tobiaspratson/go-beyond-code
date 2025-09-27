package main

import (
    "fmt"
    "sync"
    "time"
)

type Result struct {
    WorkerID int
    Value    int
    Duration time.Duration
}

func worker(id int, wg *sync.WaitGroup, results chan<- Result) {
    defer wg.Done()
    
    start := time.Now()
    
    // Simulate work
    time.Sleep(time.Duration(id) * 100 * time.Millisecond)
    
    result := Result{
        WorkerID: id,
        Value:    id * 10,
        Duration: time.Since(start),
    }
    
    results <- result
}

func main() {
    var wg sync.WaitGroup
    results := make(chan Result, 5)
    
    // Start workers
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg, results)
    }
    
    // Close results channel when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    fmt.Println("Results:")
    for result := range results {
        fmt.Printf("Worker %d: value=%d, duration=%v\n", 
            result.WorkerID, result.Value, result.Duration)
    }
}