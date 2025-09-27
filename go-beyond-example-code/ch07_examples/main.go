package main

import "fmt"

// Common Go interfaces you'll encounter
type Stringer interface {
    String() string
}

type Error interface {
    Error() string
}

// Custom interfaces for a database system
type Database interface {
    Connect() error
    Disconnect() error
}

type Queryable interface {
    Query(sql string) ([]map[string]interface{}, error)
}

type Writable interface {
    Insert(table string, data map[string]interface{}) error
    Update(table string, id int, data map[string]interface{}) error
    Delete(table string, id int) error
}

// Composed interface for full database functionality
type FullDatabase interface {
    Database
    Queryable
    Writable
}

// Mock database implementation
type MockDB struct {
    connected bool
    data      map[string][]map[string]interface{}
}

func (m *MockDB) Connect() error {
    m.connected = true
    return nil
}

func (m *MockDB) Disconnect() error {
    m.connected = false
    return nil
}

func (m *MockDB) Query(sql string) ([]map[string]interface{}, error) {
    if !m.connected {
        return nil, fmt.Errorf("not connected")
    }
    // Mock query result
    return []map[string]interface{}{
        {"id": 1, "name": "Alice"},
        {"id": 2, "name": "Bob"},
    }, nil
}

func (m *MockDB) Insert(table string, data map[string]interface{}) error {
    if !m.connected {
        return fmt.Errorf("not connected")
    }
    if m.data == nil {
        m.data = make(map[string][]map[string]interface{})
    }
    m.data[table] = append(m.data[table], data)
    return nil
}

func (m *MockDB) Update(table string, id int, data map[string]interface{}) error {
    if !m.connected {
        return fmt.Errorf("not connected")
    }
    // Mock update logic
    return nil
}

func (m *MockDB) Delete(table string, id int) error {
    if !m.connected {
        return fmt.Errorf("not connected")
    }
    // Mock delete logic
    return nil
}

func main() {
    db := &MockDB{}
    
    // Use as Database interface
    var database Database = db
    database.Connect()
    
    // Use as Queryable interface
    var queryable Queryable = db
    results, err := queryable.Query("SELECT * FROM users")
    if err != nil {
        fmt.Printf("Query error: %v\n", err)
    } else {
        fmt.Printf("Query results: %+v\n", results)
    }
    
    // Use as Writable interface
    var writable Writable = db
    writable.Insert("users", map[string]interface{}{
        "name": "Charlie",
        "age":  30,
    })
    
    // Use as FullDatabase interface
    var fullDB FullDatabase = db
    fullDB.Insert("products", map[string]interface{}{
        "name":  "Laptop",
        "price": 999.99,
    })
    
    fullDB.Disconnect()
}