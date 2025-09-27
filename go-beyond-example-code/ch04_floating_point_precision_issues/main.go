package main

import "fmt"

func main() {
    // Demonstrate precision issues
    var a float64 = 0.1
    var b float64 = 0.2
    var c float64 = a + b
    
    fmt.Printf("0.1 + 0.2 = %.17f\n", c)  // Not exactly 0.3!
    fmt.Printf("Is 0.1 + 0.2 == 0.3? %t\n", c == 0.3)
    
    // Better way to compare floats
    const epsilon = 1e-9
    fmt.Printf("Is 0.1 + 0.2 ≈ 0.3? %t\n", (c-0.3) < epsilon)
    
    // Accumulation error example
    var sum float64
    for i := 0; i < 10; i++ {
        sum += 0.1
    }
    fmt.Printf("0.1 added 10 times = %.17f\n", sum)
    fmt.Printf("Expected: 1.0, Got: %.17f\n", sum)
}