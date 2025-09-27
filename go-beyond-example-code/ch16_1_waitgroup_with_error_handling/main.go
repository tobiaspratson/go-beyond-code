package main

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

type WorkerResult struct {
    ID    int
    Value int
    Error error
}

func workerWithError(id int, wg *sync.WaitGroup, results chan<- WorkerResult) {
    defer wg.Done()
    
    // Simulate work that might fail
    time.Sleep(time.Duration(id) * 100 * time.Millisecond)
    
    var result WorkerResult
    result.ID = id
    
    // Simulate some workers failing
    if id%3 == 0 {
        result.Error = errors.New("worker failed")
    } else {
        result.Value = id * 10
    }
    
    results <- result
}

func main() {
    var wg sync.WaitGroup
    results := make(chan WorkerResult, 10)
    
    // Start workers
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go workerWithError(i, &wg, results)
    }
    
    // Close results channel when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    var successCount, errorCount int
    for result := range results {
        if result.Error != nil {
            fmt.Printf("Worker %d failed: %v\n", result.ID, result.Error)
            errorCount++
        } else {
            fmt.Printf("Worker %d succeeded: value=%d\n", result.ID, result.Value)
            successCount++
        }
    }
    
    fmt.Printf("Summary: %d succeeded, %d failed\n", successCount, errorCount)
}