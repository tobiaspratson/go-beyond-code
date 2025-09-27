package main

import (
    "fmt"
    "time"
)

// Tip 1: Use capacity hints
func efficientAppend() {
    fmt.Println("=== Efficient Append Pattern ===")
    
    // Inefficient: multiple reallocations
    start := time.Now()
    var inefficient []int
    for i := 0; i < 10000; i++ {
        inefficient = append(inefficient, i)
    }
    fmt.Printf("Inefficient: %v, reallocations: %d\n", 
        time.Since(start), countReallocations(inefficient))
    
    // Efficient: pre-allocate capacity
    start = time.Now()
    efficient := make([]int, 0, 10000)
    for i := 0; i < 10000; i++ {
        efficient = append(efficient, i)
    }
    fmt.Printf("Efficient: %v, reallocations: %d\n", 
        time.Since(start), countReallocations(efficient))
}

func countReallocations(slice []int) int {
    // This is a simplified example - in reality, you'd need to track allocations
    return 0
}

// Tip 2: Avoid unnecessary copies
func avoidUnnecessaryCopies() {
    fmt.Println("\n=== Avoiding Unnecessary Copies ===")
    
    original := []int{1, 2, 3, 4, 5}
    
    // Bad: creates unnecessary copy
    badCopy := make([]int, len(original))
    copy(badCopy, original)
    badCopy = append(badCopy, 6)
    
    // Good: work with original slice
    goodCopy := append(original, 6)
    
    fmt.Printf("Bad approach: %v\n", badCopy)
    fmt.Printf("Good approach: %v\n", goodCopy)
}

// Tip 3: Use slice operations efficiently
func efficientSliceOperations() {
    fmt.Println("\n=== Efficient Slice Operations ===")
    
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    
    // Efficient: use slice expressions
    start := time.Now()
    result1 := data[100:200]  // No copy, just reference
    fmt.Printf("Slice expression: %v, len=%d\n", 
        time.Since(start), len(result1))
    
    // Less efficient: copy the slice
    start = time.Now()
    result2 := make([]int, 100)
    copy(result2, data[100:200])
    fmt.Printf("Copy operation: %v, len=%d\n", 
        time.Since(start), len(result2))
}

// Tip 4: Batch operations
func batchOperations() {
    fmt.Println("\n=== Batch Operations ===")
    
    // Inefficient: multiple appends
    start := time.Now()
    var inefficient []int
    for i := 0; i < 1000; i++ {
        inefficient = append(inefficient, i)
    }
    fmt.Printf("Multiple appends: %v\n", time.Since(start))
    
    // Efficient: batch append
    start = time.Now()
    batch := make([]int, 0, 1000)
    for i := 0; i < 1000; i += 100 {
        batch = append(batch, make([]int, 100)...)
        for j := 0; j < 100; j++ {
            batch[len(batch)-100+j] = i + j
        }
    }
    fmt.Printf("Batch operations: %v\n", time.Since(start))
}

func main() {
    efficientAppend()
    avoidUnnecessaryCopies()
    efficientSliceOperations()
    batchOperations()
}