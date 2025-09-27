package main

import "fmt"

func main() {
    // Regular variable
    x := 42
    fmt.Printf("Value of x: %d\n", x)
    fmt.Printf("Address of x: %p\n", &x)
    
    // Pointer variable
    var p *int = &x
    fmt.Printf("Value of p (address): %p\n", p)
    fmt.Printf("Value that p points to: %d\n", *p)
    
    // Modify value through pointer
    *p = 100
    fmt.Printf("New value of x: %d\n", x)
    
    // Show that both x and *p are the same
    fmt.Printf("x == *p: %t\n", x == *p)
    fmt.Printf("&x == p: %t\n", &x == p)
}