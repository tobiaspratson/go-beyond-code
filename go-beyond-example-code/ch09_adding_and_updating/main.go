package main

import "fmt"

func main() {
    // Start with empty map
    inventory := make(map[string]int)
    
    // Add items
    inventory["apples"] = 50
    inventory["bananas"] = 30
    inventory["oranges"] = 25
    
    fmt.Printf("Initial inventory: %v\n", inventory)
    
    // Update existing item
    inventory["apples"] = 45
    fmt.Printf("After updating apples: %v\n", inventory)
    
    // Add new item
    inventory["grapes"] = 20
    fmt.Printf("After adding grapes: %v\n", inventory)
    
    // Update multiple items
    inventory["bananas"] = 35
    inventory["oranges"] = 30
    fmt.Printf("Final inventory: %v\n", inventory)
}