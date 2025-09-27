package main

import "fmt"

func main() {
    // Pattern 1: Pre-allocated slice for better performance
    fmt.Println("=== Pattern 1: Pre-allocated slice ===")
    preAllocated := make([]int, 0, 10)  // length 0, capacity 10
    fmt.Printf("Pre-allocated: %v, len: %d, cap: %d\n", 
        preAllocated, len(preAllocated), cap(preAllocated))
    
    // Append to pre-allocated slice
    for i := 1; i <= 5; i++ {
        preAllocated = append(preAllocated, i)
        fmt.Printf("After append %d: len: %d, cap: %d\n", 
            i, len(preAllocated), cap(preAllocated))
    }
    
    // Pattern 2: Append to nil slice
    fmt.Println("\n=== Pattern 2: Append to nil slice ===")
    var nilSlice []int
    fmt.Printf("Nil slice: %v, is nil: %t\n", nilSlice, nilSlice == nil)
    
    nilSlice = append(nilSlice, 1, 2, 3)
    fmt.Printf("After append: %v, is nil: %t\n", nilSlice, nilSlice == nil)
    
    // Pattern 3: Append with capacity growth
    fmt.Println("\n=== Pattern 3: Capacity growth ===")
    growing := []int{1}
    fmt.Printf("Start: len: %d, cap: %d\n", len(growing), cap(growing))
    
    for i := 2; i <= 10; i++ {
        growing = append(growing, i)
        fmt.Printf("After append %d: len: %d, cap: %d\n", 
            i, len(growing), cap(growing))
    }
    
    // Pattern 4: Append slice to itself (dangerous!)
    fmt.Println("\n=== Pattern 4: Append slice to itself ===")
    original := []int{1, 2, 3}
    fmt.Printf("Original: %v\n", original)
    
    // This can cause issues if not careful
    original = append(original, original...)
    fmt.Printf("After append to self: %v\n", original)
}