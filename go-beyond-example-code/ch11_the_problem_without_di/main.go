package main

import (
    "fmt"
    "time"
)

// Bad: Hard-coded dependencies
type EmailService struct {
    // Email service implementation
}

func (e *EmailService) SendEmail(to, subject, body string) {
    fmt.Printf("Sending email to %s: %s\n", to, subject)
    // Simulate email sending
    time.Sleep(100 * time.Millisecond)
}

type UserService struct {
    // User service implementation
}

func (u *UserService) CreateUser(name, email string) {
    fmt.Printf("Creating user: %s (%s)\n", name, email)
    
    // Problem: Hard-coded dependency
    // This creates several issues:
    // 1. UserService is tightly coupled to EmailService
    // 2. Can't easily test without sending real emails
    // 3. Can't swap email providers without changing code
    // 4. Violates Single Responsibility Principle
    emailService := &EmailService{}
    emailService.SendEmail(email, "Welcome!", "Welcome to our service!")
}

func main() {
    userService := &UserService{}
    userService.CreateUser("Alice", "alice@example.com")
}