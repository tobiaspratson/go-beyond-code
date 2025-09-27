package main

import (
    "fmt"
    "reflect"
)

// Simple DI container
type Container struct {
    services map[string]interface{}
}

func NewContainer() *Container {
    return &Container{
        services: make(map[string]interface{}),
    }
}

func (c *Container) Register(name string, service interface{}) {
    c.services[name] = service
}

func (c *Container) Get(name string) interface{} {
    service, exists := c.services[name]
    if !exists {
        panic(fmt.Sprintf("Service %s not found", name))
    }
    return service
}

// Service definitions
type DatabaseConfig struct {
    Host string
    Port int
}

type DatabaseService struct {
    config *DatabaseConfig
}

func NewDatabaseService(config *DatabaseConfig) *DatabaseService {
    return &DatabaseService{config: config}
}

func (d *DatabaseService) Connect() {
    fmt.Printf("Connecting to database at %s:%d\n", d.config.Host, d.config.Port)
}

type UserService struct {
    db *DatabaseService
}

func NewUserService(db *DatabaseService) *UserService {
    return &UserService{db: db}
}

func (u *UserService) GetUsers() {
    fmt.Println("Getting users from database")
    u.db.Connect()
}

func main() {
    // Create container
    container := NewContainer()
    
    // Register services in dependency order
    container.Register("config", &DatabaseConfig{
        Host: "localhost",
        Port: 5432,
    })
    
    // Create database service with injected config
    config := container.Get("config").(*DatabaseConfig)
    dbService := NewDatabaseService(config)
    container.Register("database", dbService)
    
    // Create user service with injected database
    db := container.Get("database").(*DatabaseService)
    userService := NewUserService(db)
    container.Register("userService", userService)
    
    // Use the service
    userService = container.Get("userService").(*UserService)
    userService.GetUsers()
}