package main

import (
    "fmt"
    "math"
    "strconv"
)

type Calculator struct {
    history []string
}

func (c *Calculator) add(a, b float64) float64 {
    result := a + b
    c.history = append(c.history, fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) subtract(a, b float64) float64 {
    result := a - b
    c.history = append(c.history, fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) multiply(a, b float64) float64 {
    result := a * b
    c.history = append(c.history, fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    result := a / b
    c.history = append(c.history, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) power(a, b float64) float64 {
    result := math.Pow(a, b)
    c.history = append(c.history, fmt.Sprintf("%.2f ^ %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) sqrt(a float64) (float64, error) {
    if a < 0 {
        return 0, fmt.Errorf("square root of negative number")
    }
    result := math.Sqrt(a)
    c.history = append(c.history, fmt.Sprintf("√%.2f = %.2f", a, result))
    return result, nil
}

func (c *Calculator) showHistory() {
    fmt.Println("Calculation History:")
    for i, entry := range c.history {
        fmt.Printf("%d. %s\n", i+1, entry)
    }
}

func main() {
    calc := &Calculator{}
    
    // Test operations
    fmt.Printf("Addition: %.2f\n", calc.add(10, 5))
    fmt.Printf("Subtraction: %.2f\n", calc.subtract(10, 5))
    fmt.Printf("Multiplication: %.2f\n", calc.multiply(10, 5))
    
    result, err := calc.divide(10, 5)
    if err == nil {
        fmt.Printf("Division: %.2f\n", result)
    }
    
    result, err = calc.divide(10, 0)
    if err != nil {
        fmt.Printf("Division error: %v\n", err)
    }
    
    fmt.Printf("Power: %.2f\n", calc.power(2, 3))
    
    result, err = calc.sqrt(16)
    if err == nil {
        fmt.Printf("Square root: %.2f\n", result)
    }
    
    result, err = calc.sqrt(-4)
    if err != nil {
        fmt.Printf("Square root error: %v\n", err)
    }
    
    calc.showHistory()
}