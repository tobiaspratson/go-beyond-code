package main

import "fmt"

func main() {
    // Pattern 1: Nil slice
    var nilSlice []int
    fmt.Printf("Nil slice: %v, len: %d, cap: %d, is nil: %t\n", 
        nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
    
    // Pattern 2: Empty slice (not nil)
    emptySlice := []int{}
    fmt.Printf("Empty slice: %v, len: %d, cap: %d, is nil: %t\n", 
        emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
    
    // Pattern 3: Pre-allocated slice
    preAllocated := make([]int, 0, 10)  // length 0, capacity 10
    fmt.Printf("Pre-allocated: %v, len: %d, cap: %d\n", 
        preAllocated, len(preAllocated), cap(preAllocated))
    
    // Pattern 4: Slice with initial values
    initialized := make([]int, 5, 10)  // length 5, capacity 10
    fmt.Printf("Initialized: %v, len: %d, cap: %d\n", 
        initialized, len(initialized), cap(initialized))
    
    // Pattern 5: Slice from another slice
    original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    subSlice := original[2:7]  // [3 4 5 6 7]
    fmt.Printf("Sub-slice: %v, len: %d, cap: %d\n", 
        subSlice, len(subSlice), cap(subSlice))
}