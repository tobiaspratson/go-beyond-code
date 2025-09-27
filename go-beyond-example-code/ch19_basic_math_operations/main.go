package main

import (
    "fmt"
    "math"
)

func main() {
    // Basic arithmetic operations
    fmt.Println("=== Basic Arithmetic ===")
    fmt.Printf("Abs(-5): %.2f\n", math.Abs(-5))           // Absolute value
    fmt.Printf("Abs(5): %.2f\n", math.Abs(5))             // Always positive
    fmt.Printf("Max(10, 20): %.2f\n", math.Max(10, 20))   // Maximum of two values
    fmt.Printf("Min(10, 20): %.2f\n", math.Min(10, 20))   // Minimum of two values
    
    // Working with multiple values
    fmt.Printf("Max of 1, 5, 3, 9: %.2f\n", math.Max(1, math.Max(5, math.Max(3, 9))))
    
    // Rounding functions - understanding the differences
    fmt.Println("\n=== Rounding Functions ===")
    testValue := 3.7
    fmt.Printf("Original value: %.2f\n", testValue)
    fmt.Printf("Ceil(%.2f): %.2f (rounds UP)\n", testValue, math.Ceil(testValue))
    fmt.Printf("Floor(%.2f): %.2f (rounds DOWN)\n", testValue, math.Floor(testValue))
    fmt.Printf("Round(%.2f): %.2f (rounds to nearest)\n", testValue, math.Round(testValue))
    fmt.Printf("Trunc(%.2f): %.2f (truncates decimal)\n", testValue, math.Trunc(testValue))
    
    // Special cases for rounding
    fmt.Println("\n=== Rounding Edge Cases ===")
    fmt.Printf("Round(2.5): %.2f\n", math.Round(2.5))     // Rounds to even (banker's rounding)
    fmt.Printf("Round(3.5): %.2f\n", math.Round(3.5))     // Rounds to even
    fmt.Printf("Round(-2.5): %.2f\n", math.Round(-2.5))   // Negative rounding
    
    // Power and roots - different ways to calculate powers
    fmt.Println("\n=== Power and Roots ===")
    fmt.Printf("Pow(2, 3): %.2f (2^3)\n", math.Pow(2, 3))
    fmt.Printf("Pow(2, 0.5): %.2f (2^0.5 = √2)\n", math.Pow(2, 0.5))
    fmt.Printf("Sqrt(16): %.2f (√16)\n", math.Sqrt(16))
    fmt.Printf("Cbrt(27): %.2f (∛27)\n", math.Cbrt(27))
    fmt.Printf("Pow(8, 1.0/3): %.2f (8^(1/3) = ∛8)\n", math.Pow(8, 1.0/3))
    
    // Comparing different power methods
    fmt.Println("\n=== Power Method Comparison ===")
    base := 2.0
    exponent := 3.0
    fmt.Printf("Pow(%.1f, %.1f): %.2f\n", base, exponent, math.Pow(base, exponent))
    fmt.Printf("Exp(%.1f * Ln(%.1f)): %.2f\n", exponent, base, math.Exp(exponent * math.Log(base)))
}