package main

import (
    "fmt"
    "runtime"
    "time"
)

func demonstrateMemoryPatterns() {
    fmt.Println("=== Memory Allocation Patterns ===")
    
    // Pattern 1: Growing slice (inefficient)
    fmt.Println("Pattern 1: Growing slice")
    start := time.Now()
    var growing []int
    for i := 0; i < 1000; i++ {
        growing = append(growing, i)
    }
    fmt.Printf("Time: %v, Final capacity: %d\n", time.Since(start), cap(growing))
    
    // Pattern 2: Pre-allocated slice (efficient)
    fmt.Println("\nPattern 2: Pre-allocated slice")
    start = time.Now()
    preAllocated := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        preAllocated = append(preAllocated, i)
    }
    fmt.Printf("Time: %v, Final capacity: %d\n", time.Since(start), cap(preAllocated))
    
    // Pattern 3: Fixed size (most efficient)
    fmt.Println("\nPattern 3: Fixed size")
    start = time.Now()
    fixed := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        fixed[i] = i
    }
    fmt.Printf("Time: %v, Final capacity: %d\n", time.Since(start), cap(fixed))
}

func demonstrateCapacityGrowth() {
    fmt.Println("\n=== Capacity Growth Pattern ===")
    
    var slice []int
    fmt.Printf("Initial: len=%d, cap=%d\n", len(slice), cap(slice))
    
    for i := 0; i < 20; i++ {
        slice = append(slice, i)
        fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(slice), cap(slice))
    }
}

func demonstrateMemoryUsage() {
    fmt.Println("\n=== Memory Usage Comparison ===")
    
    // Small slice
    small := make([]int, 10)
    fmt.Printf("Small slice: len=%d, cap=%d\n", len(small), cap(small))
    
    // Large slice
    large := make([]int, 1000000)
    fmt.Printf("Large slice: len=%d, cap=%d\n", len(large), cap(large))
    
    // Check memory stats
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory allocated: %d KB\n", m.Alloc/1024)
}

func main() {
    demonstrateMemoryPatterns()
    demonstrateCapacityGrowth()
    demonstrateMemoryUsage()
}