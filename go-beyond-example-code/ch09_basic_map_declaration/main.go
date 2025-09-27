package main

import "fmt"

func main() {
    // Declare a map (this creates a nil map)
    var colors map[string]string
    fmt.Printf("Empty map: %v\n", colors)
    fmt.Printf("Is nil: %t\n", colors == nil)
    
    // Initialize with make (this creates an empty map)
    colors = make(map[string]string)
    colors["red"] = "#FF0000"
    colors["green"] = "#00FF00"
    colors["blue"] = "#0000FF"
    
    fmt.Printf("Colors: %v\n", colors)
    fmt.Printf("Length: %d\n", len(colors))
}