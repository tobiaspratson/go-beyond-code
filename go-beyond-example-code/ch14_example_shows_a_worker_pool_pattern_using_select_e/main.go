package main

import (
    "fmt"
    "sync"
    "time"
)

type Priority int

const (
    Low Priority = iota
    Normal
    High
)

type Job struct {
    ID       int
    Data     string
    Priority Priority
}

type Worker struct {
    ID           int
    JobChan      chan Job
    QuitChan     chan bool
    TimeoutChan  <-chan time.Time
}

func NewWorker(id int, jobChan chan Job, timeout time.Duration) *Worker {
    return &Worker{
        ID:          id,
        JobChan:     jobChan,
        QuitChan:    make(chan bool),
        TimeoutChan: time.After(timeout),
    }
}

func (w *Worker) Start() {
    go func() {
        for {
            select {
            case job := <-w.JobChan:
                fmt.Printf("Worker %d processing %s job %d: %s\n", 
                    w.ID, priorityString(job.Priority), job.ID, job.Data)
                time.Sleep(time.Duration(job.Priority+1) * 200 * time.Millisecond)
                fmt.Printf("Worker %d completed job %d\n", w.ID, job.ID)
            case <-w.QuitChan:
                fmt.Printf("Worker %d quitting\n", w.ID)
                return
            case <-w.TimeoutChan:
                fmt.Printf("Worker %d timed out\n", w.ID)
                return
            }
        }
    }()
}

func (w *Worker) Stop() {
    w.QuitChan <- true
}

func priorityString(p Priority) string {
    switch p {
    case High:
        return "HIGH"
    case Normal:
        return "NORMAL"
    case Low:
        return "LOW"
    default:
        return "UNKNOWN"
    }
}

func main() {
    jobChan := make(chan Job, 20)
    
    // Start workers with different timeouts
    workers := make([]*Worker, 3)
    for i := 0; i < 3; i++ {
        workers[i] = NewWorker(i+1, jobChan, 5*time.Second)
        workers[i].Start()
    }
    
    // Send jobs with different priorities
    priorities := []Priority{High, Normal, Low, High, Normal}
    for i := 1; i <= 5; i++ {
        job := Job{
            ID:       i,
            Data:     fmt.Sprintf("Job data %d", i),
            Priority: priorities[i-1],
        }
        jobChan <- job
    }
    
    // Wait for processing
    time.Sleep(3 * time.Second)
    
    // Stop workers
    for _, worker := range workers {
        worker.Stop()
    }
}