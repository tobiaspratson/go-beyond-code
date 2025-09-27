package main

import (
	"fmt"
)

// Advanced DI container with factory support
type AdvancedContainer struct {
	services   map[string]interface{}
	factories  map[string]func() interface{}
	singletons map[string]interface{}
}

func NewAdvancedContainer() *AdvancedContainer {
	return &AdvancedContainer{
		services:   make(map[string]interface{}),
		factories:  make(map[string]func() interface{}),
		singletons: make(map[string]interface{}),
	}
}

// Register a singleton service
func (c *AdvancedContainer) RegisterSingleton(name string, factory func() interface{}) {
	c.factories[name] = factory
}

// Register a transient service
func (c *AdvancedContainer) RegisterTransient(name string, factory func() interface{}) {
	c.factories[name] = factory
}

// Get a service, creating it if necessary
func (c *AdvancedContainer) Get(name string) interface{} {
	// Check if already created singleton
	if singleton, exists := c.singletons[name]; exists {
		return singleton
	}

	// Check if factory exists
	if factory, exists := c.factories[name]; exists {
		service := factory()

		// Store as singleton for future use
		c.singletons[name] = service
		return service
	}

	panic(fmt.Sprintf("Service %s not found", name))
}

// DatabaseConfig
type DatabaseConfig struct {
	Host string
	Port int
}

// DatabaseService
type DatabaseService struct {
	config *DatabaseConfig
}

// NewDatabaseService
func NewDatabaseService(config *DatabaseConfig) *DatabaseService {
	return &DatabaseService{config: config}
}

// UserService
type UserService struct {
	db *DatabaseService
}

// NewUserService
func NewUserService(db *DatabaseService) *UserService {
	return &UserService{db: db}
}

// User
type User struct {
	ID    int
	Name  string
	Email string
}

// GetUsers
func (u *UserService) GetUsers() User {
	// u.db.Connect()
	fmt.Println("Getting users from database")
	return User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
}

// Example usage
func main() {
	container := NewAdvancedContainer()

	// Register services with factories
	container.RegisterSingleton("config", func() interface{} {
		return &DatabaseConfig{
			Host: "localhost",
			Port: 5432,
		}
	})

	container.RegisterSingleton("database", func() interface{} {
		config := container.Get("config").(*DatabaseConfig)
		return NewDatabaseService(config)
	})

	container.RegisterSingleton("userService", func() interface{} {
		db := container.Get("database").(*DatabaseService)
		return NewUserService(db)
	})

	// Use the service
	userService := container.Get("userService").(*UserService)
	userService.GetUsers()
}
