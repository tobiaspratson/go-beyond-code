package main

import (
    "fmt"
    "sync"
    "time"
)

func fanOut(input <-chan int, outputs []chan int) {
    defer func() {
        for _, output := range outputs {
            close(output)
        }
    }()
    
    for value := range input {
        for _, output := range outputs {
            select {
            case output <- value:
                // Successfully sent
            default:
                // Output channel is full, skip
            }
        }
    }
}

func worker(id int, input <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for value := range input {
        fmt.Printf("Worker %d processing: %d\n", id, value)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    input := make(chan int)
    outputs := make([]chan int, 3)
    
    // Create output channels
    for i := range outputs {
        outputs[i] = make(chan int, 2)
    }
    
    var wg sync.WaitGroup
    
    // Start fan-out
    go fanOut(input, outputs)
    
    // Start workers
    for i, output := range outputs {
        wg.Add(1)
        go worker(i+1, output, &wg)
    }
    
    // Send data
    for i := 1; i <= 10; i++ {
        input <- i
    }
    close(input)
    
    wg.Wait()
}