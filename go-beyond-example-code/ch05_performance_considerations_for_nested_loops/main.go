package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== Nested Loop Performance ===")
    
    // Example of O(n²) complexity
    n := 1000
    start := time.Now()
    
    count := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            count++  // This runs n² times
        }
    }
    
    elapsed := time.Since(start)
    fmt.Printf("Nested loop with n=%d took %v\n", n, elapsed)
    fmt.Printf("Total iterations: %d\n", count)
    
    fmt.Println("\n=== Optimizing Nested Loops ===")
    // Sometimes you can optimize by reducing inner loop iterations
    start = time.Now()
    count = 0
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {  // Start j from i instead of 0
            count++
        }
    }
    elapsed = time.Since(start)
    fmt.Printf("Optimized nested loop took %v\n", elapsed)
    fmt.Printf("Total iterations: %d\n", count)
}