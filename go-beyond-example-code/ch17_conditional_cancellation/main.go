package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int, shouldStop chan bool) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        case <-shouldStop:
            fmt.Printf("Worker %d received stop signal\n", id)
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(200 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    stopSignal := make(chan bool)
    
    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i, stopSignal)
    }
    
    // Let them work for a bit
    time.Sleep(1 * time.Second)
    
    // Option 1: Cancel via context
    fmt.Println("Cancelling via context...")
    cancel()
    
    time.Sleep(500 * time.Millisecond)
    
    // Option 2: Cancel via custom signal
    fmt.Println("Sending stop signal...")
    close(stopSignal)
    
    time.Sleep(500 * time.Millisecond)
}