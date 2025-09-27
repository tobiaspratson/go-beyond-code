package main

import (
    "fmt"
    "math"
)

func main() {
    // Financial calculations (be careful with money!)
    type Account struct {
        Balance float64  // In dollars
        InterestRate float64  // Annual rate
    }
    
    account := Account{
        Balance: 1000.50,
        InterestRate: 0.05,  // 5%
    }
    
    // Calculate monthly interest
    monthlyRate := account.InterestRate / 12
    monthlyInterest := account.Balance * monthlyRate
    
    fmt.Printf("Account balance: $%.2f\n", account.Balance)
    fmt.Printf("Monthly interest: $%.2f\n", monthlyInterest)
    
    // Scientific calculations
    type Physics struct {
        Mass float64     // kg
        Velocity float64 // m/s
        Temperature float64 // K
    }
    
    physics := Physics{
        Mass: 10.5,
        Velocity: 25.0,
        Temperature: 298.15,  // 25°C in Kelvin
    }
    
    // Kinetic energy: KE = 0.5 * m * v²
    kineticEnergy := 0.5 * physics.Mass * math.Pow(physics.Velocity, 2)
    
    fmt.Printf("Mass: %.1f kg\n", physics.Mass)
    fmt.Printf("Velocity: %.1f m/s\n", physics.Velocity)
    fmt.Printf("Kinetic energy: %.2f J\n", kineticEnergy)
}