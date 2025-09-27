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
    jobQueue   chan Job
    workers    int
    ctx        context.Context
    cancel     context.CancelFunc
    wg         sync.WaitGroup
    isRunning  bool
    mutex      sync.RWMutex
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
    return &WorkerPool{
        jobQueue: make(chan Job, queueSize),
        workers:  workers,
    }
}

func (wp *WorkerPool) Start(ctx context.Context) {
    wp.mutex.Lock()
    defer wp.mutex.Unlock()
    
    if wp.isRunning {
        return
    }
    
    wp.ctx, wp.cancel = context.WithCancel(ctx)
    wp.isRunning = true
    
    for i := 0; i < wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) Stop() {
    wp.mutex.Lock()
    defer wp.mutex.Unlock()
    
    if !wp.isRunning {
        return
    }
    
    wp.cancel()
    wp.wg.Wait()
    close(wp.jobQueue)
    wp.isRunning = false
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    for {
        select {
        case job := <-wp.jobQueue:
            fmt.Printf("Worker %d processing job %d: %s\n", id, job.ID, job.Data)
            time.Sleep(100 * time.Millisecond) // Simulate work
            fmt.Printf("Worker %d completed job %d\n", id, job.ID)
        case <-wp.ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, wp.ctx.Err())
            return
        }
    }
}

func (wp *WorkerPool) AddJob(job Job) error {
    wp.mutex.RLock()
    defer wp.mutex.RUnlock()
    
    if !wp.isRunning {
        return fmt.Errorf("worker pool is not running")
    }
    
    select {
    case wp.jobQueue <- job:
        return nil
    case <-wp.ctx.Done():
        return fmt.Errorf("worker pool cancelled: %v", wp.ctx.Err())
    }
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
        if err := pool.AddJob(job); err != nil {
            fmt.Printf("Error adding job: %v\n", err)
            break
        }
    }
    
    // Wait for context to be done
    <-ctx.Done()
    fmt.Printf("Context done: %v\n", ctx.Err())
    
    // Stop worker pool
    pool.Stop()
}