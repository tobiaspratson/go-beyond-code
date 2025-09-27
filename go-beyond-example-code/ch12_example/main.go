package main

import "fmt"

// Interface for email service
type EmailService interface {
    SendEmail(to, subject, body string) error
}

// Real email service
type SMTPEmailService struct{}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
    fmt.Printf("Sending email via SMTP to %s: %s\n", to, subject)
    // Simulate network delay
    return nil
}

// Mock email service for testing
type MockEmailService struct {
    sentEmails []Email
}

type Email struct {
    To      string
    Subject string
    Body    string
}

func NewMockEmailService() *MockEmailService {
    return &MockEmailService{
        sentEmails: make([]Email, 0),
    }
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
    // Record the email instead of sending it
    m.sentEmails = append(m.sentEmails, Email{
        To:      to,
        Subject: subject,
        Body:    body,
    })
    fmt.Printf("Mock: Email recorded for %s: %s\n", to, subject)
    return nil
}

func (m *MockEmailService) GetSentEmails() []Email {
    return m.sentEmails
}

func (m *MockEmailService) Clear() {
    m.sentEmails = make([]Email, 0)
}

// User service that uses email service
type UserService struct {
    emailService EmailService
}

func NewUserService(emailService EmailService) *UserService {
    return &UserService{
        emailService: emailService,
    }
}

func (u *UserService) CreateUser(name, email string) error {
    // Simulate user creation
    fmt.Printf("Creating user: %s\n", name)
    
    // Send welcome email
    return u.emailService.SendEmail(email, "Welcome!", "Welcome to our service!")
}

func main() {
    // Use real service in production
    realEmailService := &SMTPEmailService{}
    userService := NewUserService(realEmailService)
    userService.CreateUser("Alice", "alice@example.com")
    
    // Use mock service in tests
    mockEmailService := NewMockEmailService()
    userService = NewUserService(mockEmailService)
    userService.CreateUser("Bob", "bob@example.com")
    
    // Verify emails were sent
    sentEmails := mockEmailService.GetSentEmails()
    fmt.Printf("Mock recorded %d emails\n", len(sentEmails))
    for _, email := range sentEmails {
        fmt.Printf("  To: %s, Subject: %s\n", email.To, email.Subject)
    }
}