package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string, 2)
    
    // Fill the channel
    ch <- "First"
    ch <- "Second"
    
    // Try to send more (should fail)
    select {
    case ch <- "Third":
        fmt.Println("Message sent")
    default:
        fmt.Println("Channel is full")
    }
    
    // Try to receive
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("No message available")
    }
}