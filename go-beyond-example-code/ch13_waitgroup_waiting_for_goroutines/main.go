package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement counter when done
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(500 * time.Millisecond) // Simulate work
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // Start 5 workers
    for i := 1; i <= 5; i++ {
        wg.Add(1) // Increment counter
        go worker(i, &wg)
    }
    
    // Wait for all workers to complete
    wg.Wait()
    fmt.Println("All workers completed")
}