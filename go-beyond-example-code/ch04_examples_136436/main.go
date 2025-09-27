package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    text := "  Hello, World!  "
    
    // Basic string operations
    fmt.Printf("Original: '%s'\n", text)
    fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(text))
    fmt.Printf("Uppercase: '%s'\n", strings.ToUpper(text))
    fmt.Printf("Lowercase: '%s'\n", strings.ToLower(text))
    
    // String splitting and joining
    words := strings.Fields("apple banana cherry")
    fmt.Printf("Words: %v\n", words)
    fmt.Printf("Joined: '%s'\n", strings.Join(words, ", "))
    
    // String replacement
    original := "Hello World"
    replaced := strings.ReplaceAll(original, "World", "Go")
    fmt.Printf("Replaced: '%s'\n", replaced)
    
    // String contains and prefix/suffix
    fmt.Printf("Contains 'World': %t\n", strings.Contains(original, "World"))
    fmt.Printf("Starts with 'Hello': %t\n", strings.HasPrefix(original, "Hello"))
    fmt.Printf("Ends with 'World': %t\n", strings.HasSuffix(original, "World"))
}