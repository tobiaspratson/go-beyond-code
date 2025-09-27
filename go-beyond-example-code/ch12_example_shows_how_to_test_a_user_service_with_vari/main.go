package main

import (
    "fmt"
    "testing"
)

// User model - represents our domain entity
type User struct {
    ID    int
    Name  string
    Email string
}

// User repository interface - defines the contract for data access
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
}

// User service - contains business logic
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// CreateUser implements the business logic for user creation
func (s *UserService) CreateUser(name, email string) (*User, error) {
    // Business rule: Check if user already exists
    existingUser, _ := s.repo.FindByEmail(email)
    if existingUser != nil {
        return nil, fmt.Errorf("user with email %s already exists", email)
    }
    
    // Business rule: Create new user
    user := &User{
        Name:  name,
        Email: email,
    }
    
    // Business rule: Save user to repository
    err := s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.FindByID(id)
}

// Comprehensive mock user repository with full state management
type MockUserRepository struct {
    users      map[int]*User      // In-memory user storage
    emails     map[string]*User   // Email index for fast lookups
    nextID     int                // Auto-incrementing ID
    saveError  error              // Configurable save error
    findError  error              // Configurable find error
    saveCalls  []*User            // Track all save calls
    findCalls  []int              // Track all find calls by ID
    emailCalls []string           // Track all find calls by email
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users:     make(map[int]*User),
        emails:    make(map[string]*User),
        nextID:    1,
        saveCalls: make([]*User, 0),
        findCalls: make([]int, 0),
        emailCalls: make([]string, 0),
    }
}

// Save method with call tracking
func (m *MockUserRepository) Save(user *User) error {
    // Track the call
    m.saveCalls = append(m.saveCalls, user)
    
    if m.saveError != nil {
        return m.saveError
    }
    
    // Auto-assign ID if not set
    if user.ID == 0 {
        user.ID = m.nextID
        m.nextID++
    }
    
    // Store in both indexes
    m.users[user.ID] = user
    m.emails[user.Email] = user
    return nil
}

// FindByID method with call tracking
func (m *MockUserRepository) FindByID(id int) (*User, error) {
    // Track the call
    m.findCalls = append(m.findCalls, id)
    
    if m.findError != nil {
        return nil, m.findError
    }
    
    user, exists := m.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

// FindByEmail method with call tracking
func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
    // Track the call
    m.emailCalls = append(m.emailCalls, email)
    
    if m.findError != nil {
        return nil, m.findError
    }
    
    user, exists := m.emails[email]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

// Mock control methods for test setup
func (m *MockUserRepository) SetSaveError(err error) {
    m.saveError = err
}

func (m *MockUserRepository) SetFindError(err error) {
    m.findError = err
}

// Verification methods for test assertions
func (m *MockUserRepository) GetSaveCalls() []*User {
    return m.saveCalls
}

func (m *MockUserRepository) GetFindCalls() []int {
    return m.findCalls
}

func (m *MockUserRepository) GetEmailCalls() []string {
    return m.emailCalls
}

func (m *MockUserRepository) WasSaveCalledWith(user *User) bool {
    for _, call := range m.saveCalls {
        if call.Name == user.Name && call.Email == user.Email {
            return true
        }
    }
    return false
}

func (m *MockUserRepository) WasFindByEmailCalledWith(email string) bool {
    for _, call := range m.emailCalls {
        if call == email {
            return true
        }
    }
    return false
}

func (m *MockUserRepository) Clear() {
    m.users = make(map[int]*User)
    m.emails = make(map[string]*User)
    m.nextID = 1
    m.saveError = nil
    m.findError = nil
    m.saveCalls = make([]*User, 0)
    m.findCalls = make([]int, 0)
    m.emailCalls = make([]string, 0)
}

// Test functions demonstrating comprehensive testing scenarios

func TestCreateUser_Success(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Act
    user, err := userService.CreateUser("Alice", "alice@example.com")
    
    // Assert
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
        t.Error("Expected user ID to be set")
    }
    
    // Verify repository interactions
    if len(mockRepo.GetSaveCalls()) != 1 {
        t.Errorf("Expected 1 save call, got %d", len(mockRepo.GetSaveCalls()))
    }
    
    if !mockRepo.WasFindByEmailCalledWith("alice@example.com") {
        t.Error("Expected FindByEmail to be called with alice@example.com")
    }
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Create first user
    _, err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        t.Fatalf("Failed to create first user: %v", err)
    }
    
    // Act - Try to create user with same email
    _, err = userService.CreateUser("Bob", "alice@example.com")
    
    // Assert
    if err == nil {
        t.Error("Expected error for duplicate email")
    }
    
    if err.Error() != "user with email alice@example.com already exists" {
        t.Errorf("Expected specific error message, got %v", err)
    }
    
    // Verify that save was not called for the duplicate
    saveCalls := mockRepo.GetSaveCalls()
    if len(saveCalls) != 1 {
        t.Errorf("Expected 1 save call, got %d", len(saveCalls))
    }
}

func TestCreateUser_SaveError(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Configure mock to return error
    mockRepo.SetSaveError(fmt.Errorf("database connection failed"))
    
    // Act
    _, err := userService.CreateUser("Alice", "alice@example.com")
    
    // Assert
    if err == nil {
        t.Error("Expected error")
    }
    
    if err.Error() != "database connection failed" {
        t.Errorf("Expected database error, got %v", err)
    }
}

func TestCreateUser_FindByEmailError(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Configure mock to return error on FindByEmail
    mockRepo.SetFindError(fmt.Errorf("database query failed"))
    
    // Act
    _, err := userService.CreateUser("Alice", "alice@example.com")
    
    // Assert
    if err == nil {
        t.Error("Expected error")
    }
    
    if err.Error() != "database query failed" {
        t.Errorf("Expected database query error, got %v", err)
    }
}

func TestGetUser_Success(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Pre-populate mock with a user
    testUser := &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
    mockRepo.Save(testUser)
    
    // Act
    user, err := userService.GetUser(1)
    
    // Assert
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user.ID != 1 {
        t.Errorf("Expected ID 1, got %d", user.ID)
    }
    
    if user.Name != "Alice" {
        t.Errorf("Expected name Alice, got %s", user.Name)
    }
}

func TestGetUser_NotFound(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Act
    _, err := userService.GetUser(999)
    
    // Assert
    if err == nil {
        t.Error("Expected error for non-existent user")
    }
    
    if err.Error() != "user not found" {
        t.Errorf("Expected 'user not found' error, got %v", err)
    }
}

// Integration test demonstrating multiple operations
func TestUserService_Integration(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    // Act - Create multiple users
    user1, err1 := userService.CreateUser("Alice", "alice@example.com")
    user2, err2 := userService.CreateUser("Bob", "bob@example.com")
    
    // Assert
    if err1 != nil {
        t.Errorf("Failed to create user1: %v", err1)
    }
    if err2 != nil {
        t.Errorf("Failed to create user2: %v", err2)
    }
    
    if user1.ID == user2.ID {
        t.Error("Expected different IDs for different users")
    }
    
    // Verify all users were saved
    saveCalls := mockRepo.GetSaveCalls()
    if len(saveCalls) != 2 {
        t.Errorf("Expected 2 save calls, got %d", len(saveCalls))
    }
    
    // Verify email lookups were performed
    emailCalls := mockRepo.GetEmailCalls()
    if len(emailCalls) != 2 {
        t.Errorf("Expected 2 email lookups, got %d", len(emailCalls))
    }
}

func main() {
    // Run tests
    fmt.Println("Running comprehensive user service tests...")
    
    // Test 1: Create user
    mockRepo := NewMockUserRepository()
    userService := NewUserService(mockRepo)
    
    user, err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("✓ Created user: %+v\n", user)
    }
    
    // Test 2: Duplicate email
    _, err = userService.CreateUser("Bob", "alice@example.com")
    if err != nil {
        fmt.Printf("✓ Expected error: %v\n", err)
    }
    
    // Test 3: Get user
    retrievedUser, err := userService.GetUser(user.ID)
    if err != nil {
        fmt.Printf("Error retrieving user: %v\n", err)
    } else {
        fmt.Printf("✓ Retrieved user: %+v\n", retrievedUser)
    }
    
    fmt.Println("\n=== Mock Verification ===")
    fmt.Printf("Save calls: %d\n", len(mockRepo.GetSaveCalls()))
    fmt.Printf("Email lookup calls: %d\n", len(mockRepo.GetEmailCalls()))
    fmt.Printf("Find calls: %d\n", len(mockRepo.GetFindCalls()))
}