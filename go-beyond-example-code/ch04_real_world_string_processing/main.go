package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    // Email validation example
    email := "user@example.com"
    
    // Basic email validation
    if strings.Contains(email, "@") && strings.Contains(email, ".") {
        fmt.Printf("'%s' looks like a valid email\n", email)
    }
    
    // Text processing example
    text := "  Hello, World! This is a test.  "
    
    // Clean and process text
    cleaned := strings.TrimSpace(text)
    words := strings.Fields(cleaned)
    
    fmt.Printf("Original: '%s'\n", text)
    fmt.Printf("Cleaned: '%s'\n", cleaned)
    fmt.Printf("Words: %v\n", words)
    fmt.Printf("Word count: %d\n", len(words))
    
    // Case conversion
    fmt.Printf("Title case: '%s'\n", strings.Title(cleaned))
    fmt.Printf("Sentence case: '%s'\n", strings.ToLower(cleaned))
    
    // Password strength checker
    password := "MyP@ssw0rd123"
    
    var hasUpper, hasLower, hasDigit, hasSpecial bool
    
    for _, r := range password {
        if unicode.IsUpper(r) {
            hasUpper = true
        } else if unicode.IsLower(r) {
            hasLower = true
        } else if unicode.IsDigit(r) {
            hasDigit = true
        } else if unicode.IsPunct(r) || unicode.IsSymbol(r) {
            hasSpecial = true
        }
    }
    
    fmt.Printf("Password: '%s'\n", password)
    fmt.Printf("Has uppercase: %t\n", hasUpper)
    fmt.Printf("Has lowercase: %t\n", hasLower)
    fmt.Printf("Has digit: %t\n", hasDigit)
    fmt.Printf("Has special: %t\n", hasSpecial)
}