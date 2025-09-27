package main

import (
    "context"
    "fmt"
    "time"
)

type Database struct {
    // Simulate database connection
}

func (db *Database) Query(ctx context.Context, query string) (string, error) {
    // Simulate database query
    select {
    case <-time.After(500 * time.Millisecond):
        return fmt.Sprintf("Result for query: %s", query), nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func (db *Database) Insert(ctx context.Context, data string) error {
    // Simulate database insert
    select {
    case <-time.After(300 * time.Millisecond):
        fmt.Printf("Inserted: %s\n", data)
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    db := &Database{}
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Perform database operations
    result, err := db.Query(ctx, "SELECT * FROM users")
    if err != nil {
        fmt.Printf("Query error: %v\n", err)
        return
    }
    fmt.Printf("Query result: %s\n", result)
    
    err = db.Insert(ctx, "New user data")
    if err != nil {
        fmt.Printf("Insert error: %v\n", err)
        return
    }
}