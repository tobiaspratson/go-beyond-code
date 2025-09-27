package main

import (
    "fmt"
    "sync"
    "time"
)

type DynamicWorkerPool struct {
    wg       sync.WaitGroup
    workers  chan struct{}
    results  chan int
    mu       sync.Mutex
    active   int
}

func NewDynamicWorkerPool(maxWorkers int) *DynamicWorkerPool {
    return &DynamicWorkerPool{
        workers: make(chan struct{}, maxWorkers),
        results: make(chan int, 100),
    }
}

func (p *DynamicWorkerPool) AddWorker(id int) {
    p.workers <- struct{}{} // Acquire worker slot
    p.wg.Add(1)
    
    go func() {
        defer func() {
            p.wg.Done()
            <-p.workers // Release worker slot
        }()
        
        p.mu.Lock()
        p.active++
        active := p.active
        p.mu.Unlock()
        
        fmt.Printf("Worker %d starting (active: %d)\n", id, active)
        
        // Simulate work
        time.Sleep(time.Duration(id) * 100 * time.Millisecond)
        
        p.results <- id * 10
        
        p.mu.Lock()
        p.active--
        p.mu.Unlock()
        
        fmt.Printf("Worker %d finished\n", id)
    }()
}

func (p *DynamicWorkerPool) Wait() {
    p.wg.Wait()
    close(p.results)
}

func main() {
    pool := NewDynamicWorkerPool(3) // Max 3 concurrent workers
    
    // Add workers dynamically
    for i := 1; i <= 10; i++ {
        pool.AddWorker(i)
        time.Sleep(50 * time.Millisecond) // Stagger worker creation
    }
    
    // Wait for all workers to complete
    go pool.Wait()
    
    // Collect results
    for result := range pool.results {
        fmt.Printf("Received result: %d\n", result)
    }
}