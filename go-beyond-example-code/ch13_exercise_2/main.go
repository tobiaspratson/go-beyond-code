package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    value int
    mutex sync.Mutex
}

func (sc *SafeCounter) Increment() {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    sc.value++
}

func (sc *SafeCounter) Value() int {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    return sc.value
}

func worker(counter *SafeCounter, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        counter.Increment()
    }
}

func main() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup
    
    // Start 10 workers
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(counter, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.Value())
}