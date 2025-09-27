package main

import "fmt"

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    // Check odd divisors only
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var limit int
    fmt.Print("Find primes up to: ")
    fmt.Scanln(&limit)
    
    if limit < 2 {
        fmt.Println("No prime numbers below 2")
        return
    }
    
    fmt.Printf("Prime numbers up to %d:\n", limit)
    count := 0
    for i := 2; i <= limit; i++ {
        if isPrime(i) {
            fmt.Printf("%3d ", i)
            count++
            if count%10 == 0 { // New line every 10 primes
                fmt.Println()
            }
        }
    }
    if count%10 != 0 {
        fmt.Println()
    }
    fmt.Printf("Found %d prime numbers\n", count)
    
    // Show first N primes
    fmt.Print("\nHow many primes to show? ")
    var n int
    fmt.Scanln(&n)
    
    fmt.Printf("First %d prime numbers:\n", n)
    found := 0
    num := 2
    for found < n {
        if isPrime(num) {
            fmt.Printf("%3d ", num)
            found++
            if found%10 == 0 {
                fmt.Println()
            }
        }
        num++
    }
    if found%10 != 0 {
        fmt.Println()
    }
}