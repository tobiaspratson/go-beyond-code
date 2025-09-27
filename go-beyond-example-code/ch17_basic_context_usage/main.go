package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Create a background context
    ctx := context.Background()
    fmt.Printf("Context: %v\n", ctx)
    
    // Create a TODO context
    todoCtx := context.TODO()
    fmt.Printf("TODO Context: %v\n", todoCtx)
    
    // Check if context is done
    select {
    case <-ctx.Done():
        fmt.Println("Context is done")
    default:
        fmt.Println("Context is not done")
    }
}