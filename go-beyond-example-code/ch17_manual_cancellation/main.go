package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(200 * time.Millisecond)
        }
    }
}

func main() {
    // Create cancellable context
    ctx, cancel := context.WithCancel(context.Background())
    
    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }
    
    // Let them work for a bit
    time.Sleep(1 * time.Second)
    
    // Cancel all workers
    fmt.Println("Cancelling all workers...")
    cancel()
    
    // Wait a bit to see cancellation
    time.Sleep(500 * time.Millisecond)
}