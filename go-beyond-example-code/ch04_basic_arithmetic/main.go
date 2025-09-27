package main

import "fmt"

func main() {
    a, b := 10, 3
    
    fmt.Printf("a = %d, b = %d\n", a, b)
    fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
    fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
    fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
    fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
    fmt.Printf("Remainder: %d %% %d = %d\n", a, b, a%b)
    
    // Floating-point division
    fmt.Printf("Float division: %.2f / %.2f = %.2f\n", 
        float64(a), float64(b), float64(a)/float64(b))
}