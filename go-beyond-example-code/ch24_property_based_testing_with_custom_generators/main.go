package main

import (
    "fmt"
    "math/rand"
    "reflect"
    "testing"
    "testing/quick"
    "time"
)

// Custom generator for email addresses
func emailGenerator() string {
    domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "example.com"}
    names := []string{"alice", "bob", "charlie", "david", "eve"}
    
    name := names[rand.Intn(len(names))]
    domain := domains[rand.Intn(len(domains))]
    return fmt.Sprintf("%s@%s", name, domain)
}

// Custom generator for phone numbers
func phoneGenerator() string {
    return fmt.Sprintf("+1-%d-%d-%d", 
        rand.Intn(900)+100, 
        rand.Intn(900)+100, 
        rand.Intn(9000)+1000)
}

// Custom generator for user data
func userGenerator() User {
    names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
    emails := []string{"alice@example.com", "bob@test.com", "charlie@demo.org"}
    
    return User{
        ID:    rand.Intn(1000) + 1,
        Name:  names[rand.Intn(len(names))],
        Email: emails[rand.Intn(len(emails))],
        Age:   rand.Intn(50) + 18,
    }
}

// User struct for testing
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

// Function to test
func ValidateEmail(email string) bool {
    return len(email) > 0 && 
           len(email) < 100 && 
           contains(email, "@") && 
           contains(email, ".")
}

func ValidatePhone(phone string) bool {
    return len(phone) > 0 && 
           len(phone) < 20 && 
           contains(phone, "+")
}

func ValidateUser(user User) bool {
    return user.ID > 0 && 
           len(user.Name) > 0 && 
           ValidateEmail(user.Email) && 
           user.Age >= 18
}

func contains(s, substr string) bool {
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}

func TestEmailValidation(t *testing.T) {
    property := func(email string) bool {
        if ValidateEmail(email) {
            return len(email) > 0 && contains(email, "@")
        }
        return true // Invalid emails are acceptable
    }
    
    config := &quick.Config{
        Values: func(args []reflect.Value, rand *rand.Rand) {
            args[0] = reflect.ValueOf(emailGenerator())
        },
    }
    
    if err := quick.Check(property, config); err != nil {
        t.Error(err)
    }
}

func TestPhoneValidation(t *testing.T) {
    property := func(phone string) bool {
        if ValidatePhone(phone) {
            return len(phone) > 0 && contains(phone, "+")
        }
        return true // Invalid phones are acceptable
    }
    
    config := &quick.Config{
        Values: func(args []reflect.Value, rand *rand.Rand) {
            args[0] = reflect.ValueOf(phoneGenerator())
        },
    }
    
    if err := quick.Check(property, config); err != nil {
        t.Error(err)
    }
}

func TestUserValidation(t *testing.T) {
    property := func(user User) bool {
        if ValidateUser(user) {
            return user.ID > 0 && len(user.Name) > 0 && user.Age >= 18
        }
        return true // Invalid users are acceptable
    }
    
    config := &quick.Config{
        Values: func(args []reflect.Value, rand *rand.Rand) {
            args[0] = reflect.ValueOf(userGenerator())
        },
    }
    
    if err := quick.Check(property, config); err != nil {
        t.Error(err)
    }
}

func main() {
    // Test the functions
    emails := []string{"alice@gmail.com", "bob@yahoo.com", "invalid"}
    for _, email := range emails {
        fmt.Printf("Email '%s' is valid: %t\n", email, ValidateEmail(email))
    }
    
    phones := []string{"+1-555-123-4567", "+1-800-555-0199", "invalid"}
    for _, phone := range phones {
        fmt.Printf("Phone '%s' is valid: %t\n", phone, ValidatePhone(phone))
    }
}