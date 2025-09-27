package main

import "fmt"

func main() {
    // Create an array
    arr := [5]int{10, 20, 30, 40, 50}
    fmt.Printf("Original array: %v\n", arr)
    
    // Create slices from the array
    slice1 := arr[1:4]  // [20 30 40]
    slice2 := arr[0:3]   // [10 20 30]
    slice3 := arr[2:]    // [30 40 50]
    
    fmt.Printf("Slice1 [1:4]: %v, len: %d, cap: %d\n", 
        slice1, len(slice1), cap(slice1))
    fmt.Printf("Slice2 [0:3]: %v, len: %d, cap: %d\n", 
        slice2, len(slice2), cap(slice2))
    fmt.Printf("Slice3 [2:]: %v, len: %d, cap: %d\n", 
        slice3, len(slice3), cap(slice3))
    
    // Demonstrate shared underlying array
    fmt.Println("\n--- Demonstrating shared underlying array ---")
    slice1[0] = 999  // This changes the underlying array
    fmt.Printf("After changing slice1[0] to 999:\n")
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Slice1: %v\n", slice1)
    fmt.Printf("Slice2: %v\n", slice2)  // Also affected!
    fmt.Printf("Slice3: %v\n", slice3)  // Also affected!
}