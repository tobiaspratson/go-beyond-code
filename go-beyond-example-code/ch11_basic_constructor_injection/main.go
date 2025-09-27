package main

import "fmt"

// Database interface - defines what we need, not how it's implemented
type Database interface {
    Save(data string) error
    Find(id string) (string, error)
}

// User repository with injected database dependency
type UserRepository struct {
    db Database  // Depends on interface, not concrete type
}

// Constructor with dependency injection
// This is the key: we inject the dependency through the constructor
func NewUserRepository(db Database) *UserRepository {
    return &UserRepository{
        db: db,  // Store the injected dependency
    }
}

func (r *UserRepository) SaveUser(name string) error {
    // Use the injected database
    return r.db.Save(fmt.Sprintf("user:%s", name))
}

func (r *UserRepository) FindUser(id string) (string, error) {
    // Use the injected database
    return r.db.Find(id)
}

// Database implementations
type MySQLDatabase struct{}

func (m *MySQLDatabase) Save(data string) error {
    fmt.Printf("MySQL: Saving %s\n", data)
    return nil
}

func (m *MySQLDatabase) Find(id string) (string, error) {
    fmt.Printf("MySQL: Finding %s\n", id)
    return fmt.Sprintf("data for %s", id), nil
}

type PostgreSQLDatabase struct{}

func (p *PostgreSQLDatabase) Save(data string) error {
    fmt.Printf("PostgreSQL: Saving %s\n", data)
    return nil
}

func (p *PostgreSQLDatabase) Find(id string) (string, error) {
    fmt.Printf("PostgreSQL: Finding %s\n", id)
    return fmt.Sprintf("data for %s", id), nil
}

func main() {
    // Use MySQL - inject MySQL implementation
    mysqlDB := &MySQLDatabase{}
    userRepo1 := NewUserRepository(mysqlDB)
    userRepo1.SaveUser("Alice")
    
    // Use PostgreSQL - inject PostgreSQL implementation
    postgresDB := &PostgreSQLDatabase{}
    userRepo2 := NewUserRepository(postgresDB)
    userRepo2.SaveUser("Bob")
    
    // The beauty: UserRepository doesn't know or care which database it's using!
}