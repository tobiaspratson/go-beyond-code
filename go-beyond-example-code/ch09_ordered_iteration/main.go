package main

import (
    "fmt"
    "sort"
)

func main() {
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
        "David":   78,
    }
    
    // Get sorted keys
    var names []string
    for name := range scores {
        names = append(names, name)
    }
    sort.Strings(names)
    
    // Iterate in sorted order
    fmt.Println("Scores in alphabetical order:")
    for _, name := range names {
        fmt.Printf("%s: %d\n", name, scores[name])
    }
}