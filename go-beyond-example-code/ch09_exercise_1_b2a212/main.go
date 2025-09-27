package main

import "fmt"

func main() {
    cart := make(map[string]int)
    
    // Add items to cart
    cart["apples"] = 3
    cart["bananas"] = 2
    cart["oranges"] = 1
    
    // Display cart
    fmt.Println("Shopping Cart:")
    total := 0
    for item, quantity := range cart {
        fmt.Printf("%s: %d\n", item, quantity)
        total += quantity
    }
    fmt.Printf("Total items: %d\n", total)
    
    // Remove an item
    delete(cart, "bananas")
    fmt.Println("\nAfter removing bananas:")
    for item, quantity := range cart {
        fmt.Printf("%s: %d\n", item, quantity)
    }
}