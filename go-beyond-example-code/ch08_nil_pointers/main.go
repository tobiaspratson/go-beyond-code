package main

import "fmt"

func main() {
    // Declare a nil pointer
    var p *int
    fmt.Printf("Nil pointer: %v\n", p)
    
    // Check if pointer is nil
    if p == nil {
        fmt.Println("Pointer is nil")
    }
    
    // DON'T do this - it will panic!
    // fmt.Printf("Value: %d\n", *p) // panic: runtime error: invalid memory address
    
    // Safe way to check and use pointer
    if p != nil {
        fmt.Printf("Value: %d\n", *p)
    } else {
        fmt.Println("Cannot dereference nil pointer")
    }
    
    // Initialize pointer
    x := 42
    p = &x
    fmt.Printf("Now p points to: %d\n", *p)
}