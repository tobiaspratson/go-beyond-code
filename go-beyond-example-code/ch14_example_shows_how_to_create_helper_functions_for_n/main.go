package main

import (
    "fmt"
    "time"
)

func processChannels(channels []chan string) {
    for {
        allEmpty := true
        
        // Check all channels non-blocking
        for i, ch := range channels {
            select {
            case msg := <-ch:
                fmt.Printf("Channel %d: %s\n", i, msg)
                allEmpty = false
            default:
                // This channel is empty
            }
        }
        
        if allEmpty {
            fmt.Println("All channels empty, waiting...")
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func main() {
    channels := make([]chan string, 3)
    for i := range channels {
        channels[i] = make(chan string, 2)
    }
    
    // Start processor
    go processChannels(channels)
    
    // Send messages to different channels
    go func() {
        for i := 1; i <= 5; i++ {
            channels[0] <- fmt.Sprintf("Message %d to channel 0", i)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            channels[1] <- fmt.Sprintf("Message %d to channel 1", i)
            time.Sleep(300 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 2; i++ {
            channels[2] <- fmt.Sprintf("Message %d to channel 2", i)
            time.Sleep(500 * time.Millisecond)
        }
    }()
    
    time.Sleep(3 * time.Second)
}