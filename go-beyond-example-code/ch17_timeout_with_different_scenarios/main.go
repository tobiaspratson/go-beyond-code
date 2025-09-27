package main

import (
    "context"
    "fmt"
    "time"
)

func processWithTimeout(ctx context.Context, name string, duration time.Duration) {
    fmt.Printf("Starting %s (will take %v)\n", name, duration)
    
    select {
    case <-time.After(duration):
        fmt.Printf("%s completed successfully\n", name)
    case <-ctx.Done():
        fmt.Printf("%s cancelled: %v\n", name, ctx.Err())
    }
}

func main() {
    // Scenario 1: Task completes before timeout
    fmt.Println("=== Scenario 1: Task completes before timeout ===")
    ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel1()
    
    go processWithTimeout(ctx1, "Quick Task", 500*time.Millisecond)
    time.Sleep(1 * time.Second)
    
    // Scenario 2: Task times out
    fmt.Println("\n=== Scenario 2: Task times out ===")
    ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel2()
    
    go processWithTimeout(ctx2, "Slow Task", 2*time.Second)
    time.Sleep(2 * time.Second)
    
    fmt.Println("All scenarios completed")
}