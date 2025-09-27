package main

import (
    "fmt"
    "math"
)

func main() {
    // Modulo operations
    fmt.Println("=== Modulo Operations ===")
    fmt.Printf("Mod(10, 3): %.2f\n", math.Mod(10, 3))     // 10 % 3
    fmt.Printf("Mod(-10, 3): %.2f\n", math.Mod(-10, 3))   // Negative modulo
    fmt.Printf("Remainder(10, 3): %.2f\n", math.Remainder(10, 3)) // IEEE remainder
    
    // Sign and comparison functions
    fmt.Println("\n=== Sign and Comparison ===")
    fmt.Printf("Signbit(-5.0): %t\n", math.Signbit(-5.0)) // Check if negative
    fmt.Printf("Signbit(5.0): %t\n", math.Signbit(5.0))   // Check if negative
    fmt.Printf("IsNaN(math.NaN()): %t\n", math.IsNaN(math.NaN()))
    fmt.Printf("IsInf(math.Inf(1), 1): %t\n", math.IsInf(math.Inf(1), 1))
    
    // Copy sign function
    fmt.Println("\n=== Copy Sign ===")
    fmt.Printf("Copysign(5.0, -1.0): %.2f\n", math.Copysign(5.0, -1.0)) // Copy sign from -1 to 5
    fmt.Printf("Copysign(-5.0, 1.0): %.2f\n", math.Copysign(-5.0, 1.0)) // Copy sign from 1 to -5
    
    // Next and previous representable values
    fmt.Println("\n=== Next/Previous Values ===")
    value := 1.0
    fmt.Printf("Nextafter(%.2f, 2.0): %.10f\n", value, math.Nextafter(value, 2.0))
    fmt.Printf("Nextafter(%.2f, 0.0): %.10f\n", value, math.Nextafter(value, 0.0))
}