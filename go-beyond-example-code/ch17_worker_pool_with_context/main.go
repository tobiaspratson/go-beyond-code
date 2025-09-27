package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID   int
    Data string
}

type WorkerPool struct {
    jobQueue chan Job
    workers  int
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
    return &WorkerPool{
        jobQueue: make(chan Job, queueSize),
        workers:  workers,
    }
}

func (wp *WorkerPool) Start(ctx context.Context) {
    var wg sync.WaitGroup
    
    for i := 0; i < wp.workers; i++ {
        wg.Add(1)
        go wp.worker(ctx, i, &wg)
    }
    
    go func() {
        wg.Wait()
        close(wp.jobQueue)
    }()
}

func (wp *WorkerPool) worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for {
        select {
        case job := <-wp.jobQueue:
            fmt.Printf("Worker %d processing job %d: %s\n", id, job.ID, job.Data)
            time.Sleep(100 * time.Millisecond) // Simulate work
            fmt.Printf("Worker %d completed job %d\n", id, job.ID)
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        }
    }
}

func (wp *WorkerPool) AddJob(job Job) {
    wp.jobQueue <- job
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Create and start worker pool
    pool := NewWorkerPool(3, 10)
    pool.Start(ctx)
    
    // Add jobs
    for i := 1; i <= 10; i++ {
        job := Job{
            ID:   i,
            Data: fmt.Sprintf("Job data %d", i),
        }
        pool.AddJob(job)
    }
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
}