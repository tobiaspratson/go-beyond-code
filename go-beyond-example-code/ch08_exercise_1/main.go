package main

import (
    "errors"
    "fmt"
    "math"
)

// Custom error types for calculator
type CalculatorError struct {
    Operation string
    Message   string
    Value     float64
}

func (e CalculatorError) Error() string {
    return fmt.Sprintf("calculator error in %s: %s (value: %.2f)", 
        e.Operation, e.Message, e.Value)
}

var (
    ErrDivisionByZero = errors.New("division by zero")
    ErrInvalidInput   = errors.New("invalid input")
    ErrOverflow       = errors.New("result overflow")
    ErrUnderflow      = errors.New("result underflow")
)

type Calculator struct {
    history []string
}

func NewCalculator() *Calculator {
    return &Calculator{
        history: make([]string, 0),
    }
}

func (c *Calculator) Add(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, CalculatorError{
            Operation: "ADD",
            Message:   "invalid input: NaN",
            Value:     a,
        }
    }
    
    result := a + b
    if math.IsInf(result, 0) {
        return 0, CalculatorError{
            Operation: "ADD",
            Message:   "result overflow",
            Value:     result,
        }
    }
    
    c.history = append(c.history, fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) Subtract(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, CalculatorError{
            Operation: "SUBTRACT",
            Message:   "invalid input: NaN",
            Value:     a,
        }
    }
    
    result := a - b
    if math.IsInf(result, 0) {
        return 0, CalculatorError{
            Operation: "SUBTRACT",
            Message:   "result overflow",
            Value:     result,
        }
    }
    
    c.history = append(c.history, fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) Multiply(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, CalculatorError{
            Operation: "MULTIPLY",
            Message:   "invalid input: NaN",
            Value:     a,
        }
    }
    
    result := a * b
    if math.IsInf(result, 0) {
        return 0, CalculatorError{
            Operation: "MULTIPLY",
            Message:   "result overflow",
            Value:     result,
        }
    }
    
    c.history = append(c.history, fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if math.IsNaN(a) || math.IsNaN(b) {
        return 0, CalculatorError{
            Operation: "DIVIDE",
            Message:   "invalid input: NaN",
            Value:     a,
        }
    }
    
    if b == 0 {
        return 0, CalculatorError{
            Operation: "DIVIDE",
            Message:   "division by zero",
            Value:     b,
        }
    }
    
    result := a / b
    if math.IsInf(result, 0) {
        return 0, CalculatorError{
            Operation: "DIVIDE",
            Message:   "result overflow",
            Value:     result,
        }
    }
    
    c.history = append(c.history, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) Power(base, exponent float64) (float64, error) {
    if math.IsNaN(base) || math.IsNaN(exponent) {
        return 0, CalculatorError{
            Operation: "POWER",
            Message:   "invalid input: NaN",
            Value:     base,
        }
    }
    
    result := math.Pow(base, exponent)
    if math.IsNaN(result) {
        return 0, CalculatorError{
            Operation: "POWER",
            Message:   "result is NaN",
            Value:     result,
        }
    }
    
    if math.IsInf(result, 0) {
        return 0, CalculatorError{
            Operation: "POWER",
            Message:   "result overflow",
            Value:     result,
        }
    }
    
    c.history = append(c.history, fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
    return result, nil
}

func (c *Calculator) Sqrt(value float64) (float64, error) {
    if math.IsNaN(value) {
        return 0, CalculatorError{
            Operation: "SQRT",
            Message:   "invalid input: NaN",
            Value:     value,
        }
    }
    
    if value < 0 {
        return 0, CalculatorError{
            Operation: "SQRT",
            Message:   "cannot take square root of negative number",
            Value:     value,
        }
    }
    
    result := math.Sqrt(value)
    c.history = append(c.history, fmt.Sprintf("√%.2f = %.2f", value, result))
    return result, nil
}

func (c *Calculator) GetHistory() []string {
    return c.history
}

func (c *Calculator) ClearHistory() {
    c.history = make([]string, 0)
}

func main() {
    calc := NewCalculator()
    
    fmt.Println("=== Calculator Tests ===")
    
    // Test normal operations
    operations := []struct {
        name string
        fn   func(float64, float64) (float64, error)
        a, b float64
    }{
        {"Add", calc.Add, 10, 5},
        {"Subtract", calc.Subtract, 10, 3},
        {"Multiply", calc.Multiply, 4, 7},
        {"Divide", calc.Divide, 15, 3},
        {"Power", calc.Power, 2, 8},
    }
    
    for _, op := range operations {
        result, err := op.fn(op.a, op.b)
        if err != nil {
            fmt.Printf("%s error: %v\n", op.name, err)
        } else {
            fmt.Printf("%s: %.2f\n", op.name, result)
        }
    }
    
    // Test square root
    result, err := calc.Sqrt(16)
    if err != nil {
        fmt.Printf("Sqrt error: %v\n", err)
    } else {
        fmt.Printf("Sqrt: %.2f\n", result)
    }
    
    // Test error conditions
    fmt.Println("\n=== Error Tests ===")
    
    // Division by zero
    _, err = calc.Divide(10, 0)
    if err != nil {
        fmt.Printf("Division by zero: %v\n", err)
    }
    
    // Negative square root
    _, err = calc.Sqrt(-4)
    if err != nil {
        fmt.Printf("Negative sqrt: %v\n", err)
    }
    
    // NaN input
    _, err = calc.Add(math.NaN(), 5)
    if err != nil {
        fmt.Printf("NaN input: %v\n", err)
    }
    
    // Print history
    fmt.Println("\n=== Calculation History ===")
    for i, entry := range calc.GetHistory() {
        fmt.Printf("%d. %s\n", i+1, entry)
    }
}