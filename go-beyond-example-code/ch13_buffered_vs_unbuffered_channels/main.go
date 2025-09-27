package main

import (
    "fmt"
    "time"
)

func main() {
    // Unbuffered channel (synchronous)
    fmt.Println("=== Unbuffered Channel ===")
    unbuffered := make(chan string)
    
    go func() {
        fmt.Println("Sending to unbuffered channel...")
        unbuffered <- "Hello"
        fmt.Println("Sent to unbuffered channel")
    }()
    
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Receiving from unbuffered channel...")
    msg := <-unbuffered
    fmt.Printf("Received: %s\n", msg)
    
    // Buffered channel (asynchronous)
    fmt.Println("\n=== Buffered Channel ===")
    buffered := make(chan string, 2)
    
    fmt.Println("Sending to buffered channel...")
    buffered <- "Hello"
    buffered <- "World"
    fmt.Println("Sent to buffered channel")
    
    fmt.Println("Receiving from buffered channel...")
    fmt.Printf("Received: %s\n", <-buffered)
    fmt.Printf("Received: %s\n", <-buffered)
}