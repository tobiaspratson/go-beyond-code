package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type Service struct {
    running int32 // 0 = stopped, 1 = running
    workers int32
}

func (s *Service) Start() bool {
    // Try to start the service
    if atomic.CompareAndSwapInt32(&s.running, 0, 1) {
        fmt.Println("Service started")
        return true
    }
    return false // Already running
}

func (s *Service) Stop() bool {
    // Try to stop the service
    if atomic.CompareAndSwapInt32(&s.running, 1, 0) {
        fmt.Println("Service stopped")
        return true
    }
    return false // Already stopped
}

func (s *Service) IsRunning() bool {
    return atomic.LoadInt32(&s.running) == 1
}

func (s *Service) AddWorker() {
    atomic.AddInt32(&s.workers, 1)
}

func (s *Service) RemoveWorker() {
    atomic.AddInt32(&s.workers, -1)
}

func (s *Service) WorkerCount() int32 {
    return atomic.LoadInt32(&s.workers)
}

func main() {
    service := &Service{}
    var wg sync.WaitGroup
    
    // Try to start service multiple times
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            if service.Start() {
                fmt.Printf("Goroutine %d started service\n", id)
                service.AddWorker()
                time.Sleep(100 * time.Millisecond)
                service.RemoveWorker()
                service.Stop()
            } else {
                fmt.Printf("Goroutine %d: Service already running\n", id)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final worker count: %d\n", service.WorkerCount())
}