package main

import (
    "fmt"
    "strings"
)

func main() {
    // Inefficient string concatenation
    var result string
    for i := 0; i < 5; i++ {
        result += fmt.Sprintf("Item %d, ", i)  // Creates new string each time!
    }
    fmt.Printf("Inefficient: %s\n", result)
    
    // Efficient string building with strings.Builder
    var builder strings.Builder
    for i := 0; i < 5; i++ {
        builder.WriteString(fmt.Sprintf("Item %d, ", i))
    }
    efficient := builder.String()
    fmt.Printf("Efficient: %s\n", efficient)
    
    // Using strings.Join for known elements
    items := []string{"Apple", "Banana", "Cherry"}
    joined := strings.Join(items, ", ")
    fmt.Printf("Joined: %s\n", joined)
}