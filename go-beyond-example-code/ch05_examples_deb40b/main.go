package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(100) + 1
    
    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I'm thinking of a number between 1 and 100.")
    fmt.Printf("(Secret number: %d - for testing)\n", secretNumber)
    
    var guess int
    attempts := 0
    maxAttempts := 10
    
    for {
        fmt.Printf("Attempt %d/%d - Enter your guess: ", attempts+1, maxAttempts)
        fmt.Scanln(&guess)
        attempts++
        
        if guess < 1 || guess > 100 {
            fmt.Println("Please enter a number between 1 and 100.")
            attempts-- // Don't count invalid attempts
            continue
        }
        
        if guess < secretNumber {
            fmt.Println("Too low! Try again.")
        } else if guess > secretNumber {
            fmt.Println("Too high! Try again.")
        } else {
            fmt.Printf("🎉 Congratulations! You guessed it in %d attempts!\n", attempts)
            break
        }
        
        if attempts >= maxAttempts {
            fmt.Printf("Game over! You've used all %d attempts.\n", maxAttempts)
            fmt.Printf("The secret number was %d.\n", secretNumber)
            break
        }
    }
}