package main

import "fmt"

// Logger interface - for logging operations
type Logger interface {
    Log(message string)
}

// Email service interface - for sending emails
type EmailService interface {
    SendEmail(to, subject, body string)
}

// User service with multiple dependencies
type UserService struct {
    logger       Logger        // For logging operations
    emailService EmailService  // For sending emails
    userRepo     *UserRepository // For data persistence
}

// Constructor with multiple dependencies
// All dependencies are injected at construction time
func NewUserService(logger Logger, emailService EmailService, userRepo *UserRepository) *UserService {
    return &UserService{
        logger:       logger,
        emailService: emailService,
        userRepo:     userRepo,
    }
}

func (u *UserService) CreateUser(name, email string) error {
    // Use injected logger
    u.logger.Log(fmt.Sprintf("Creating user: %s", name))
    
    // Use injected repository to save user
    err := u.userRepo.SaveUser(name)
    if err != nil {
        u.logger.Log(fmt.Sprintf("Failed to save user %s: %v", name, err))
        return err
    }
    
    // Use injected email service to send welcome email
    u.emailService.SendEmail(email, "Welcome!", "Welcome to our service!")
    
    u.logger.Log(fmt.Sprintf("User %s created successfully", name))
    return nil
}

// Simple logger implementation
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
    fmt.Printf("[LOG] %s\n", message)
}

// Simple email service implementation
type ConsoleEmailService struct{}

func (c *ConsoleEmailService) SendEmail(to, subject, body string) {
    fmt.Printf("[EMAIL] To: %s, Subject: %s, Body: %s\n", to, subject, body)
}

func main() {
    // Create all dependencies
    logger := &ConsoleLogger{}
    emailService := &ConsoleEmailService{}
    userRepo := NewUserRepository(&MySQLDatabase{})
    
    // Inject all dependencies through constructor
    userService := NewUserService(logger, emailService, userRepo)
    
    // Use the service
    userService.CreateUser("Alice", "alice@example.com")
}