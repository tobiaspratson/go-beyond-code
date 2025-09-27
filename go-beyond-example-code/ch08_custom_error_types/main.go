package main

import (
    "fmt"
    "time"
)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

// Function that returns custom error
func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "age cannot be negative",
        }
    }
    if age > 150 {
        return ValidationError{
            Field:   "age",
            Message: "age cannot be greater than 150",
        }
    }
    return nil
}

func main() {
    // Test with valid age
    err := validateAge(25)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println("Age is valid")
    }
    
    // Test with invalid age
    err = validateAge(-5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Test with another invalid age
    err = validateAge(200)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}