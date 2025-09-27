package main

import "fmt"

func main() {
    // Start with empty slice
    numbers := []int{}
    fmt.Printf("Initial: %v, len: %d, cap: %d\n", 
        numbers, len(numbers), cap(numbers))
    
    // Append single element
    numbers = append(numbers, 1)
    fmt.Printf("After append 1: %v, len: %d, cap: %d\n", 
        numbers, len(numbers), cap(numbers))
    
    // Append multiple elements
    numbers = append(numbers, 2, 3, 4)
    fmt.Printf("After append 2,3,4: %v, len: %d, cap: %d\n", 
        numbers, len(numbers), cap(numbers))
    
    // Append another slice
    moreNumbers := []int{5, 6, 7}
    numbers = append(numbers, moreNumbers...)
    fmt.Printf("After append slice: %v, len: %d, cap: %d\n", 
        numbers, len(numbers), cap(numbers))
}