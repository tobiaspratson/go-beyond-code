package main

import "fmt"

func main() {
    // Method 1: Using make (recommended for empty maps)
    scores := make(map[string]int)
    scores["Alice"] = 95
    scores["Bob"] = 87
    scores["Charlie"] = 92
    
    // Method 2: Map literal (recommended for known values)
    ages := map[string]int{
        "Alice":   25,
        "Bob":     30,
        "Charlie": 28,
    }
    
    // Method 3: Empty map literal (alternative to make)
    empty := map[string]int{}
    
    // Method 4: Pre-allocated map with capacity hint
    largeMap := make(map[string]int, 100) // Hint: expect ~100 elements
    
    fmt.Printf("Scores: %v\n", scores)
    fmt.Printf("Ages: %v\n", ages)
    fmt.Printf("Empty: %v\n", empty)
    fmt.Printf("Large map: %v (len: %d)\n", largeMap, len(largeMap))
}