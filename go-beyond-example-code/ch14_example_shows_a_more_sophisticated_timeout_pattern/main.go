package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start operations with different durations
    go func() {
        time.Sleep(500 * time.Millisecond)
        ch1 <- "Fast operation"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Slow operation"
    }()
    
    // Wait for either operation or timeout
    select {
    case msg1 := <-ch1:
        fmt.Printf("Received from ch1: %s\n", msg1)
    case msg2 := <-ch2:
        fmt.Printf("Received from ch2: %s\n", msg2)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout - neither operation completed")
    }
    
    // Wait for remaining operation
    select {
    case msg2 := <-ch2:
        fmt.Printf("Received from ch2: %s\n", msg2)
    case <-time.After(2 * time.Second):
        fmt.Println("Second timeout")
    }
}