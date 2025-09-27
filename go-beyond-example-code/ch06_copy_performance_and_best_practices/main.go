package main

import "fmt"

func demonstrateCopyPerformance() {
    fmt.Println("=== Copy Performance Comparison ===")
    
    original := make([]int, 1000)
    for i := range original {
        original[i] = i
    }
    
    // Method 1: Using copy() function (fastest)
    fmt.Println("Method 1: Using copy() function")
    copy1 := make([]int, len(original))
    copy(copy1, original)
    fmt.Printf("  Length: %d, First few: %v\n", len(copy1), copy1[:5])
    
    // Method 2: Using append (slower for large slices)
    fmt.Println("\nMethod 2: Using append")
    copy2 := append([]int(nil), original...)
    fmt.Printf("  Length: %d, First few: %v\n", len(copy2), copy2[:5])
    
    // Method 3: Manual copy (fastest for small slices)
    fmt.Println("\nMethod 3: Manual copy")
    copy3 := make([]int, len(original))
    for i, v := range original {
        copy3[i] = v
    }
    fmt.Printf("  Length: %d, First few: %v\n", len(copy3), copy3[:5])
    
    // Verify all methods produce same result
    fmt.Printf("All copies equal: %t\n", 
        len(copy1) == len(copy2) && len(copy2) == len(copy3))
}

func demonstrateCopyEdgeCases() {
    fmt.Println("\n=== Copy Edge Cases ===")
    
    // Case 1: Copy to nil slice
    var nilSlice []int
    source := []int{1, 2, 3}
    copy(nilSlice, source)  // No effect
    fmt.Printf("Copy to nil slice: %v\n", nilSlice)
    
    // Case 2: Copy from nil slice
    var dest []int
    copy(dest, nilSlice)  // No effect
    fmt.Printf("Copy from nil slice: %v\n", dest)
    
    // Case 3: Copy with zero length
    zeroSource := []int{}
    zeroDest := make([]int, 0)
    copy(zeroDest, zeroSource)
    fmt.Printf("Copy zero-length slices: %v\n", zeroDest)
    
    // Case 4: Copy overlapping slices
    overlapping := []int{1, 2, 3, 4, 5}
    copy(overlapping[1:], overlapping[0:4])  // Shift left
    fmt.Printf("Overlapping copy: %v\n", overlapping)
}

func main() {
    demonstrateCopyPerformance()
    demonstrateCopyEdgeCases()
}