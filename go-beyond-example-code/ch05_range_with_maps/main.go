package main

import "fmt"

func main() {
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    
    fmt.Println("=== Key and Value ===")
    // Range over map
    for color, hex := range colors {
        fmt.Printf("%s: %s\n", color, hex)
    }
    
    fmt.Println("\n=== Key Only ===")
    // Range with key only
    for color := range colors {
        fmt.Printf("Color: %s\n", color)
    }
    
    fmt.Println("\n=== Value Only ===")
    // Range with value only
    for _, hex := range colors {
        fmt.Printf("Hex code: %s\n", hex)
    }
    
    // Working with different map types
    fmt.Println("\n=== Integer Map ===")
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    
    total := 0
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
        total += score
    }
    fmt.Printf("Average score: %.2f\n", float64(total)/float64(len(scores)))
}