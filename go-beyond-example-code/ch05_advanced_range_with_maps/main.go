package main

import "fmt"

func main() {
    // Map iteration order is random!
    fmt.Println("=== Map Iteration Order ===")
    data := map[string]int{
        "first":  1,
        "second": 2,
        "third":  3,
        "fourth": 4,
    }
    
    fmt.Println("First iteration:")
    for k, v := range data {
        fmt.Printf("%s: %d\n", k, v)
    }
    
    fmt.Println("\nSecond iteration (order may differ):")
    for k, v := range data {
        fmt.Printf("%s: %d\n", k, v)
    }
    
    // Working with nested maps
    fmt.Println("\n=== Nested Maps ===")
    students := map[string]map[string]int{
        "Alice": {"math": 95, "science": 88, "english": 92},
        "Bob":   {"math": 87, "science": 91, "english": 85},
    }
    
    for student, subjects := range students {
        fmt.Printf("\n%s's grades:\n", student)
        total := 0
        for subject, grade := range subjects {
            fmt.Printf("  %s: %d\n", subject, grade)
            total += grade
        }
        fmt.Printf("  Average: %.2f\n", float64(total)/float64(len(subjects)))
    }
    
    // Modifying map during iteration
    fmt.Println("\n=== Modifying Map During Iteration ===")
    numbers := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    fmt.Println("Before modification:", numbers)
    for key, value := range numbers {
        if value%2 == 0 {
            numbers[key] = value * 2  // Double even values
        }
    }
    fmt.Println("After modification:", numbers)
}