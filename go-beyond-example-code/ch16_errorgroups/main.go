package main

import (
    "fmt"
    "sync"
    "time"
)

// Errorgroup for handling errors from workers
type Errorgroup struct {
    wg     sync.WaitGroup
    errors chan error
    mu     sync.Mutex
}

func NewErrorgroup() *Errorgroup {
    return &Errorgroup{
        errors: make(chan error, 10),
    }
}

func (eg *Errorgroup) AddWorker(fn func() error) {
    eg.wg.Add(1)
    go func() {
        defer eg.wg.Done()
        if err := fn(); err != nil {
            eg.errors <- err
        }
    }()
}

func (eg *Errorgroup) Wait() []error {
    go func() {
        eg.wg.Wait()
        close(eg.errors)
    }()
    
    var errors []error
    for err := range eg.errors {
        errors = append(errors, err)
    }
    
    return errors
}

func main() {
    // Create an errorgroup
    eg := NewErrorgroup()
    
    // Add workers that might fail
    eg.AddWorker(func() error {
        time.Sleep(1 * time.Second)
        fmt.Println("Worker 1 completed successfully")
        return nil
    })
    
    eg.AddWorker(func() error {
        time.Sleep(2 * time.Second)
        fmt.Println("Worker 2 failed")
        return fmt.Errorf("worker 2 failed")
    })
    
    eg.AddWorker(func() error {
        time.Sleep(3 * time.Second)
        fmt.Println("Worker 3 completed successfully")
        return nil
    })
    
    // Wait for all workers and collect errors
    errors := eg.Wait()
    
    if len(errors) == 0 {
        fmt.Println("All workers completed successfully!")
    } else {
        fmt.Printf("Workers completed with %d errors:\n", len(errors))
        for i, err := range errors {
            fmt.Printf("Error %d: %v\n", i+1, err)
        }
    }
}