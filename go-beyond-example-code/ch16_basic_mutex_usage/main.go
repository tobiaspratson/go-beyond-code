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

func (sc *SafeCounter) Add(amount int) {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    sc.value += amount
}

func main() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup
    
    // Start multiple goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                counter.Increment()
                if j%50 == 0 {
                    fmt.Printf("Goroutine %d: counter = %d\n", id, counter.Value())
                }
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter.Value())
}