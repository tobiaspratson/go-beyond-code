package main

import (
    "errors"
    "fmt"
    "os"
)

// Pattern 1: Simple error return
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Pattern 2: Error with context
func readConfig(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read config file %s: %w", filename, err)
    }
    return string(data), nil
}

// Pattern 3: Multiple error conditions
func validateUser(name, email string) error {
    if name == "" {
        return errors.New("name cannot be empty")
    }
    if email == "" {
        return errors.New("email cannot be empty")
    }
    if len(name) < 2 {
        return errors.New("name must be at least 2 characters")
    }
    return nil
}

func main() {
    // Test division
    result, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Division error: %v\n", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }
    
    // Test division by zero
    result, err = divide(10, 0)
    if err != nil {
        fmt.Printf("Division error: %v\n", err)
    }
    
    // Test config reading
    config, err := readConfig("config.txt")
    if err != nil {
        fmt.Printf("Config error: %v\n", err)
    } else {
        fmt.Printf("Config: %s\n", config)
    }
    
    // Test validation
    err = validateUser("", "test@example.com")
    if err != nil {
        fmt.Printf("Validation error: %v\n", err)
    }
}