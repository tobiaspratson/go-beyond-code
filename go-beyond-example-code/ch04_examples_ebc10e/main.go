package main

import "fmt"

func main() {
    // Float32 - less precise, smaller memory
    var price float32 = 19.99
    var temperature float32 = -5.5
    
    // Float64 - more precise, larger memory (default for literals)
    var pi float64 = 3.141592653589793
    var e float64 = 2.718281828459045
    
    // Scientific notation
    var bigNumber float64 = 1.23e6  // 1,230,000
    var smallNumber float64 = 1.23e-6  // 0.00000123
    
    fmt.Printf("Price: %.2f\n", price)
    fmt.Printf("Temperature: %.1f°C\n", temperature)
    fmt.Printf("Pi: %.10f\n", pi)
    fmt.Printf("E: %.10f\n", e)
    fmt.Printf("Big number: %.0f\n", bigNumber)
    fmt.Printf("Small number: %.6f\n", smallNumber)
}