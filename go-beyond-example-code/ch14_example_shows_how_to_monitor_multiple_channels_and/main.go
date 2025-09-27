package main

import "fmt"

func main() {
    ch := make(chan string)
    
    // Try to receive from channel
    select {
    case msg := <-ch:
        fmt.Printf("Received: %s\n", msg)
    default:
        fmt.Println("No message available")
    }
    
    // Try to send to channel
    select {
    case ch <- "Hello":
        fmt.Println("Message sent")
    default:
        fmt.Println("Channel is full or blocked")
    }
}