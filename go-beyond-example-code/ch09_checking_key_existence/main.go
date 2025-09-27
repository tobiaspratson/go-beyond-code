package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
    }
    
    // Method 1: Check if key exists (comma ok idiom)
    score, exists := scores["Alice"]
    if exists {
        fmt.Printf("Alice's score: %d\n", score)
    } else {
        fmt.Println("Alice not found")
    }
    
    // Method 2: Check for non-existent key
    score, exists = scores["David"]
    if exists {
        fmt.Printf("David's score: %d\n", score)
    } else {
        fmt.Printf("David not found (score would be: %d)\n", score)
    }
    
    // Method 3: Inline check (most common pattern)
    if score, ok := scores["Bob"]; ok {
        fmt.Printf("Bob's score: %d\n", score)
    } else {
        fmt.Println("Bob not found")
    }
    
    // Method 4: Check without using the value
    if _, exists := scores["Charlie"]; exists {
        fmt.Println("Charlie is in the map")
    }
    
    // Method 5: Using a helper function
    if hasScore("Alice", scores) {
        fmt.Println("Alice has a score")
    }
}

// Helper function to check key existence
func hasScore(name string, scores map[string]int) bool {
    _, exists := scores[name]
    return exists
}