package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("=== Basic Random Number Generation ===")
    
    // Seed the random number generator with current time
    rand.Seed(time.Now().UnixNano())
    
    // Generate random integers
    fmt.Printf("Random int: %d\n", rand.Int())
    fmt.Printf("Random int (0-100): %d\n", rand.Intn(101))
    fmt.Printf("Random int (0-9): %d\n", rand.Intn(10))
    
    // Generate random floats
    fmt.Printf("Random float (0-1): %.6f\n", rand.Float64())
    fmt.Printf("Random float (0-10): %.6f\n", rand.Float64()*10)
    fmt.Printf("Random float (5-15): %.6f\n", rand.Float64()*10+5)
    
    // Generate random numbers in specific ranges
    fmt.Println("\n=== Random Numbers in Ranges ===")
    min := 10
    max := 20
    randomInRange := rand.Intn(max-min+1) + min
    fmt.Printf("Random int (%d-%d): %d\n", min, max, randomInRange)
    
    // Float in range
    minFloat := 5.5
    maxFloat := 10.5
    randomFloat := rand.Float64()*(maxFloat-minFloat) + minFloat
    fmt.Printf("Random float (%.1f-%.1f): %.6f\n", minFloat, maxFloat, randomFloat)
    
    // Multiple random numbers
    fmt.Println("\n=== Multiple Random Numbers ===")
    fmt.Print("10 random integers (0-9): ")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", rand.Intn(10))
    }
    fmt.Println()
    
    fmt.Print("10 random floats (0-1): ")
    for i := 0; i < 10; i++ {
        fmt.Printf("%.3f ", rand.Float64())
    }
    fmt.Println()
    
    // Random boolean
    fmt.Println("\n=== Random Boolean ===")
    fmt.Printf("Random boolean: %t\n", rand.Intn(2) == 1)
    fmt.Printf("Random boolean: %t\n", rand.Intn(2) == 1)
}