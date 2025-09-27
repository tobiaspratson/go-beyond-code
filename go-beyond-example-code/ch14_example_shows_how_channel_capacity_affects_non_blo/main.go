package main

import (
    "fmt"
    "time"
)

func nonBlockingWithTimeout(ch <-chan string, timeout time.Duration) (string, bool) {
    select {
    case msg := <-ch:
        return msg, true
    case <-time.After(timeout):
        return "", false
    }
}

func main() {
    ch := make(chan string)
    
    // Test with empty channel (should timeout)
    if msg, ok := nonBlockingWithTimeout(ch, 100*time.Millisecond); ok {
        fmt.Printf("Received: %s\n", msg)
    } else {
        fmt.Println("Timeout - no message received")
    }
    
    // Send a message after timeout
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch <- "Delayed message"
    }()
    
    // Test with message coming after timeout
    if msg, ok := nonBlockingWithTimeout(ch, 100*time.Millisecond); ok {
        fmt.Printf("Received: %s\n", msg)
    } else {
        fmt.Println("Timeout - no message received")
    }
    
    // Test with message coming before timeout
    if msg, ok := nonBlockingWithTimeout(ch, 500*time.Millisecond); ok {
        fmt.Printf("Received: %s\n", msg)
    } else {
        fmt.Println("Timeout - no message received")
    }
}