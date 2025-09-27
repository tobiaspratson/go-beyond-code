package main

import (
    "fmt"
    "math"
)

func main() {
    x := 16.0
    
    fmt.Printf("x = %.2f\n", x)
    fmt.Printf("Square root: %.2f\n", math.Sqrt(x))
    fmt.Printf("Square: %.2f\n", math.Pow(x, 2))
    fmt.Printf("Cube: %.2f\n", math.Pow(x, 3))
    fmt.Printf("Absolute value of -5: %.2f\n", math.Abs(-5))
    fmt.Printf("Ceiling of 3.2: %.2f\n", math.Ceil(3.2))
    fmt.Printf("Floor of 3.8: %.2f\n", math.Floor(3.8))
    fmt.Printf("Round of 3.6: %.2f\n", math.Round(3.6))
}