package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start receiver
    go func() {
        for {
            select {
            case msg1 := <-ch1:
                fmt.Printf("Received from ch1: %s\n", msg1)
            case msg2 := <-ch2:
                fmt.Printf("Received from ch2: %s\n", msg2)
            }
        }
    }()
    
    // Send messages
    ch1 <- "Hello from ch1"
    ch2 <- "Hello from ch2"
    
    time.Sleep(100 * time.Millisecond)
}