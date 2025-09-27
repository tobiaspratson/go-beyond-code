package main

import (
    "fmt"
    "math"
)

func main() {
    // Understanding angle conversion
    fmt.Println("=== Angle Conversion ===")
    degrees := 45.0
    radians := degrees * math.Pi / 180
    
    fmt.Printf("Degrees: %.2f°\n", degrees)
    fmt.Printf("Radians: %.4f rad\n", radians)
    fmt.Printf("Degrees to radians: %.2f° × π/180 = %.4f rad\n", degrees, radians)
    fmt.Printf("Radians to degrees: %.4f rad × 180/π = %.2f°\n", radians, radians * 180 / math.Pi)
    
    // Basic trigonometric functions
    fmt.Println("\n=== Basic Trigonometric Functions ===")
    fmt.Printf("Sin(45°): %.6f\n", math.Sin(radians))
    fmt.Printf("Cos(45°): %.6f\n", math.Cos(radians))
    fmt.Printf("Tan(45°): %.6f\n", math.Tan(radians))
    
    // Understanding the unit circle
    fmt.Println("\n=== Unit Circle Values ===")
    angles := []float64{0, 30, 45, 60, 90, 180, 270, 360}
    for _, angle := range angles {
        rad := angle * math.Pi / 180
        fmt.Printf("Angle: %3.0f° | Sin: %8.4f | Cos: %8.4f | Tan: %8.4f\n", 
            angle, math.Sin(rad), math.Cos(rad), math.Tan(rad))
    }
    
    // Inverse trigonometric functions (arc functions)
    fmt.Println("\n=== Inverse Trigonometric Functions ===")
    sinValue := 0.7071  // sin(45°)
    cosValue := 0.7071  // cos(45°)
    tanValue := 1.0     // tan(45°)
    
    fmt.Printf("Asin(%.4f): %.4f rad (%.2f°)\n", sinValue, math.Asin(sinValue), math.Asin(sinValue) * 180 / math.Pi)
    fmt.Printf("Acos(%.4f): %.4f rad (%.2f°)\n", cosValue, math.Acos(cosValue), math.Acos(cosValue) * 180 / math.Pi)
    fmt.Printf("Atan(%.1f): %.4f rad (%.2f°)\n", tanValue, math.Atan(tanValue), math.Atan(tanValue) * 180 / math.Pi)
    
    // Atan2 function - more robust for finding angles
    fmt.Println("\n=== Atan2 Function ===")
    x, y := 1.0, 1.0
    angle := math.Atan2(y, x)
    fmt.Printf("Atan2(%.1f, %.1f): %.4f rad (%.2f°)\n", y, x, angle, angle * 180 / math.Pi)
    
    // Hyperbolic functions
    fmt.Println("\n=== Hyperbolic Functions ===")
    value := 1.0
    fmt.Printf("Sinh(%.1f): %.6f\n", value, math.Sinh(value))
    fmt.Printf("Cosh(%.1f): %.6f\n", value, math.Cosh(value))
    fmt.Printf("Tanh(%.1f): %.6f\n", value, math.Tanh(value))
    
    // Hyperbolic identities
    fmt.Println("\n=== Hyperbolic Identities ===")
    sinh := math.Sinh(value)
    cosh := math.Cosh(value)
    fmt.Printf("cosh²(x) - sinh²(x) = %.6f (should be 1)\n", cosh*cosh - sinh*sinh)
}