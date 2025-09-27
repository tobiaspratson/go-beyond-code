package main

import (
    "fmt"
    "sync"
    "time"
)

type Counter struct {
    value int
    mutex sync.Mutex
}

func (c *Counter) Increment() {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.value
}

func worker(counter *Counter, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        counter.Increment()
    }
}

func main() {
    counter := &Counter{}
    var wg sync.WaitGroup
    
    // Start 10 workers
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(counter, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.Value())
}