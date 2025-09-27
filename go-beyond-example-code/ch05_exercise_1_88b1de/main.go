package main

import "fmt"

func main() {
    var n int
    fmt.Print("How many Fibonacci numbers? ")
    fmt.Scanln(&n)
    
    if n <= 0 {
        fmt.Println("Please enter a positive number")
        return
    }
    
    a, b := 0, 1
    fmt.Printf("First %d Fibonacci numbers: ", n)
    
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b
    }
    fmt.Println()
}