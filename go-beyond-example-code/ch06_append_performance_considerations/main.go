package main

import "fmt"

func demonstrateAppendPerformance() {
    fmt.Println("=== Append Performance Comparison ===")
    
    // Method 1: No pre-allocation (slower)
    fmt.Println("Method 1: No pre-allocation")
    var slow []int
    for i := 0; i < 5; i++ {
        slow = append(slow, i)
        fmt.Printf("  Append %d: len: %d, cap: %d\n", i, len(slow), cap(slow))
    }
    
    // Method 2: Pre-allocation (faster)
    fmt.Println("\nMethod 2: Pre-allocation")
    fast := make([]int, 0, 5)  // Pre-allocate capacity
    for i := 0; i < 5; i++ {
        fast = append(fast, i)
        fmt.Printf("  Append %d: len: %d, cap: %d\n", i, len(fast), cap(fast))
    }
    
    // Method 3: Known size (fastest)
    fmt.Println("\nMethod 3: Known size")
    known := make([]int, 5)  // Pre-allocate length
    for i := 0; i < 5; i++ {
        known[i] = i  // Direct assignment
        fmt.Printf("  Set %d: len: %d, cap: %d\n", i, len(known), cap(known))
    }
    
    fmt.Printf("Results: slow=%v, fast=%v, known=%v\n", slow, fast, known)
}

func main() {
    demonstrateAppendPerformance()
}