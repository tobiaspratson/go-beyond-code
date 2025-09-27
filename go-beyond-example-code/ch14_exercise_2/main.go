package main

import (
    "fmt"
    "time"
)

func dynamicMultiplexer(inputs []<-chan string, output chan<- string, done <-chan bool) {
    defer close(output)
    
    for {
        select {
        case <-done:
            fmt.Println("Multiplexer shutting down")
            return
        default:
            // Try to receive from any input channel
            for i, input := range inputs {
                select {
                case msg := <-input:
                    output <- fmt.Sprintf("Channel %d: %s", i, msg)
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
    // Create variable number of channels
    numChannels := 5
    inputs := make([]chan string, numChannels)
    for i := range inputs {
        inputs[i] = make(chan string, 2)
    }
    
    output := make(chan string, 10)
    done := make(chan bool)
    
    // Start multiplexer
    go dynamicMultiplexer(convertToReadOnly(inputs), output, done)
    
    // Start producers
    for i, ch := range inputs {
        go func(id int, ch chan<- string) {
            for j := 1; j <= 3; j++ {
                ch <- fmt.Sprintf("Message %d", j)
                time.Sleep(time.Duration(id*100+200) * time.Millisecond)
            }
        }(i, ch)
    }
    
    // Collect results
    go func() {
        for msg := range output {
            fmt.Printf("Received: %s\n", msg)
        }
    }()
    
    time.Sleep(3 * time.Second)
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