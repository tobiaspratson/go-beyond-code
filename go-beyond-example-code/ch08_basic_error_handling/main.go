package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Function that can return an error
    num, err := strconv.Atoi("42")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Number: %d\n", num)
    
    // Function that returns an error
    num2, err := strconv.Atoi("not-a-number")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Number: %d\n", num2)
}