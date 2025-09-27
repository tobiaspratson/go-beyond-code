package main

import (
    "fmt"
    "sync"
    "time"
)

type BankAccount struct {
    balance int
    mutex   sync.Mutex
}

func (ba *BankAccount) Deposit(amount int) {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    
    fmt.Printf("Depositing %d, current balance: %d\n", amount, ba.balance)
    ba.balance += amount
    fmt.Printf("New balance: %d\n", ba.balance)
}

func (ba *BankAccount) Withdraw(amount int) bool {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    
    if ba.balance >= amount {
        fmt.Printf("Withdrawing %d, current balance: %d\n", amount, ba.balance)
        ba.balance -= amount
        fmt.Printf("New balance: %d\n", ba.balance)
        return true
    }
    
    fmt.Printf("Insufficient funds for withdrawal of %d\n", amount)
    return false
}

func (ba *BankAccount) Balance() int {
    ba.mutex.Lock()
    defer ba.mutex.Unlock()
    return ba.balance
}

func main() {
    account := &BankAccount{balance: 1000}
    var wg sync.WaitGroup
    
    // Multiple depositors
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 5; j++ {
                account.Deposit(100)
                time.Sleep(10 * time.Millisecond)
            }
        }(i)
    }
    
    // Multiple withdrawers
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 3; j++ {
                account.Withdraw(200)
                time.Sleep(15 * time.Millisecond)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final balance: %d\n", account.Balance())
}