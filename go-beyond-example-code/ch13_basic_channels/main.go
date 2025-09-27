package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a channel
    messages := make(chan string)
    
    // Send a message in a goroutine
    go func() {
        messages <- "Hello from goroutine!"
    }()
    
    // Receive the message
    msg := <-messages
    fmt.Println(msg)
}