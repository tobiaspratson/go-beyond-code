package main

import (
    "fmt"
    "time"
)

// Database interface
type Database interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
    DeleteUser(id int) error
    GetUsers() ([]*User, error)
}

// User model
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

// Service that depends on database
type UserService struct {
    db Database
}

func NewUserService(db Database) *UserService {
    return &UserService{db: db}
}

func (us *UserService) GetUserProfile(id int) (*User, error) {
    return us.db.GetUser(id)
}

func (us *UserService) CreateUser(name, email string, age int) (*User, error) {
    user := &User{
        Name:  name,
        Email: email,
        Age:   age,
    }
    
    err := us.db.SaveUser(user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (us *UserService) UpdateUser(id int, name, email string, age int) (*User, error) {
    user, err := us.db.GetUser(id)
    if err != nil {
        return nil, err
    }
    
    user.Name = name
    user.Email = email
    user.Age = age
    
    err = us.db.SaveUser(user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (us *UserService) GetAllUsers() ([]*User, error) {
    return us.db.GetUsers()
}

// Mock database for testing
type MockDatabase struct {
    users map[int]*User
    nextID int
    shouldError bool
    errorMessage string
}

func NewMockDatabase() *MockDatabase {
    return &MockDatabase{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

func (md *MockDatabase) GetUser(id int) (*User, error) {
    if md.shouldError {
        return nil, fmt.Errorf(md.errorMessage)
    }
    
    user, exists := md.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (md *MockDatabase) SaveUser(user *User) error {
    if md.shouldError {
        return fmt.Errorf(md.errorMessage)
    }
    
    if user.ID == 0 {
        user.ID = md.nextID
        md.nextID++
    }
    md.users[user.ID] = user
    return nil
}

func (md *MockDatabase) DeleteUser(id int) error {
    if md.shouldError {
        return fmt.Errorf(md.errorMessage)
    }
    
    delete(md.users, id)
    return nil
}

func (md *MockDatabase) GetUsers() ([]*User, error) {
    if md.shouldError {
        return nil, fmt.Errorf(md.errorMessage)
    }
    
    users := make([]*User, 0, len(md.users))
    for _, user := range md.users {
        users = append(users, user)
    }
    return users, nil
}

// Mock configuration methods
func (md *MockDatabase) SetError(shouldError bool, message string) {
    md.shouldError = shouldError
    md.errorMessage = message
}

func (md *MockDatabase) GetUserCount() int {
    return len(md.users)
}

func (md *MockDatabase) Clear() {
    md.users = make(map[int]*User)
    md.nextID = 1
    md.shouldError = false
    md.errorMessage = ""
}

// Test functions
func TestUserService_CreateUser(t *testing.T) {
    mockDB := NewMockDatabase()
    service := NewUserService(mockDB)
    
    // Test successful creation
    user, err := service.CreateUser("Alice", "alice@example.com", 25)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user.Name != "Alice" {
        t.Errorf("Expected name 'Alice', got %s", user.Name)
    }
    
    if user.ID == 0 {
        t.Error("Expected user ID to be set")
    }
    
    if mockDB.GetUserCount() != 1 {
        t.Errorf("Expected 1 user in database, got %d", mockDB.GetUserCount())
    }
}

func TestUserService_GetUserProfile(t *testing.T) {
    mockDB := NewMockDatabase()
    service := NewUserService(mockDB)
    
    // Add a user to the mock database
    testUser := &User{ID: 1, Name: "Bob", Email: "bob@example.com", Age: 30}
    mockDB.SaveUser(testUser)
    
    // Test getting the user
    user, err := service.GetUserProfile(1)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user.Name != "Bob" {
        t.Errorf("Expected name 'Bob', got %s", user.Name)
    }
}

func TestUserService_ErrorHandling(t *testing.T) {
    mockDB := NewMockDatabase()
    service := NewUserService(mockDB)
    
    // Configure mock to return error
    mockDB.SetError(true, "database connection failed")
    
    // Test that errors are properly propagated
    _, err := service.CreateUser("Alice", "alice@example.com", 25)
    if err == nil {
        t.Error("Expected error, got nil")
    }
    
    if err.Error() != "database connection failed" {
        t.Errorf("Expected 'database connection failed', got %s", err.Error())
    }
}

func main() {
    fmt.Println("=== Mock Testing Example ===")
    
    // Create mock database
    mockDB := NewMockDatabase()
    
    // Create service with mock
    service := NewUserService(mockDB)
    
    // Test creating a user
    user, err := service.CreateUser("Alice", "alice@example.com", 25)
    if err != nil {
        fmt.Printf("Error creating user: %v\n", err)
        return
    }
    
    fmt.Printf("Created user: %+v\n", user)
    
    // Test getting the user
    retrievedUser, err := service.GetUserProfile(user.ID)
    if err != nil {
        fmt.Printf("Error getting user: %v\n", err)
        return
    }
    
    fmt.Printf("Retrieved user: %+v\n", retrievedUser)
    
    // Test error handling
    mockDB.SetError(true, "simulated database error")
    _, err = service.CreateUser("Bob", "bob@example.com", 30)
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }
}