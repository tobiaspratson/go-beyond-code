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
    jobs     []Job
    mutex    sync.Mutex
    cond     *sync.Cond
    workers  int
    wg       sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
    wp := &WorkerPool{
        workers: workers,
    }
    wp.cond = sync.NewCond(&wp.mutex)
    return wp
}

func (wp *WorkerPool) AddJob(job Job) {
    wp.mutex.Lock()
    defer wp.mutex.Unlock()
    
    wp.jobs = append(wp.jobs, job)
    fmt.Printf("Added job %d, queue length: %d\n", job.ID, len(wp.jobs))
    wp.cond.Signal() // Wake up one worker
}

func (wp *WorkerPool) GetJob() (Job, bool) {
    wp.mutex.Lock()
    defer wp.mutex.Unlock()
    
    // Wait until there's a job available
    for len(wp.jobs) == 0 {
        wp.cond.Wait()
    }
    
    // Get the first job
    job := wp.jobs[0]
    wp.jobs = wp.jobs[1:]
    return job, true
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
        job, ok := wp.GetJob()
        if !ok {
            break
        }
        
        fmt.Printf("Worker %d processing job %d: %s\n", id, job.ID, job.Data)
        time.Sleep(100 * time.Millisecond) // Simulate work
        fmt.Printf("Worker %d completed job %d\n", id, job.ID)
    }
}

func (wp *WorkerPool) Stop() {
    wp.mutex.Lock()
    wp.jobs = nil // Clear jobs to signal workers to stop
    wp.mutex.Unlock()
    wp.cond.Broadcast() // Wake up all workers
    wp.wg.Wait()
}

func main() {
    pool := NewWorkerPool(3)
    pool.Start()
    
    // Add jobs
    for i := 1; i <= 10; i++ {
        job := Job{
            ID:   i,
            Data: fmt.Sprintf("Job data %d", i),
        }
        pool.AddJob(job)
        time.Sleep(50 * time.Millisecond)
    }
    
    time.Sleep(1 * time.Second)
    pool.Stop()
    fmt.Println("All workers stopped")
}