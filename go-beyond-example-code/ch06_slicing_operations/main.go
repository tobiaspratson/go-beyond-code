package main

import "fmt"

func main() {
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // Basic slicing: [start:end]
    slice1 := numbers[2:5]  // [2 3 4]
    fmt.Printf("numbers[2:5]: %v\n", slice1)
    
    // From beginning: [:end]
    slice2 := numbers[:5]  // [0 1 2 3 4]
    fmt.Printf("numbers[:5]: %v\n", slice2)
    
    // To end: [start:]
    slice3 := numbers[5:]  // [5 6 7 8 9]
    fmt.Printf("numbers[5:]: %v\n", slice3)
    
    // Full slice: [:]
    slice4 := numbers[:]  // [0 1 2 3 4 5 6 7 8 9]
    fmt.Printf("numbers[:]: %v\n", slice4)
    
    // Step by 2: [start:end:step] (Go 1.2+)
    // Note: Go doesn't have step syntax like Python
    // You need to use a loop for stepping
}