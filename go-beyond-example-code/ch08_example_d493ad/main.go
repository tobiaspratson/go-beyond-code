package main

import (
	"errors"
	"fmt"
)

// Use panic for programming errors
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero - programming error")
	}
	return a / b
}

// Use error for expected conditions
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Use panic for impossible conditions
func getElement(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		panic("index out of bounds - programming error")
	}
	return slice[index]
}

// Use error for user input validation
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > 150 {
		return errors.New("age seems unrealistic")
	}
	return nil
}

// Use panic for initialization failures
func initializeDatabase() {
	// Simulate initialization failure
	panic("database connection failed - cannot continue")
}

// Use error for recoverable failures
func connectToDatabase() error {
	// Simulate connection failure
	return errors.New("database connection failed")
}

func main() {
	// Examples of when to use panic
	fmt.Println("=== Panic Examples ===")

	// Programming error - use panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from programming error: %v\n", r)
		}
	}()

	// This is a programming error - should never happen
	result := divide(10, 0)
	fmt.Printf("Result: %f\n", result)

	// Examples of when to use error
	fmt.Println("\n=== Error Examples ===")

	// User input - use error
	err := validateAge(-5)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	// Recoverable failure - use error
	err = connectToDatabase()
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
	}

	// Safe operation - use error
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}
