package main

import "fmt"

func main() {
    fmt.Println("=== Inverted Triangle ===")
    // Inverted triangle
    for i := 5; i >= 1; i-- {
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    
    fmt.Println("\n=== Diamond Pattern ===")
    // Diamond pattern
    n := 5
    // Upper half
    for i := 1; i <= n; i++ {
        // Spaces
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        // Stars
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    // Lower half
    for i := n-1; i >= 1; i-- {
        // Spaces
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        // Stars
        for j := 1; j <= 2*i-1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    
    fmt.Println("\n=== Number Pyramid ===")
    // Number pyramid
    for i := 1; i <= 4; i++ {
        // Spaces
        for j := 1; j <= 4-i; j++ {
            fmt.Print(" ")
        }
        // Numbers
        for j := 1; j <= i; j++ {
            fmt.Printf("%d ", j)
        }
        fmt.Println()
    }
}