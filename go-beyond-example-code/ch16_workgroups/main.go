package main

import (
    "fmt"
    "sync"
    "time"
)

// Workgroup for coordinating workers
type Workgroup struct {
    wg     sync.WaitGroup
    workers int
}

func NewWorkgroup(workers int) *Workgroup {
    return &Workgroup{workers: workers}
}

func (wg *Workgroup) AddWorker(fn func()) {
    wg.wg.Add(1)
    go func() {
        defer wg.wg.Done()
        fn()
    }()
}

func (wg *Workgroup) Wait() {
    wg.wg.Wait()
}

func (wg *Workgroup) WorkerCount() int {
    return wg.workers
}

func main() {
    // Create a workgroup
    wg := NewWorkgroup(3)
    
    // Add workers
    for i := 0; i < 3; i++ {
        workerID := i
        wg.AddWorker(func() {
            fmt.Printf("Worker %d starting\n", workerID)
            time.Sleep(time.Duration(workerID+1) * time.Second)
            fmt.Printf("Worker %d finished\n", workerID)
        })
    }
    
    fmt.Printf("Waiting for %d workers to complete...\n", wg.WorkerCount())
    wg.Wait()
    fmt.Println("All workers completed!")
}