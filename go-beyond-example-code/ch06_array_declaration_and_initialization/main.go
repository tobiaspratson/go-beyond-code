package main

import "fmt"

func main() {
    // Method 1: Declare an array of 5 integers (zero-initialized)
    var numbers [5]int
    fmt.Printf("Numbers: %v\n", numbers)  // [0 0 0 0 0]
    fmt.Printf("Type: %T\n", numbers)     // [5]int
    
    // Method 2: Initialize with values using full syntax
    var primes [5]int = [5]int{2, 3, 5, 7, 11}
    fmt.Printf("Primes: %v\n", primes)  // [2 3 5 7 11]
    
    // Method 3: Short declaration with initialization
    colors := [3]string{"red", "green", "blue"}
    fmt.Printf("Colors: %v\n", colors)  // [red green blue]
    
    // Method 4: Let Go infer the size using ellipsis
    temperatures := [...]float64{20.5, 25.3, 18.7, 22.1}
    fmt.Printf("Temperatures: %v\n", temperatures)  // [20.5 25.3 18.7 22.1]
    fmt.Printf("Array length: %d\n", len(temperatures))  // 4
    fmt.Printf("Type: %T\n", temperatures)  // [4]float64
    
    // Method 5: Partial initialization (remaining elements get zero values)
    partial := [5]int{1, 2}  // [1 2 0 0 0]
    fmt.Printf("Partial: %v\n", partial)
    
    // Method 6: Initialize specific indices
    sparse := [5]int{1: 10, 3: 30}  // [0 10 0 30 0]
    fmt.Printf("Sparse: %v\n", sparse)
}