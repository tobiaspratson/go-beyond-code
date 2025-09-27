package main

import "fmt"

// User model - represents a user in our system
type User struct {
    ID    int
    Name  string
    Email string
}

// User repository interface - defines what we need from a user repository
// This is the key: we define what we need, not how it's implemented
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
}

// User service interface - defines what our user service can do
type UserService interface {
    CreateUser(name, email string) (*User, error)
    GetUser(id int) (*User, error)
    GetUserByEmail(email string) (*User, error)
}

// Implementation of UserService
// Note: we return the interface, not the concrete type
type userService struct {
    repo UserRepository  // Depends on interface, not concrete type
}

// Constructor that injects the repository dependency
func NewUserService(repo UserRepository) UserService {
    return &userService{
        repo: repo,  // Store the injected dependency
    }
}

func (s *userService) CreateUser(name, email string) (*User, error) {
    // Check if user already exists using injected repository
    existingUser, _ := s.repo.FindByEmail(email)
    if existingUser != nil {
        return nil, fmt.Errorf("user with email %s already exists", email)
    }
    
    // Create new user
    user := &User{
        Name:  name,
        Email: email,
    }
    
    // Save user using injected repository
    err := s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *userService) GetUser(id int) (*User, error) {
    // Use injected repository to find user
    return s.repo.FindByID(id)
}

func (s *userService) GetUserByEmail(email string) (*User, error) {
    // Use injected repository to find user by email
    return s.repo.FindByEmail(email)
}

// In-memory repository implementation
// This implements the UserRepository interface
type InMemoryUserRepository struct {
    users  map[int]*User
    nextID int
}

func NewInMemoryUserRepository() UserRepository {
    return &InMemoryUserRepository{
        users:  make(map[int]*User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepository) Save(user *User) error {
    if user.ID == 0 {
        user.ID = r.nextID
        r.nextID++
    }
    r.users[user.ID] = user
    return nil
}

func (r *InMemoryUserRepository) FindByID(id int) (*User, error) {
    user, exists := r.users[id]
    if !exists {
        return nil, fmt.Errorf("user with ID %d not found", id)
    }
    return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*User, error) {
    for _, user := range r.users {
        if user.Email == email {
            return user, nil
        }
    }
    return nil, fmt.Errorf("user with email %s not found", email)
}

func main() {
    // Create repository - could be any implementation of UserRepository
    repo := NewInMemoryUserRepository()
    
    // Create service with injected repository
    // The service doesn't know it's using an in-memory repository
    userService := NewUserService(repo)
    
    // Use the service
    user, err := userService.CreateUser("Alice", "alice@example.com")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Created user: %+v\n", user)
    
    // Retrieve user
    retrievedUser, err := userService.GetUser(user.ID)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Retrieved user: %+v\n", retrievedUser)
}