package main

import "fmt"

func main() {
    colors := map[string]string{
        "red":    "#FF0000",
        "green":  "#00FF00",
        "blue":   "#0000FF",
        "yellow": "#FFFF00",
    }
    
    fmt.Printf("Before deletion: %v\n", colors)
    
    // Delete a key
    delete(colors, "yellow")
    fmt.Printf("After deleting yellow: %v\n", colors)
    
    // Try to delete non-existent key (safe)
    delete(colors, "purple")
    fmt.Printf("After trying to delete purple: %v\n", colors)
    
    // Check if key exists before deleting
    if _, exists := colors["green"]; exists {
        delete(colors, "green")
        fmt.Printf("After deleting green: %v\n", colors)
    }
}