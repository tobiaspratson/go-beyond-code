package main

import (
    "errors"
    "fmt"
)

type Account struct {
    Balance float64
    Owner   string
}

var (
    ErrInsufficientFunds = errors.New("insufficient funds")
    ErrInvalidAmount     = errors.New("invalid amount")
)

func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return ErrInvalidAmount
    }
    if amount > a.Balance {
        return ErrInsufficientFunds
    }
    
    a.Balance -= amount
    return nil
}

func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return ErrInvalidAmount
    }
    
    a.Balance += amount
    return nil
}

func (a Account) GetBalance() float64 {
    return a.Balance
}

func main() {
    account := Account{Balance: 1000.0, Owner: "Alice"}
    
    // Successful withdrawal
    err := account.Withdraw(100.0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Withdrawal successful. Balance: $%.2f\n", account.GetBalance())
    }
    
    // Failed withdrawal - insufficient funds
    err = account.Withdraw(2000.0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Failed withdrawal - invalid amount
    err = account.Withdraw(-50.0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Successful deposit
    err = account.Deposit(500.0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Deposit successful. Balance: $%.2f\n", account.GetBalance())
    }
}