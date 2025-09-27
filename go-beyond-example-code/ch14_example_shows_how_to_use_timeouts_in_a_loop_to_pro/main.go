package main

import (
    "fmt"
    "time"
)

func adaptiveTimeout(initialTimeout time.Duration, maxTimeout time.Duration) {
    timeout := initialTimeout
    ch := make(chan string)
    
    // Start a goroutine that sends messages at random intervals
    go func() {
        for i := 1; i <= 10; i++ {
            time.Sleep(time.Duration(i*100) * time.Millisecond)
            ch <- fmt.Sprintf("Message %d", i)
        }
    }()
    
    for i := 0; i < 10; i++ {
        select {
        case msg := <-ch:
            fmt.Printf("Received: %s (timeout was %v)\n", msg, timeout)
            // Reset timeout to initial value on successful receive
            timeout = initialTimeout
        case <-time.After(timeout):
            fmt.Printf("Timeout after %v, increasing to %v\n", timeout, timeout*2)
            // Increase timeout for next iteration
            if timeout*2 <= maxTimeout {
                timeout *= 2
            }
        }
    }
}

func main() {
    fmt.Println("Adaptive timeout pattern:")
    adaptiveTimeout(200*time.Millisecond, 2*time.Second)
}