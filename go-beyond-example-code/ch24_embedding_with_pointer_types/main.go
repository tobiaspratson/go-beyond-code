package main

import "fmt"

// Base type
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) GetCount() int {
    return c.count
}

func (c *Counter) Reset() {
    c.count = 0
}

// Embedding pointer type
type Service struct {
    *Counter  // Embedded pointer - methods promoted
    Name      string
}

func NewService(name string) *Service {
    return &Service{
        Counter: &Counter{},  // Initialize the embedded counter
        Name:    name,
    }
}

func (s *Service) GetStatus() string {
    return fmt.Sprintf("Service %s: %d operations", s.Name, s.GetCount())
}

func main() {
    service := NewService("DatabaseService")
    
    fmt.Println("=== Pointer Embedding ===")
    fmt.Println("Initial status:", service.GetStatus())
    
    // Use embedded methods
    service.Increment()
    service.Increment()
    service.Increment()
    
    fmt.Println("After operations:", service.GetStatus())
    
    // Reset and check
    service.Reset()
    fmt.Println("After reset:", service.GetStatus())
}