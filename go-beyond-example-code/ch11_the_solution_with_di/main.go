package main

import (
    "fmt"
    "time"
)

// Define interface for email service
// This is the key: we depend on an interface, not a concrete type
type EmailSender interface {
    SendEmail(to, subject, body string)
}

// Email service implementation
type EmailService struct{}

func (e *EmailService) SendEmail(to, subject, body string) {
    fmt.Printf("Sending email to %s: %s\n", to, subject)
    time.Sleep(100 * time.Millisecond)
}

// User service with injected dependency
type UserService struct {
    emailSender EmailSender  // Depends on interface, not concrete type
}

// Constructor that injects dependencies
// This is where the "injection" happens
func NewUserService(emailSender EmailSender) *UserService {
    return &UserService{
        emailSender: emailSender,
    }
}

func (u *UserService) CreateUser(name, email string) {
    fmt.Printf("Creating user: %s (%s)\n", name, email)
    
    // Use injected dependency
    // UserService doesn't know or care about the concrete implementation
    u.emailSender.SendEmail(email, "Welcome!", "Welcome to our service!")
}

func main() {
    // Create dependencies
    emailService := &EmailService{}
    
    // Inject dependencies through constructor
    userService := NewUserService(emailService)
    userService.CreateUser("Alice", "alice@example.com")
}