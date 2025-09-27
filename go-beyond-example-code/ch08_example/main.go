package main

import "fmt"

func main() {
    // Different types of pointers
    var intPtr *int
    var strPtr *string
    var floatPtr *float64
    
    // Initialize variables
    num := 42
    text := "Hello"
    price := 19.99
    
    // Assign addresses to pointers
    intPtr = &num
    strPtr = &text
    floatPtr = &price
    
    fmt.Printf("num: %d, address: %p\n", num, &num)
    fmt.Printf("intPtr points to: %d\n", *intPtr)
    
    fmt.Printf("text: %s, address: %p\n", text, &text)
    fmt.Printf("strPtr points to: %s\n", *strPtr)
    
    fmt.Printf("price: %.2f, address: %p\n", price, &price)
    fmt.Printf("floatPtr points to: %.2f\n", *floatPtr)
    
    // Modify values through pointers
    *intPtr = 100
    *strPtr = "World"
    *floatPtr = 29.99
    
    fmt.Printf("After modification:\n")
    fmt.Printf("num: %d\n", num)
    fmt.Printf("text: %s\n", text)
    fmt.Printf("price: %.2f\n", price)
}