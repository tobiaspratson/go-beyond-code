package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    // Start a goroutine that might be slow
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Slow response"
    }()
    
    // Wait with timeout
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout! Operation took too long")
    }
}