package main

import "fmt"

func main() {
    // This will cause a panic
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    fmt.Println("Before panic")
    panic("Something went wrong!")
    fmt.Println("This won't be printed")
}