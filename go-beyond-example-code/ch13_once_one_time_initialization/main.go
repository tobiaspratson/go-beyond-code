package main

import (
    "fmt"
    "sync"
    "time"
)

type Singleton struct {
    data string
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        fmt.Println("Creating singleton instance...")
        time.Sleep(100 * time.Millisecond) // Simulate expensive initialization
        instance = &Singleton{data: "Initialized"}
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    
    // Multiple goroutines trying to get singleton
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            instance := GetInstance()
            fmt.Printf("Goroutine %d got: %s\n", id, instance.data)
        }(i)
    }
    
    wg.Wait()
}