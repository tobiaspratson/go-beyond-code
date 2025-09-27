package main

import (
    "fmt"
    "sync"
    "time"
)

type WorkerPool struct {
    workers    int
    jobs       chan func()
    wg         sync.WaitGroup
    quit       chan bool
}

func NewWorkerPool(workers int) *WorkerPool {
    return &WorkerPool{
        workers: workers,
        jobs:    make(chan func()),
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
            fmt.Printf("Worker %d executing job\n", id)
            job()
        case <-wp.quit:
            fmt.Printf("Worker %d shutting down\n", id)
            return
        }
    }
}

func (wp *WorkerPool) Submit(job func()) {
    wp.jobs <- job
}

func (wp *WorkerPool) Stop() {
    close(wp.quit)
    wp.wg.Wait()
}

func main() {
    pool := NewWorkerPool(3)
    pool.Start()
    
    // Submit jobs
    for i := 0; i < 10; i++ {
        i := i // Capture loop variable
        pool.Submit(func() {
            fmt.Printf("Executing job %d\n", i)
            time.Sleep(100 * time.Millisecond)
        })
    }
    
    time.Sleep(2 * time.Second)
    pool.Stop()
}