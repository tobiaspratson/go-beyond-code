package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type AtomicCounter struct {
    value int64
}

func (ac *AtomicCounter) Increment() {
    atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&ac.value)
}

func worker(counter *AtomicCounter, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        counter.Increment()
    }
}

func main() {
    counter := &AtomicCounter{}
    var wg sync.WaitGroup
    
    // Start 10 workers
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(counter, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.Value())
}