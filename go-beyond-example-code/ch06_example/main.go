package main

import "fmt"

func main() {
    // Create an array
    scores := [5]int{85, 92, 78, 96, 88}
    
    // Access elements
    fmt.Printf("First score: %d\n", scores[0])  // 85
    fmt.Printf("Last score: %d\n", scores[4])   // 88
    
    // Modify elements
    scores[2] = 90
    fmt.Printf("Updated scores: %v\n", scores)  // [85 92 90 96 88]
    
    // Calculate average
    sum := 0
    for i := 0; i < len(scores); i++ {
        sum += scores[i]
    }
    average := float64(sum) / float64(len(scores))
    fmt.Printf("Average score: %.2f\n", average)
    
    // Using range
    total := 0
    for _, score := range scores {
        total += score
    }
    fmt.Printf("Total: %d\n", total)
}