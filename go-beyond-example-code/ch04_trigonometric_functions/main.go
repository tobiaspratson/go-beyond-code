package main

import (
    "fmt"
    "math"
)

func main() {
    angle := math.Pi / 4  // 45 degrees in radians
    
    fmt.Printf("Angle: %.2f radians (%.1f degrees)\n", angle, angle*180/math.Pi)
    fmt.Printf("sin(π/4): %.6f\n", math.Sin(angle))
    fmt.Printf("cos(π/4): %.6f\n", math.Cos(angle))
    fmt.Printf("tan(π/4): %.6f\n", math.Tan(angle))
    
    // Inverse trigonometric functions
    fmt.Printf("arcsin(0.5): %.6f radians\n", math.Asin(0.5))
    fmt.Printf("arccos(0.5): %.6f radians\n", math.Acos(0.5))
    fmt.Printf("arctan(1): %.6f radians\n", math.Atan(1))
    
    // Hyperbolic functions
    fmt.Printf("sinh(1): %.6f\n", math.Sinh(1))
    fmt.Printf("cosh(1): %.6f\n", math.Cosh(1))
    fmt.Printf("tanh(1): %.6f\n", math.Tanh(1))
}