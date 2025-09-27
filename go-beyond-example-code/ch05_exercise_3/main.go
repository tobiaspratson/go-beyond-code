package main

import "fmt"

func main() {
    // Square pattern
    fmt.Println("Square Pattern:")
    for i := 1; i <= 4; i++ {
        for j := 1; j <= 4; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
    
    // Right triangle
    fmt.Println("\nRight Triangle:")
    for i := 1; i <= 4; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
    
    // Inverted triangle
    fmt.Println("\nInverted Triangle:")
    for i := 4; i >= 1; i-- {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}