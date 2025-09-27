package main

import "fmt"

func main() {
    fmt.Println("=== Basic Break ===")
    // Break out of loop
    for i := 1; i <= 10; i++ {
        if i == 5 {
            fmt.Println("Breaking at 5!")
            break
        }
        fmt.Printf("Count: %d\n", i)
    }
    
    fmt.Println("\n=== Break in While-style Loop ===")
    // Break in while-style loop
    num := 1
    for num < 100 {
        if num > 20 {
            fmt.Println("Breaking at", num)
            break
        }
        fmt.Printf("Number: %d\n", num)
        num *= 2
    }
    
    fmt.Println("\n=== Break with Range ===")
    // Break with range loop
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i, num := range numbers {
        if num > 5 {
            fmt.Printf("Breaking at index %d, value %d\n", i, num)
            break
        }
        fmt.Printf("Processing: %d\n", num)
    }
    
    fmt.Println("\n=== Break in Nested Loops ===")
    // Break only exits the innermost loop
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 5; j++ {
            if j == 3 {
                fmt.Printf("  Breaking inner loop at j=%d\n", j)
                break  // Only breaks the inner loop
            }
            fmt.Printf("  Inner loop: %d\n", j)
        }
    }
}