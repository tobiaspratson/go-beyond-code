package main

import (
    "fmt"
    "runtime"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(500 * time.Millisecond) // Simulate work
        results <- job * 2
    }
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    // Create channels
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    
    // Start 3 workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Send jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Collect results
    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
    
    // Show goroutine count
    fmt.Printf("Goroutines running: %d\n", runtime.NumGoroutine())
}