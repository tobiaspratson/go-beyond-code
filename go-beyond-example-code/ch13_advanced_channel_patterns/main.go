package main

import (
    "fmt"
    "time"
)

func fanOut(input <-chan int, outputs ...chan<- int) {
    defer func() {
        for _, output := range outputs {
            close(output)
        }
    }()
    
    for value := range input {
        for _, output := range outputs {
            output <- value
        }
    }
}

func main() {
    input := make(chan int)
    output1 := make(chan int)
    output2 := make(chan int)
    
    go fanOut(input, output1, output2)
    
    // Start receivers
    go func() {
        for value := range output1 {
            fmt.Printf("Receiver 1 got: %d\n", value)
        }
    }()
    
    go func() {
        for value := range output2 {
            fmt.Printf("Receiver 2 got: %d\n", value)
        }
    }()
    
    // Send data
    for i := 1; i <= 5; i++ {
        input <- i
        time.Sleep(100 * time.Millisecond)
    }
    close(input)
    
    time.Sleep(1 * time.Second)
}