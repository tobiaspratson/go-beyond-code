package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}
    
    fmt.Println("=== Index and Value ===")
    // Range with index and value
    for i, num := range numbers {
        fmt.Printf("Index %d: %d\n", i, num)
    }
    
    fmt.Println("\n=== Index Only ===")
    // Range with index only
    for i := range numbers {
        fmt.Printf("Index: %d\n", i)
    }
    
    fmt.Println("\n=== Value Only ===")
    // Range with value only
    for _, num := range numbers {
        fmt.Printf("Value: %d\n", num)
    }
    
    // Working with different slice types
    fmt.Println("\n=== String Slice ===")
    fruits := []string{"apple", "banana", "cherry", "date"}
    for i, fruit := range fruits {
        fmt.Printf("%d: %s (length: %d)\n", i, fruit, len(fruit))
    }
    
    fmt.Println("\n=== Float Slice ===")
    prices := []float64{1.99, 2.49, 3.99, 4.50}
    total := 0.0
    for i, price := range prices {
        total += price
        fmt.Printf("Item %d: $%.2f\n", i+1, price)
    }
    fmt.Printf("Total: $%.2f\n", total)
}