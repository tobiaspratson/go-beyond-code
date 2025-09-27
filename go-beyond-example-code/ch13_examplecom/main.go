package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID   int
    Data string
}

func producer(jobs chan<- Job, wg *sync.WaitGroup) {
    defer wg.Done()
    defer close(jobs)
    
    for i := 1; i <= 10; i++ {
        job := Job{
            ID:   i,
            Data: fmt.Sprintf("Job data %d", i),
        }
        
        fmt.Printf("Producing job %d\n", i)
        jobs <- job
        time.Sleep(100 * time.Millisecond)
    }
}

func consumer(workerID int, jobs <-chan Job, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d: %s\n", workerID, job.ID, job.Data)
        time.Sleep(200 * time.Millisecond) // Simulate work
        fmt.Printf("Worker %d completed job %d\n", workerID, job.ID)
    }
}

func main() {
    jobs := make(chan Job, 5) // Buffered channel
    var wg sync.WaitGroup
    
    // Start producer
    wg.Add(1)
    go producer(jobs, &wg)
    
    // Start consumers
    numWorkers := 3
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go consumer(i, jobs, &wg)
    }
    
    // Wait for all goroutines to complete
    wg.Wait()
    fmt.Println("All jobs completed")
}