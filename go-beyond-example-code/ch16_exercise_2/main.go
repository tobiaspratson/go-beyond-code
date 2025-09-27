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

type WorkerPool struct {
    jobQueue chan Job
    workers  int
    wg       sync.WaitGroup
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
    return &WorkerPool{
        jobQueue: make(chan Job, queueSize),
        workers:  workers,
    }
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    for job := range wp.jobQueue {
        fmt.Printf("Worker %d processing job %d: %s\n", id, job.ID, job.Data)
        time.Sleep(100 * time.Millisecond) // Simulate work
        fmt.Printf("Worker %d completed job %d\n", id, job.ID)
    }
}

func (wp *WorkerPool) AddJob(job Job) {
    wp.jobQueue <- job
}

func (wp *WorkerPool) Stop() {
    close(wp.jobQueue)
    wp.wg.Wait()
}

func main() {
    pool := NewWorkerPool(3, 10)
    pool.Start()
    
    // Add jobs
    for i := 1; i <= 10; i++ {
        job := Job{
            ID:   i,
            Data: fmt.Sprintf("Job data %d", i),
        }
        pool.AddJob(job)
    }
    
    pool.Stop()
    fmt.Println("All jobs completed")
}