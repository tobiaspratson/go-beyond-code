package main

import "fmt"

func main() {
    // While loop equivalent
    i := 0
    for i < 5 {
        fmt.Printf("While count: %d\n", i)
        i++
    }
    
    // Another while-style example
    sum := 0
    for sum < 100 {
        sum += 10
        fmt.Printf("Sum: %d\n", sum)
    }
}