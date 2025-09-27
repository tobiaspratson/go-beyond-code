package main

import (
    "fmt"
    "sync"
    "time"
)

func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for value := range ch {
                output <- value
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

func main() {
    input1 := make(chan int)
    input2 := make(chan int)
    input3 := make(chan int)
    
    output := fanIn(input1, input2, input3)
    
    // Start senders
    go func() {
        for i := 1; i <= 3; i++ {
            input1 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(input1)
    }()
    
    go func() {
        for i := 4; i <= 6; i++ {
            input2 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(input2)
    }()
    
    go func() {
        for i := 7; i <= 9; i++ {
            input3 <- i
            time.Sleep(100 * time.Millisecond)
        }
        close(input3)
    }()
    
    // Receive from combined output
    for value := range output {
        fmt.Printf("Received: %d\n", value)
    }
}