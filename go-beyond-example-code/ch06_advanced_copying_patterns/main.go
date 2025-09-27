package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("=== Advanced Copying Patterns ===")
    
    // Pattern 1: Partial copy
    fmt.Println("Pattern 1: Partial copy")
    partial := make([]int, 3)
    copy(partial, original[2:5])  // Copy elements 2,3,4
    fmt.Printf("Partial copy: %v\n", partial)
    
    // Pattern 2: Copy with different sizes
    fmt.Println("\nPattern 2: Copy with different sizes")
    smaller := make([]int, 3)
    larger := make([]int, 8)
    
    copy(smaller, original)  // Only copies first 3 elements
    copy(larger, original)   // Copies all 10 elements
    
    fmt.Printf("Smaller copy: %v\n", smaller)
    fmt.Printf("Larger copy: %v\n", larger)
    
    // Pattern 3: Copy to existing slice
    fmt.Println("\nPattern 3: Copy to existing slice")
    existing := []int{99, 99, 99, 99, 99}
    copy(existing, original[1:4])  // Overwrites first 3 elements
    fmt.Printf("Existing after copy: %v\n", existing)
    
    // Pattern 4: Deep copy demonstration
    fmt.Println("\nPattern 4: Deep copy vs shallow copy")
    deepCopy := make([]int, len(original))
    copy(deepCopy, original)
    
    shallowCopy := original[:]  // This is a slice, not a copy!
    
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Deep copy: %v\n", deepCopy)
    fmt.Printf("Shallow copy: %v\n", shallowCopy)
    
    // Modify original
    original[0] = 999
    fmt.Printf("\nAfter original[0] = 999:\n")
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Deep copy: %v\n", deepCopy)      // Unchanged
    fmt.Printf("Shallow copy: %v\n", shallowCopy)  // Changed!
}