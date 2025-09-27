package main

import "fmt"

func main() {
    data := [6]int{10, 20, 30, 40, 50, 60}
    
    // Pattern 1: Traditional for loop with index
    fmt.Println("Pattern 1 - Index and value:")
    for i := 0; i < len(data); i++ {
        fmt.Printf("Index %d: %d\n", i, data[i])
    }
    
    // Pattern 2: Range with index and value
    fmt.Println("\nPattern 2 - Range with index:")
    for i, value := range data {
        fmt.Printf("Index %d: %d\n", i, value)
    }
    
    // Pattern 3: Range with value only
    fmt.Println("\nPattern 3 - Range with value only:")
    for _, value := range data {
        fmt.Printf("Value: %d\n", value)
    }
    
    // Pattern 4: Range with index only
    fmt.Println("\nPattern 4 - Range with index only:")
    for i := range data {
        fmt.Printf("Index %d: %d\n", i, data[i])
    }
    
    // Pattern 5: Reverse iteration
    fmt.Println("\nPattern 5 - Reverse iteration:")
    for i := len(data) - 1; i >= 0; i-- {
        fmt.Printf("Index %d: %d\n", i, data[i])
    }
}