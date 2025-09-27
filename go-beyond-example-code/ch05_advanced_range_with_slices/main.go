package main

import "fmt"

func main() {
    // Working with slice modifications during iteration
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Original slice:", numbers)
    
    // Modifying slice during iteration
    fmt.Println("\n=== Modifying During Iteration ===")
    for i, num := range numbers {
        if num%2 == 0 {
            numbers[i] = num * 2  // Double even numbers
        }
        fmt.Printf("Index %d: %d\n", i, num)
    }
    fmt.Println("Modified slice:", numbers)
    
    // Range with slice bounds
    fmt.Println("\n=== Range with Subslice ===")
    data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("Full slice:", data)
    
    // Iterate over a subslice
    subslice := data[2:7]  // elements 2, 3, 4, 5, 6
    fmt.Println("Subslice [2:7]:", subslice)
    for i, val := range subslice {
        fmt.Printf("Subslice[%d] = %d (original index: %d)\n", 
            i, val, i+2)
    }
    
    // Range with empty slice
    fmt.Println("\n=== Empty Slice ===")
    empty := []int{}
    fmt.Println("Empty slice length:", len(empty))
    for i, val := range empty {
        fmt.Printf("This won't print: %d, %d\n", i, val)
    }
    fmt.Println("No iterations for empty slice")
}