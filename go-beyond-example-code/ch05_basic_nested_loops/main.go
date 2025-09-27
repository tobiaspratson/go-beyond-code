package main

import "fmt"

func main() {
    fmt.Println("=== Understanding Nested Loop Execution ===")
    // Simple nested loop with explanation
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop iteration %d:\n", i)
        for j := 1; j <= 3; j++ {
            fmt.Printf("  Inner loop: i=%d, j=%d\n", i, j)
        }
        fmt.Println()
    }
    
    fmt.Println("=== Multiplication Table ===")
    // Multiplication table
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            fmt.Printf("%2d ", i*j)
        }
        fmt.Println()  // New line after each row
    }
    
    fmt.Println("\n=== Triangle Pattern ===")
    // Pattern printing
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}