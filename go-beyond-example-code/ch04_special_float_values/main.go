package main

import (
    "fmt"
    "math"
)

func main() {
    // Special floating-point values
    var positiveInf = math.Inf(1)   // Positive infinity
    var negativeInf = math.Inf(-1)  // Negative infinity
    var notANumber = math.NaN()     // Not a Number
    
    fmt.Printf("Positive infinity: %f\n", positiveInf)
    fmt.Printf("Negative infinity: %f\n", negativeInf)
    fmt.Printf("Not a Number: %f\n", notANumber)
    
    // Check for special values
    fmt.Printf("Is infinity? %t\n", math.IsInf(positiveInf, 1))
    fmt.Printf("Is NaN? %t\n", math.IsNaN(notANumber))
    
    // Operations with special values
    fmt.Printf("∞ + 1 = %f\n", positiveInf + 1)
    fmt.Printf("∞ * 0 = %f\n", positiveInf * 0)
    fmt.Printf("NaN + 1 = %f\n", notANumber + 1)
}