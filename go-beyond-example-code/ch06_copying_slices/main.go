package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5}
    
    // Method 1: Using make and copy
    copy1 := make([]int, len(original))
    copy(copy1, original)
    fmt.Printf("Copy 1: %v\n", copy1)
    
    // Method 2: Using append
    copy2 := append([]int(nil), original...)
    fmt.Printf("Copy 2: %v\n", copy2)
    
    // Method 3: Using slice expression
    copy3 := original[:]
    fmt.Printf("Copy 3: %v\n", copy3)
    
    // Verify they're different slices
    original[0] = 999
    fmt.Printf("Original after change: %v\n", original)
    fmt.Printf("Copy 1 after change: %v\n", copy1)  // Unchanged
    fmt.Printf("Copy 2 after change: %v\n", copy2)  // Unchanged
    fmt.Printf("Copy 3 after change: %v\n", copy3)  // Changed! (same underlying array)
}