package main

import "fmt"

func main() {
    fmt.Println("=== Basic Continue ===")
    // Skip even numbers
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue  // Skip to next iteration
        }
        fmt.Printf("Odd number: %d\n", i)
    }
    
    fmt.Println("\n=== Continue with Range ===")
    // Skip numbers divisible by 3
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
    for _, num := range numbers {
        if num%3 == 0 {
            fmt.Printf("Skipping %d (divisible by 3)\n", num)
            continue
        }
        fmt.Printf("Processing: %d\n", num)
    }
    
    fmt.Println("\n=== Continue in Nested Loops ===")
    // Continue only affects the innermost loop
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 5; j++ {
            if j == 3 {
                fmt.Printf("  Skipping j=%d\n", j)
                continue  // Skip to next j iteration
            }
            fmt.Printf("  Inner loop: %d\n", j)
        }
    }
}