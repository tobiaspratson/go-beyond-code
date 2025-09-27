package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&n)
    
    if n < 0 {
        fmt.Println("Factorial is not defined for negative numbers")
        return
    }
    
    if n > 20 {
        fmt.Println("Warning: Factorial of numbers > 20 may cause integer overflow")
    }
    
    factorial := 1
    fmt.Printf("Calculating %d! step by step:\n", n)
    
    for i := 1; i <= n; i++ {
        factorial *= i
        fmt.Printf("Step %d: %d! = %d\n", i, i, factorial)
    }
    
    fmt.Printf("\nFinal result: %d! = %d\n", n, factorial)
    
    // Show factorial table for small numbers
    fmt.Println("\nFactorial table:")
    for i := 0; i <= 10; i++ {
        fact := 1
        for j := 1; j <= i; j++ {
            fact *= j
        }
        fmt.Printf("%d! = %d\n", i, fact)
    }
}