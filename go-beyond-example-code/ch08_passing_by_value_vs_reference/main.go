package main

import "fmt"

// Function that takes a value (creates a copy)
func modifyValue(x int) {
    x = 100
    fmt.Printf("Inside modifyValue: x = %d\n", x)
}

// Function that takes a pointer (modifies original)
func modifyPointer(x *int) {
    *x = 100
    fmt.Printf("Inside modifyPointer: *x = %d\n", *x)
}

func main() {
    num := 42
    
    fmt.Printf("Original: num = %d\n", num)
    
    // Pass by value
    modifyValue(num)
    fmt.Printf("After modifyValue: num = %d\n", num)
    
    // Pass by reference
    modifyPointer(&num)
    fmt.Printf("After modifyPointer: num = %d\n", num)
}