package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start goroutines that send to channels
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from channel 1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from channel 2"
    }()
    
    // Select statement - waits for any channel to be ready
    select {
    case msg1 := <-ch1:
        fmt.Printf("Received from ch1: %s\n", msg1)
    case msg2 := <-ch2:
        fmt.Printf("Received from ch2: %s\n", msg2)
    }
    
    fmt.Println("Select completed")
}