package main

import (
    "context"
    "fmt"
    "time"
)

type Database struct {
    connectionPool chan bool
}

func NewDatabase() *Database {
    // Simulate connection pool
    pool := make(chan bool, 3)
    for i := 0; i < 3; i++ {
        pool <- true
    }
    return &Database{connectionPool: pool}
}

func (db *Database) Query(ctx context.Context, query string) (string, error) {
    // Acquire connection
    select {
    case <-db.connectionPool:
        defer func() { db.connectionPool <- true }()
    case <-ctx.Done():
        return "", fmt.Errorf("connection timeout: %v", ctx.Err())
    }
    
    // Simulate query execution
    select {
    case <-time.After(500 * time.Millisecond):
        return fmt.Sprintf("Result for query: %s", query), nil
    case <-ctx.Done():
        return "", fmt.Errorf("query cancelled: %v", ctx.Err())
    }
}

func (db *Database) Transaction(ctx context.Context, operations []string) error {
    fmt.Println("Starting transaction")
    
    for i, op := range operations {
        select {
        case <-ctx.Done():
            return fmt.Errorf("transaction cancelled at operation %d: %v", i, ctx.Err())
        default:
            fmt.Printf("Executing operation %d: %s\n", i+1, op)
            time.Sleep(200 * time.Millisecond)
        }
    }
    
    fmt.Println("Transaction completed")
    return nil
}

func main() {
    db := NewDatabase()
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Perform query
    result, err := db.Query(ctx, "SELECT * FROM users")
    if err != nil {
        fmt.Printf("Query error: %v\n", err)
        return
    }
    fmt.Printf("Query result: %s\n", result)
    
    // Perform transaction
    operations := []string{"INSERT user", "UPDATE profile", "INSERT preferences"}
    err = db.Transaction(ctx, operations)
    if err != nil {
        fmt.Printf("Transaction error: %v\n", err)
        return
    }
}