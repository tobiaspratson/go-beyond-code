package main

import (
	"errors"
	"fmt"
	"log"
)

// Pattern 1: Fail fast
func validateInput(input string) error {
	if input == "" {
		return errors.New("input cannot be empty")
	}
	if len(input) < 3 {
		return errors.New("input too short")
	}
	return nil
}

// Pattern 2: Accumulate errors
func validateUser(user map[string]string) []error {
	var errs []error

	if user["name"] == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if user["email"] == "" {
		errs = append(errs, errors.New("email is required"))
	}
	if user["age"] == "" {
		errs = append(errs, errors.New("age is required"))
	}

	return errs
}

// Pattern 3: Error with recovery
func riskyOperation() (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("operation panicked: %v", r)
		}
	}()

	// Simulate risky operation
	panic("something went wrong")
}

// Pattern 4: Retry with exponential backoff
func retryOperation(operation func() error, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := operation()
		if err == nil {
			return nil
		}

		if i == maxRetries-1 {
			return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
		}

		fmt.Printf("Attempt %d failed: %v, retrying...\n", i+1, err)
	}
	return nil
}

func main() {
	// Test fail fast
	err := validateInput("")
	if err != nil {
		log.Printf("Validation failed: %v", err)
	}

	// Test accumulate errors
	user := map[string]string{
		"name":  "",
		"email": "test@example.com",
		"age":   "",
	}

	errs := validateUser(user)
	if len(errs) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errs {
			fmt.Printf("- %v\n", err)
		}
	}

	// Test recovery
	result, err := riskyOperation()
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	// Test retry
	attempt := 0
	err = retryOperation(func() error {
		attempt++
		if attempt < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}, 5)

	if err != nil {
		fmt.Printf("Retry failed: %v\n", err)
	} else {
		fmt.Println("Operation succeeded after retries")
	}
}
