package main

import (
    "fmt"
    "time"
)

func processMessages(channels []<-chan string, done <-chan bool) {
    for {
        select {
        case <-done:
            fmt.Println("Shutdown signal received")
            return
        default:
            // Try to receive from any channel
            for i, ch := range channels {
                select {
                case msg := <-ch:
                    fmt.Printf("Channel %d: %s\n", i, msg)
                    goto nextIteration
                default:
                    // This channel is not ready, try next
                }
            }
            // No channels ready, wait a bit
            time.Sleep(10 * time.Millisecond)
        }
    nextIteration:
    }
}

func main() {
    channels := make([]chan string, 3)
    for i := range channels {
        channels[i] = make(chan string, 2)
    }
    
    done := make(chan bool)
    
    // Start processor
    go processMessages(convertToReadOnly(channels), done)
    
    // Send messages to different channels
    go func() {
        for i := 1; i <= 5; i++ {
            channels[0] <- fmt.Sprintf("Message %d to channel 0", i)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 5; i++ {
            channels[1] <- fmt.Sprintf("Message %d to channel 1", i)
            time.Sleep(150 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 5; i++ {
            channels[2] <- fmt.Sprintf("Message %d to channel 2", i)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    time.Sleep(2 * time.Second)
    done <- true
    time.Sleep(100 * time.Millisecond)
}

func convertToReadOnly(channels []chan string) []<-chan string {
    result := make([]<-chan string, len(channels))
    for i, ch := range channels {
        result[i] = ch
    }
    return result
}