package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Duration(id) * 100 * time.Millisecond)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // Start 5 workers
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    fmt.Println("Waiting for workers to finish...")
    wg.Wait()
    fmt.Println("All workers completed!")
}