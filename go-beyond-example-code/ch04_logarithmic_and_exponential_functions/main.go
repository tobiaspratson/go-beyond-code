package main

import (
    "fmt"
    "math"
)

func main() {
    x := 2.0
    
    // Natural logarithm (base e)
    fmt.Printf("ln(%f) = %.6f\n", x, math.Log(x))
    
    // Base 10 logarithm
    fmt.Printf("log10(%f) = %.6f\n", x, math.Log10(x))
    
    // Base 2 logarithm
    fmt.Printf("log2(%f) = %.6f\n", x, math.Log2(x))
    
    // Exponential functions
    fmt.Printf("e^%f = %.6f\n", x, math.Exp(x))
    fmt.Printf("2^%f = %.6f\n", x, math.Exp2(x))
    fmt.Printf("10^%f = %.6f\n", x, math.Pow10(int(x)))
    
    // Power function
    fmt.Printf("%f^%f = %.6f\n", x, 3.0, math.Pow(x, 3.0))
    
    // Square root and cube root
    fmt.Printf("sqrt(%f) = %.6f\n", x, math.Sqrt(x))
    fmt.Printf("cbrt(%f) = %.6f\n", x*x*x, math.Cbrt(x*x*x))
}