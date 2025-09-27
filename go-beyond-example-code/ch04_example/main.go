package main

import "fmt"

func main() {
    // Same calculation with different precision
    var f32 float32 = 3.141592653589793
    var f64 float64 = 3.141592653589793
    
    fmt.Printf("float32: %.15f\n", f32)
    fmt.Printf("float64: %.15f\n", f64)
    fmt.Printf("Difference: %.15f\n", float64(f32)-f64)
    
    // Memory usage demonstration
    fmt.Printf("float32 size: %d bytes\n", 4)
    fmt.Printf("float64 size: %d bytes\n", 8)
    
    // When precision matters
    var precise float64 = 1.0 / 3.0
    var lessPrecise float32 = 1.0 / 3.0
    
    fmt.Printf("float64 1/3: %.20f\n", precise)
    fmt.Printf("float32 1/3: %.20f\n", lessPrecise)
}