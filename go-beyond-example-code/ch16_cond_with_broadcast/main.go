package main

import (
    "fmt"
    "sync"
    "time"
)

type Event struct {
    message string
    mutex   sync.Mutex
    cond    *sync.Cond
    ready   bool
}

func NewEvent() *Event {
    e := &Event{}
    e.cond = sync.NewCond(&e.mutex)
    return e
}

func (e *Event) Wait() {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    for !e.ready {
        fmt.Println("Waiting for event...")
        e.cond.Wait()
    }
    fmt.Printf("Event received: %s\n", e.message)
}

func (e *Event) Signal(message string) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    e.message = message
    e.ready = true
    e.cond.Broadcast() // Wake up all waiting goroutines
}

func main() {
    event := NewEvent()
    var wg sync.WaitGroup
    
    // Multiple waiters
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Waiter %d starting\n", id)
            event.Wait()
            fmt.Printf("Waiter %d finished\n", id)
        }(i)
    }
    
    // Signal after delay
    time.Sleep(1 * time.Second)
    event.Signal("Hello, World!")
    
    wg.Wait()
}