package main

import (
    "errors"
    "fmt"
    "time"
)

// Error with additional context
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
    Timestamp time.Time
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database error in %s on table %s at %v: %v", 
        e.Operation, e.Table, e.Timestamp, e.Err)
}

func (e DatabaseError) Unwrap() error {
    return e.Err
}

// Network error with retry information
type NetworkError struct {
    URL     string
    Code    int
    Message string
    Retries int
}

func (e NetworkError) Error() string {
    return fmt.Sprintf("network error %d for %s (retries: %d): %s", 
        e.Code, e.URL, e.Retries, e.Message)
}

// Business logic error
type BusinessError struct {
    Code    string
    Message string
    Details map[string]interface{}
}

func (e BusinessError) Error() string {
    return fmt.Sprintf("business error %s: %s", e.Code, e.Message)
}

func (e BusinessError) GetDetails() map[string]interface{} {
    return e.Details
}

// Simulate database operation
func saveUser(name string) error {
    if name == "" {
        return DatabaseError{
            Operation: "INSERT",
            Table:     "users",
            Err:       errors.New("name cannot be empty"),
            Timestamp: time.Now(),
        }
    }
    
    // Simulate database error
    if name == "error" {
        return DatabaseError{
            Operation: "INSERT",
            Table:     "users",
            Err:       errors.New("connection timeout"),
            Timestamp: time.Now(),
        }
    }
    
    fmt.Printf("User %s saved successfully\n", name)
    return nil
}

// Simulate network operation
func fetchData(url string) (string, error) {
    if url == "" {
        return "", NetworkError{
            URL:     url,
            Code:    400,
            Message: "URL cannot be empty",
            Retries: 0,
        }
    }
    
    // Simulate network error
    if url == "error" {
        return "", NetworkError{
            URL:     url,
            Code:    500,
            Message: "server error",
            Retries: 3,
        }
    }
    
    return fmt.Sprintf("Data from %s", url), nil
}

// Simulate business logic
func processOrder(amount float64) error {
    if amount <= 0 {
        return BusinessError{
            Code:    "INVALID_AMOUNT",
            Message: "Order amount must be positive",
            Details: map[string]interface{}{
                "amount": amount,
                "min_amount": 1.0,
            },
        }
    }
    
    if amount > 10000 {
        return BusinessError{
            Code:    "AMOUNT_TOO_LARGE",
            Message: "Order amount exceeds limit",
            Details: map[string]interface{}{
                "amount": amount,
                "max_amount": 10000.0,
            },
        }
    }
    
    fmt.Printf("Order processed for amount: $%.2f\n", amount)
    return nil
}

func main() {
    // Test database error
    err := saveUser("")
    if err != nil {
        fmt.Printf("Database error: %v\n", err)
        
        // Type assertion to get specific error details
        if dbErr, ok := err.(DatabaseError); ok {
            fmt.Printf("Operation: %s, Table: %s\n", dbErr.Operation, dbErr.Table)
        }
    }
    
    // Test network error
    data, err := fetchData("")
    if err != nil {
        fmt.Printf("Network error: %v\n", err)
        
        if netErr, ok := err.(NetworkError); ok {
            fmt.Printf("URL: %s, Code: %d, Retries: %d\n", 
                netErr.URL, netErr.Code, netErr.Retries)
        }
    } else {
        fmt.Printf("Data: %s\n", data)
    }
    
    // Test business error
    err = processOrder(-100)
    if err != nil {
        fmt.Printf("Business error: %v\n", err)
        
        if bizErr, ok := err.(BusinessError); ok {
            fmt.Printf("Error code: %s\n", bizErr.Code)
            fmt.Printf("Details: %v\n", bizErr.GetDetails())
        }
    }
}