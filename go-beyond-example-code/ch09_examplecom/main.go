package main

import "fmt"

func main() {
    original := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    // Method 1: Manual copy (shallow copy)
    copy1 := make(map[string]int)
    for k, v := range original {
        copy1[k] = v
    }
    
    // Method 2: Using a helper function
    copy2 := copyMap(original)
    
    // Method 3: Copy with capacity hint
    copy3 := make(map[string]int, len(original))
    for k, v := range original {
        copy3[k] = v
    }
    
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Copy 1: %v\n", copy1)
    fmt.Printf("Copy 2: %v\n", copy2)
    fmt.Printf("Copy 3: %v\n", copy3)
    
    // Modify original
    original["d"] = 4
    fmt.Printf("\nAfter modifying original:\n")
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Copy 1: %v (unchanged)\n", copy1)
    fmt.Printf("Copy 2: %v (unchanged)\n", copy2)
    fmt.Printf("Copy 3: %v (unchanged)\n", copy3)
}

// Helper function for copying maps
func copyMap(original map[string]int) map[string]int {
    copy := make(map[string]int, len(original))
    for k, v := range original {
        copy[k] = v
    }
    return copy
}