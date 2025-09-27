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
}

// User model
type User struct {
    ID    int
    Name  string
    Email string
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

func (us *UserService) CreateUser(name, email string) (*User, error) {
    user := &User{
        Name:  name,
        Email: email,
    }
    
    err := us.db.SaveUser(user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

// Mock database for testing
type MockDatabase struct {
    users map[int]*User
    nextID int
}

func NewMockDatabase() *MockDatabase {
    return &MockDatabase{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

func (md *MockDatabase) GetUser(id int) (*User, error) {
    user, exists := md.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (md *MockDatabase) SaveUser(user *User) error {
    if user.ID == 0 {
        user.ID = md.nextID
        md.nextID++
    }
    md.users[user.ID] = user
    return nil
}

func (md *MockDatabase) DeleteUser(id int) error {
    delete(md.users, id)
    return nil
}

func main() {
    // Create mock database
    mockDB := NewMockDatabase()
    
    // Create service with mock
    service := NewUserService(mockDB)
    
    // Test creating a user
    user, err := service.CreateUser("Alice", "alice@example.com")
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
}