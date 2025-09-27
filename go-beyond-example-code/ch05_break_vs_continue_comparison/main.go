package main

import "fmt"

func main() {
    fmt.Println("=== Break vs Continue Comparison ===")
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Using BREAK (stops at first even number):")
    for _, num := range numbers {
        if num%2 == 0 {
            fmt.Printf("Found even number %d, breaking\n", num)
            break
        }
        fmt.Printf("Processing odd: %d\n", num)
    }
    
    fmt.Println("\nUsing CONTINUE (skips all even numbers):")
    for _, num := range numbers {
        if num%2 == 0 {
            fmt.Printf("Skipping even number %d\n", num)
            continue
        }
        fmt.Printf("Processing odd: %d\n", num)
    }
    
    fmt.Println("\n=== When to Use Each ===")
    fmt.Println("BREAK: Use when you want to stop processing entirely")
    fmt.Println("CONTINUE: Use when you want to skip some items but keep processing")
}