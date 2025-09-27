package main

import "fmt"

func main() {
    x := 42
    
    // Multiple pointers to the same variable
    p1 := &x
    p2 := &x
    p3 := &x
    
    fmt.Printf("x = %d\n", x)
    fmt.Printf("p1 points to: %d\n", *p1)
    fmt.Printf("p2 points to: %d\n", *p2)
    fmt.Printf("p3 points to: %d\n", *p3)
    
    // All pointers have the same address
    fmt.Printf("p1 == p2: %t\n", p1 == p2)
    fmt.Printf("p2 == p3: %t\n", p2 == p3)
    
    // Modify through any pointer affects all
    *p1 = 100
    fmt.Printf("After modifying through p1:\n")
    fmt.Printf("x = %d\n", x)
    fmt.Printf("p2 points to: %d\n", *p2)
    fmt.Printf("p3 points to: %d\n", *p3)
}