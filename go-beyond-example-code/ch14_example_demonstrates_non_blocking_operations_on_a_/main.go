package main

import (
    "fmt"
    "time"
)

func testChannelCapacity(ch chan string, capacity int) {
    fmt.Printf("Testing channel with capacity %d\n", capacity)
    
    // Try to fill the channel
    for i := 1; i <= capacity+2; i++ {
        msg := fmt.Sprintf("Message %d", i)
        select {
        case ch <- msg:
            fmt.Printf("✓ Sent: %s\n", msg)
        default:
            fmt.Printf("✗ Failed to send: %s (channel full)\n", msg)
        }
    }
    
    // Try to receive all messages
    for i := 0; i < capacity+2; i++ {
        select {
        case msg := <-ch:
            fmt.Printf("✓ Received: %s\n", msg)
        default:
            fmt.Printf("✗ No message available\n")
        }
    }
    
    fmt.Println("---")
}

func main() {
    // Test different channel capacities
    testChannelCapacity(make(chan string, 0), 0)  // Unbuffered
    testChannelCapacity(make(chan string, 1), 1)  // Capacity 1
    testChannelCapacity(make(chan string, 3), 3)  // Capacity 3
}