package main

import "fmt"

func main() {
    data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    fmt.Println("=== Advanced Slicing Patterns ===")
    
    // Pattern 1: Three-index slicing [start:end:capacity]
    fmt.Println("Pattern 1: Three-index slicing")
    slice1 := data[2:6:8]  // start:2, end:6, capacity:8
    fmt.Printf("data[2:6:8]: %v, len: %d, cap: %d\n", 
        slice1, len(slice1), cap(slice1))
    
    // Pattern 2: Slice bounds demonstration
    fmt.Println("\nPattern 2: Slice bounds")
    fmt.Printf("Original: %v, len: %d, cap: %d\n", 
        data, len(data), cap(data))
    
    // Different slice ranges
    slices := []struct {
        name string
        slice []int
    }{
        {"data[1:3]", data[1:3]},
        {"data[3:7]", data[3:7]},
        {"data[5:]", data[5:]},
        {"data[:4]", data[:4]},
    }
    
    for _, s := range slices {
        fmt.Printf("%s: %v, len: %d, cap: %d\n", 
            s.name, s.slice, len(s.slice), cap(s.slice))
    }
    
    // Pattern 3: Demonstrating shared underlying array
    fmt.Println("\nPattern 3: Shared underlying array")
    original := []int{10, 20, 30, 40, 50}
    sliceA := original[1:4]  // [20 30 40]
    sliceB := original[2:5]  // [30 40 50]
    
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Slice A [1:4]: %v\n", sliceA)
    fmt.Printf("Slice B [2:5]: %v\n", sliceB)
    
    // Modify slice A
    sliceA[0] = 999
    fmt.Printf("After sliceA[0] = 999:\n")
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Slice A: %v\n", sliceA)
    fmt.Printf("Slice B: %v\n", sliceB)  // Also affected!
}