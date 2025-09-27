package main

import (
    "fmt"
    "runtime"
)

func demonstrateMemoryManagement() {
    fmt.Println("=== Memory Management Best Practices ===")
    
    // 1. Reuse slices when possible
    fmt.Println("1. Reuse slices")
    var reusable []int
    for i := 0; i < 5; i++ {
        reusable = reusable[:0]  // Reset length, keep capacity
        for j := 0; j < 10; j++ {
            reusable = append(reusable, i*10+j)
        }
        fmt.Printf("Iteration %d: len=%d, cap=%d\n", i, len(reusable), cap(reusable))
    }
    
    // 2. Avoid memory leaks
    fmt.Println("\n2. Avoid memory leaks")
    leakyFunction()
    runtime.GC()  // Force garbage collection
    
    // 3. Use appropriate data structures
    fmt.Println("\n3. Choose appropriate data structures")
    demonstrateDataStructureChoice()
}

func leakyFunction() {
    // This function creates a slice that holds references to large data
    largeData := make([]int, 1000000)
    for i := range largeData {
        largeData[i] = i
    }
    
    // If we return a slice that references this data, it won't be garbage collected
    // This is a potential memory leak
    _ = largeData[100:200]  // This slice keeps the entire largeData in memory
}

func demonstrateDataStructureChoice() {
    // For small, fixed-size collections, arrays might be better
    smallFixed := [5]int{1, 2, 3, 4, 5}
    fmt.Printf("Small fixed array: %v\n", smallFixed)
    
    // For dynamic collections, slices are better
    dynamic := make([]int, 0, 10)
    for i := 0; i < 5; i++ {
        dynamic = append(dynamic, i)
    }
    fmt.Printf("Dynamic slice: %v\n", dynamic)
}

func main() {
    demonstrateMemoryManagement()
}