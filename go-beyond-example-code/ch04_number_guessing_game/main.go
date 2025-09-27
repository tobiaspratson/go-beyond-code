package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())
    
    // Generate random number between 1 and 100
    secretNumber := rand.Intn(100) + 1
    
    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I'm thinking of a number between 1 and 100.")
    
    var guess int
    attempts := 0
    
    for {
        fmt.Print("Enter your guess: ")
        fmt.Scanln(&guess)
        attempts++
        
        if guess < secretNumber {
            fmt.Println("Too low! Try again.")
        } else if guess > secretNumber {
            fmt.Println("Too high! Try again.")
        } else {
            fmt.Printf("Congratulations! You guessed it in %d attempts!\n", attempts)
            break
        }
    }
}