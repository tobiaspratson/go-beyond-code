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
        time.Sleep(100 * time.Millisecond) // Simulate initialization
        instance = &Singleton{
            data: "Initialized data",
        }
        fmt.Println("Singleton instance created!")
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    
    // Try to get instance from multiple goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            instance := GetInstance()
            fmt.Printf("Goroutine %d got instance: %p\n", id, instance)
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final instance: %+v\n", instance)
}