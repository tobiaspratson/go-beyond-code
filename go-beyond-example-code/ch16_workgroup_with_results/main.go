package main

import (
    "fmt"
    "sync"
    "time"
)

// Workgroup with result collection
type ResultWorkgroup struct {
    wg      sync.WaitGroup
    results chan string
    mu      sync.Mutex
}

func NewResultWorkgroup() *ResultWorkgroup {
    return &ResultWorkgroup{
        results: make(chan string, 10),
    }
}

func (rwg *ResultWorkgroup) AddWorker(fn func() string) {
    rwg.wg.Add(1)
    go func() {
        defer rwg.wg.Done()
        result := fn()
        rwg.results <- result
    }()
}

func (rwg *ResultWorkgroup) Wait() []string {
    go func() {
        rwg.wg.Wait()
        close(rwg.results)
    }()
    
    var results []string
    for result := range rwg.results {
        results = append(results, result)
    }
    
    return results
}

func main() {
    // Create a result workgroup
    rwg := NewResultWorkgroup()
    
    // Add workers with different tasks
    rwg.AddWorker(func() string {
        time.Sleep(1 * time.Second)
        return "Task 1 completed"
    })
    
    rwg.AddWorker(func() string {
        time.Sleep(2 * time.Second)
        return "Task 2 completed"
    })
    
    rwg.AddWorker(func() string {
        time.Sleep(3 * time.Second)
        return "Task 3 completed"
    })
    
    // Wait for all workers and collect results
    results := rwg.Wait()
    
    fmt.Println("All workers completed!")
    for i, result := range results {
        fmt.Printf("Result %d: %s\n", i+1, result)
    }
}