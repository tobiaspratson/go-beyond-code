package main

import (
    "fmt"
    "math"
)

// Calculate compound interest
func CompoundInterest(principal, rate, time float64) float64 {
    return principal * math.Exp(rate * time)
}

// Calculate half-life decay
func HalfLifeDecay(initial, halfLife, time float64) float64 {
    return initial * math.Exp(-math.Ln2 * time / halfLife)
}

// Calculate pH from hydrogen ion concentration
func CalculatePH(hydrogenIonConcentration float64) float64 {
    return -math.Log10(hydrogenIonConcentration)
}

func main() {
    // Compound interest example
    principal := 1000.0
    rate := 0.05  // 5% annual rate
    time := 10.0  // 10 years
    
    finalAmount := CompoundInterest(principal, rate, time)
    fmt.Printf("Compound interest: $%.2f -> $%.2f\n", principal, finalAmount)
    
    // Half-life decay example
    initial := 100.0
    halfLife := 5.0  // 5 years
    time = 10.0      // 10 years
    
    remaining := HalfLifeDecay(initial, halfLife, time)
    fmt.Printf("Half-life decay: %.2f -> %.2f\n", initial, remaining)
    
    // pH calculation example
    hConcentration := 0.001  // 0.001 M
    ph := CalculatePH(hConcentration)
    fmt.Printf("pH of solution: %.2f\n", ph)
}