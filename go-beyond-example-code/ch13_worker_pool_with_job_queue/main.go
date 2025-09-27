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
    workers    int
    jobs       chan Job
    results    chan string
    quit       chan bool
    wg         sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
    return &WorkerPool{
        workers: workers,
        jobs:    make(chan Job, 100),
        results: make(chan string, 100),
        quit:    make(chan bool),
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
    
    for {
        select {
        case job := <-wp.jobs:
            // Process job
            time.Sleep(100 * time.Millisecond) // Simulate work
            result := fmt.Sprintf("Worker %d processed job %d: %s", id, job.ID, job.Data)
            wp.results <- result
            
        case <-wp.quit:
            fmt.Printf("Worker %d shutting down\n", id)
            return
        }
    }
}

func (wp *WorkerPool) Submit(job Job) {
    wp.jobs <- job
}

func (wp *WorkerPool) Stop() {
    close(wp.quit)
    wp.wg.Wait()
    close(wp.results)
}

func (wp *WorkerPool) Results() <-chan string {
    return wp.results
}

func main() {
    pool := NewWorkerPool(3)
    pool.Start()
    
    // Start result collector
    go func() {
        for result := range pool.Results() {
            fmt.Println(result)
        }
    }()
    
    // Submit jobs
    for i := 1; i <= 10; i++ {
        job := Job{
            ID:   i,
            Data: fmt.Sprintf("Job data %d", i),
        }
        pool.Submit(job)
    }
    
    time.Sleep(2 * time.Second)
    pool.Stop()
}