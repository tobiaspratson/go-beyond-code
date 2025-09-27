package main

import (
    "fmt"
    "time"
)

func monitorChannels(channels []chan string, done <-chan bool) {
    for {
        select {
        case <-done:
            fmt.Println("Monitoring stopped")
            return
        default:
            // Check each channel's state
            for i, ch := range channels {
                select {
                case msg := <-ch:
                    fmt.Printf("Channel %d: %s\n", i, msg)
                default:
                    fmt.Printf("Channel %d: No data\n", i)
                }
            }
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    channels := make([]chan string, 3)
    for i := range channels {
        channels[i] = make(chan string, 1)
    }
    
    done := make(chan bool)
    
    // Start monitor
    go monitorChannels(channels, done)
    
    // Send data to channels at different times
    go func() {
        time.Sleep(1 * time.Second)
        channels[0] <- "Message to channel 0"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        channels[1] <- "Message to channel 1"
    }()
    
    go func() {
        time.Sleep(3 * time.Second)
        channels[2] <- "Message to channel 2"
    }()
    
    time.Sleep(4 * time.Second)
    done <- true
    time.Sleep(100 * time.Millisecond)
}