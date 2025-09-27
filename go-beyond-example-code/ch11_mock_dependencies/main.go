package main

import (
    "fmt"
    "testing"
)

// Mock user repository for testing
// This implements the UserRepository interface
type MockUserRepository struct {
    users     map[int]*User
    saveError error  // Can be set to simulate save errors
    findError error  // Can be set to simulate find errors
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users: make(map[int]*User),
    }
}

// SetError methods to control mock behavior
func (m *MockUserRepository) SetSaveError(err error) {
    m.saveError = err
}

func (m *MockUserRepository) SetFindError(err error) {
    m.findError = err
}

func (m *MockUserRepository) Save(user *User) error {
    if m.saveError != nil {
        return m.saveError
    }
    m.users[user.ID] = user
    return nil
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    if m.findError != nil {
        return nil, m.findError
    }
    user, exists := m.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
    if m.findError != nil {
        return nil, m.findError
    }
    for _, user := range m.users {
        if user.Email == email {
            return user, nil
        }
    }
    return nil, fmt.Errorf("user not found")
}

// Test function for successful user creation
func TestUserService_CreateUser_Success(t *testing.T) {
    // Create mock repository
    mockRepo := NewMockUserRepository()
    
    // Create service with mock
    userService := NewUserService(mockRepo)
    
    // Test creating a user
    user, err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user.Name != "Alice" {
        t.Errorf("Expected name Alice, got %s", user.Name)
    }
    
    if user.Email != "alice@example.com" {
        t.Errorf("Expected email alice@example.com, got %s", user.Email)
    }
    
    if user.ID == 0 {
        t.Error("Expected user to have an ID")
    }
}

// Test function for duplicate email
func TestUserService_CreateUser_DuplicateEmail(t *testing.T) {
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Create first user
    _, err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        t.Errorf("Expected no error for first user, got %v", err)
    }
    
    // Try to create user with same email
    _, err = userService.CreateUser("Bob", "alice@example.com")
    if err == nil {
        t.Error("Expected error for duplicate email, got nil")
    }
    
    expectedError := "user with email alice@example.com already exists"
    if err.Error() != expectedError {
        t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
    }
}

// Test function for save error
func TestUserService_CreateUser_SaveError(t *testing.T) {
    mockRepo := NewMockUserRepository()
    mockRepo.SetSaveError(fmt.Errorf("database connection failed"))
    
    userService := NewUserService(mockRepo)
    
    // Try to create user - should fail due to save error
    _, err := userService.CreateUser("Alice", "alice@example.com")
    if err == nil {
        t.Error("Expected error for save failure, got nil")
    }
    
    if err.Error() != "database connection failed" {
        t.Errorf("Expected 'database connection failed', got '%s'", err.Error())
    }
}

func main() {
    // Run the test
    fmt.Println("Running test...")
    // In a real test, you would use testing.T
    // TestUserService(&testing.T{})
}