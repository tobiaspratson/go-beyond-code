package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    // Start a goroutine that sends messages
    go func() {
        for i := 1; i <= 5; i++ {
            time.Sleep(500 * time.Millisecond)
            ch <- fmt.Sprintf("Message %d", i)
        }
    }()
    
    // Process messages with periodic timeout
    for {
        select {
        case msg := <-ch:
            fmt.Printf("Received: %s\n", msg)
        case <-time.After(1 * time.Second):
            fmt.Println("Timeout - no message received")
            return
        }
    }
}