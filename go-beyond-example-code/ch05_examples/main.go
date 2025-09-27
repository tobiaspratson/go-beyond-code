package main

import "fmt"

func main() {
    fmt.Println("=== Break with Search ===")
    // Search for a specific value
    numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
    target := 50
    found := false
    
    for i, num := range numbers {
        if num == target {
            fmt.Printf("Found %d at index %d\n", target, i)
            found = true
            break
        }
    }
    
    if !found {
        fmt.Printf("%d not found in the slice\n", target)
    }
    
    fmt.Println("\n=== Break with Multiple Conditions ===")
    // Break when multiple conditions are met
    sum := 0
    for i := 1; i <= 20; i++ {
        sum += i
        fmt.Printf("Adding %d, sum is now %d\n", i, sum)
        
        if sum > 50 || i > 10 {
            fmt.Printf("Breaking: sum=%d, i=%d\n", sum, i)
            break
        }
    }
    
    fmt.Println("\n=== Break with Error Handling ===")
    // Break on error condition
    data := []int{1, 2, -3, 4, 5, -6, 7, 8}
    for i, value := range data {
        if value < 0 {
            fmt.Printf("Error: negative value %d found at index %d\n", value, i)
            break
        }
        fmt.Printf("Processing positive value: %d\n", value)
    }
}