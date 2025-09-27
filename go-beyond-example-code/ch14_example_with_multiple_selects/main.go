package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Start goroutines
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(500 * time.Millisecond)
            ch1 <- fmt.Sprintf("ch1 message %d", i)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(750 * time.Millisecond)
            ch2 <- fmt.Sprintf("ch2 message %d", i)
        }
    }()
    
    // Multiple select statements
    for i := 0; i < 6; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Printf("Received from ch1: %s\n", msg1)
        case msg2 := <-ch2:
            fmt.Printf("Received from ch2: %s\n", msg2)
        }
    }
    
    fmt.Println("All messages received")
}