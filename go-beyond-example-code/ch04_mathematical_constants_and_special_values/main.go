package main

import (
    "fmt"
    "math"
)

func main() {
    // Mathematical constants
    fmt.Printf("Pi: %.10f\n", math.Pi)
    fmt.Printf("E: %.10f\n", math.E)
    fmt.Printf("Phi (golden ratio): %.10f\n", math.Phi)
    fmt.Printf("Sqrt2: %.10f\n", math.Sqrt2)
    fmt.Printf("SqrtE: %.10f\n", math.SqrtE)
    fmt.Printf("Ln2: %.10f\n", math.Ln2)
    fmt.Printf("Ln10: %.10f\n", math.Ln10)
    
    // Special values
    fmt.Printf("Max float64: %.2e\n", math.MaxFloat64)
    fmt.Printf("Smallest positive float64: %.2e\n", math.SmallestNonzeroFloat64)
    fmt.Printf("Max int64: %d\n", math.MaxInt64)
    fmt.Printf("Min int64: %d\n", math.MinInt64)
}