package main

import "fmt"

func main() {
    // Traditional for loop - the most common pattern
    for i := 1; i <= 5; i++ {
        fmt.Printf("Count: %d\n", i)
    }
    
    // Let's see what happens step by step
    fmt.Println("\nStep-by-step breakdown:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("Before iteration %d: i = %d\n", i, i)
        fmt.Printf("During iteration %d: processing...\n", i)
        fmt.Printf("After iteration %d: i will be %d\n", i, i+1)
    }
}