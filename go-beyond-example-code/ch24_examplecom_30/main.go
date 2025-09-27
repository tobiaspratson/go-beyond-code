package main

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

// Function to fuzz test
func ProcessString(input string) (string, error) {
    if len(input) == 0 {
        return "", fmt.Errorf("empty input")
    }
    
    if len(input) > 1000 {
        return "", fmt.Errorf("input too long")
    }
    
    // Simulate some processing
    result := ""
    for _, char := range input {
        if char >= 'a' && char <= 'z' {
            result += string(char - 32) // Convert to uppercase
        } else if char >= 'A' && char <= 'Z' {
            result += string(char + 32) // Convert to lowercase
        } else {
            result += string(char) // Keep as is
        }
    }
    
    return result, nil
}

// Fuzz test function
func FuzzProcessString(f *testing.F) {
    // Add seed corpus
    f.Add("hello")
    f.Add("WORLD")
    f.Add("Hello World!")
    f.Add("")
    f.Add("a")
    
    f.Fuzz(func(t *testing.T, input string) {
        result, err := ProcessString(input)
        
        // Test properties that should always hold
        if err == nil {
            if len(result) != len(input) {
                t.Errorf("Result length %d != input length %d", len(result), len(input))
            }
        } else {
            // If there's an error, it should be for empty or too long input
            if len(input) > 0 && len(input) <= 1000 {
                t.Errorf("Unexpected error for valid input: %v", err)
            }
        }
    })
}

// Custom fuzz test with specific generators
func TestProcessStringWithCustomFuzz(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    for i := 0; i < 1000; i++ {
        // Generate random string
        length := rand.Intn(100) + 1
        input := make([]byte, length)
        for j := range input {
            input[j] = byte(rand.Intn(256))
        }
        
        result, err := ProcessString(string(input))
        
        // Test properties
        if err == nil {
            if len(result) != len(input) {
                t.Errorf("Result length %d != input length %d", len(result), len(input))
            }
        }
    }
}

func main() {
    fmt.Println("=== Fuzz Testing Example ===")
    
    // Test various inputs
    testCases := []string{
        "hello",
        "WORLD",
        "Hello World!",
        "",
        "a",
        "This is a very long string that might cause issues if not handled properly",
    }
    
    for _, input := range testCases {
        result, err := ProcessString(input)
        if err != nil {
            fmt.Printf("Input: '%s' -> Error: %v\n", input, err)
        } else {
            fmt.Printf("Input: '%s' -> Output: '%s'\n", input, result)
        }
    }
}