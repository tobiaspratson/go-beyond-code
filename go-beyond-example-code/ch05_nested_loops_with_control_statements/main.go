package main

import "fmt"

func main() {
    fmt.Println("=== Nested Loops with Break ===")
    // Break in nested loops
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 5; j++ {
            if j == 3 {
                fmt.Printf("  Breaking inner loop at j=%d\n", j)
                break  // Only breaks inner loop
            }
            fmt.Printf("  Inner: %d\n", j)
        }
    }
    
    fmt.Println("\n=== Nested Loops with Continue ===")
    // Continue in nested loops
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 5; j++ {
            if j == 3 {
                fmt.Printf("  Skipping j=%d\n", j)
                continue  // Skip to next j iteration
            }
            fmt.Printf("  Inner: %d\n", j)
        }
    }
    
    fmt.Println("\n=== Breaking Out of Multiple Loops ===")
    // Using labeled break to exit multiple loops
    outer:
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 5; j++ {
            if i == 2 && j == 3 {
                fmt.Printf("  Breaking both loops at i=%d, j=%d\n", i, j)
                break outer  // Breaks out of both loops
            }
            fmt.Printf("  Inner: %d\n", j)
        }
    }
}