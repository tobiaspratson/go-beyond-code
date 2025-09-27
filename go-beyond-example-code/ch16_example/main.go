package main

import (
    "fmt"
    "sync"
    "time"
)

var counter int
var mutex sync.Mutex

func increment() {
    mutex.Lock()
    defer mutex.Unlock()
    counter++
}

func main() {
    var wg sync.WaitGroup
    
    // Start 1000 goroutines
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final counter: %d\n", counter)
}