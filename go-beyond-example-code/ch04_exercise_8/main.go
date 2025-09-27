package main

import (
    "fmt"
    "math"
)

// Safe division that handles division by zero
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Calculate compound interest
func compoundInterest(principal, rate float64, years int) float64 {
    return principal * math.Pow(1+rate, float64(years))
}

// Calculate distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Sqrt(dx*dx + dy*dy)
}

// Calculate factorial
func factorial(n int) int64 {
    if n < 0 {
        return -1  // Error indicator
    }
    if n <= 1 {
        return 1
    }
    result := int64(1)
    for i := 2; i <= n; i++ {
        result *= int64(i)
    }
    return result
}

// Calculate greatest common divisor
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Calculate least common multiple
func lcm(a, b int) int {
    return (a * b) / gcd(a, b)
}

func main() {
    // Test safe division
    result, err := safeDivide(10, 2)
    if err == nil {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }
    
    result, err = safeDivide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Test compound interest
    final := compoundInterest(1000, 0.05, 10)
    fmt.Printf("Compound interest: $%.2f\n", final)
    
    // Test distance
    dist := distance(0, 0, 3, 4)
    fmt.Printf("Distance: %.2f\n", dist)
    
    // Test factorial
    fact := factorial(5)
    fmt.Printf("5! = %d\n", fact)
    
    // Test GCD and LCM
    fmt.Printf("GCD(48, 18) = %d\n", gcd(48, 18))
    fmt.Printf("LCM(12, 18) = %d\n", lcm(12, 18))
}