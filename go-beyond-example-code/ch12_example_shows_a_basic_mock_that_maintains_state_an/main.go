package main

import "fmt"

// Database interface - this is what we want to mock
type Database interface {
    Save(data string) error
    Find(id string) (string, error)
    Delete(id string) error
}

// Mock database with state - implements the Database interface
type MockDatabase struct {
    data        map[string]string  // In-memory storage to simulate database
    saveError   error             // Configurable error for Save method
    findError   error             // Configurable error for Find method
    deleteError error             // Configurable error for Delete method
}

// Constructor for mock database
func NewMockDatabase() *MockDatabase {
    return &MockDatabase{
        data: make(map[string]string),
    }
}

// Save method - simulates database save operation
func (m *MockDatabase) Save(data string) error {
    // Check if we should return an error
    if m.saveError != nil {
        return m.saveError
    }
    // Store data in our in-memory map
    m.data[data] = data
    return nil
}

// Find method - simulates database find operation
func (m *MockDatabase) Find(id string) (string, error) {
    // Check if we should return an error
    if m.findError != nil {
        return "", m.findError
    }
    // Look up data in our in-memory map
    data, exists := m.data[id]
    if !exists {
        return "", fmt.Errorf("not found")
    }
    return data, nil
}

// Delete method - simulates database delete operation
func (m *MockDatabase) Delete(id string) error {
    // Check if we should return an error
    if m.deleteError != nil {
        return m.deleteError
    }
    // Remove data from our in-memory map
    delete(m.data, id)
    return nil
}

// Methods to control mock behavior - these are the key to effective mocking
func (m *MockDatabase) SetSaveError(err error) {
    m.saveError = err
}

func (m *MockDatabase) SetFindError(err error) {
    m.findError = err
}

func (m *MockDatabase) SetDeleteError(err error) {
    m.deleteError = err
}

// Utility methods for testing and verification
func (m *MockDatabase) GetData() map[string]string {
    return m.data
}

func (m *MockDatabase) Clear() {
    m.data = make(map[string]string)
    m.saveError = nil
    m.findError = nil
    m.deleteError = nil
}

// Service that uses database - this is what we want to test
type DataService struct {
    db Database
}

func NewDataService(db Database) *DataService {
    return &DataService{db: db}
}

func (s *DataService) StoreData(data string) error {
    return s.db.Save(data)
}

func (s *DataService) RetrieveData(id string) (string, error) {
    return s.db.Find(id)
}

func (s *DataService) RemoveData(id string) error {
    return s.db.Delete(id)
}

func main() {
    // Create mock database
    mockDB := NewMockDatabase()
    dataService := NewDataService(mockDB)
    
    fmt.Println("=== Testing Successful Operations ===")
    
    // Test successful operations
    err := dataService.StoreData("test-data")
    if err != nil {
        fmt.Printf("Error storing data: %v\n", err)
    } else {
        fmt.Println("✓ Data stored successfully")
    }
    
    data, err := dataService.RetrieveData("test-data")
    if err != nil {
        fmt.Printf("Error retrieving data: %v\n", err)
    } else {
        fmt.Printf("✓ Retrieved data: %s\n", data)
    }
    
    fmt.Println("\n=== Testing Error Conditions ===")
    
    // Test error conditions
    mockDB.SetSaveError(fmt.Errorf("database connection failed"))
    err = dataService.StoreData("another-data")
    if err != nil {
        fmt.Printf("✓ Expected error: %v\n", err)
    }
    
    // Reset mock for next test
    mockDB.Clear()
    mockDB.SetFindError(fmt.Errorf("query timeout"))
    _, err = dataService.RetrieveData("some-id")
    if err != nil {
        fmt.Printf("✓ Expected find error: %v\n", err)
    }
    
    fmt.Println("\n=== Mock State Verification ===")
    fmt.Printf("Mock database contains %d items\n", len(mockDB.GetData()))
}