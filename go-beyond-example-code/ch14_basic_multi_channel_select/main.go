package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)
    
    // Start goroutines
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "First message"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Second message"
    }()
    
    go func() {
        time.Sleep(3 * time.Second)
        ch3 <- "Third message"
    }()
    
    // Select from multiple channels
    for i := 0; i < 3; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received: %s\n", msg2)
        case msg3 := <-ch3:
            fmt.Printf("Received: %s\n", msg3)
        }
    }
}