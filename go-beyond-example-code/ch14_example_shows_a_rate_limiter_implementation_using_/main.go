package main

import (
    "fmt"
    "time"
)

func multiplexer(inputs []<-chan string, output chan<- string) {
    defer close(output)
    
    for {
        select {
        case msg1 := <-inputs[0]:
            output <- msg1
        case msg2 := <-inputs[1]:
            output <- msg2
        case msg3 := <-inputs[2]:
            output <- msg3
        }
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)
    output := make(chan string)
    
    inputs := []<-chan string{ch1, ch2, ch3}
    
    // Start multiplexer
    go multiplexer(inputs, output)
    
    // Start producers
    go func() {
        for i := 1; i <= 3; i++ {
            ch1 <- fmt.Sprintf("Message from ch1: %d", i)
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch2 <- fmt.Sprintf("Message from ch2: %d", i)
            time.Sleep(300 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            ch3 <- fmt.Sprintf("Message from ch3: %d", i)
            time.Sleep(400 * time.Millisecond)
        }
    }()
    
    // Read from multiplexed output
    for i := 0; i < 9; i++ {
        msg := <-output
        fmt.Printf("Received: %s\n", msg)
    }
}