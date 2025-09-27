package main

import (
    "fmt"
    "time"
)

func nonBlockingReceive(ch <-chan string) (string, bool) {
    select {
    case msg := <-ch:
        return msg, true
    default:
        return "", false
    }
}

func nonBlockingSend(ch chan<- string, msg string) bool {
    select {
    case ch <- msg:
        return true
    default:
        return false
    }
}

func main() {
    ch := make(chan string, 2)
    
    // Test non-blocking receive on empty channel
    if msg, ok := nonBlockingReceive(ch); ok {
        fmt.Printf("Received: %s\n", msg)
    } else {
        fmt.Println("No message available")
    }
    
    // Test non-blocking send
    if nonBlockingSend(ch, "Hello") {
        fmt.Println("Message sent successfully")
    } else {
        fmt.Println("Failed to send message")
    }
    
    // Test non-blocking receive on channel with data
    if msg, ok := nonBlockingReceive(ch); ok {
        fmt.Printf("Received: %s\n", msg)
    } else {
        fmt.Println("No message available")
    }
}