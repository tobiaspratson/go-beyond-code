package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
        "David":   78,
    }
    
    // Iterate over key-value pairs
    fmt.Println("All scores:")
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
    
    // Iterate over keys only
    fmt.Println("\nAll names:")
    for name := range scores {
        fmt.Printf("Name: %s\n", name)
    }
    
    // Iterate over values only
    fmt.Println("\nAll scores:")
    for _, score := range scores {
        fmt.Printf("Score: %d\n", score)
    }
}