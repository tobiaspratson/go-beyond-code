package main

import (
    "errors"
    "fmt"
    "time"
)

// Custom error types for banking
type BankingError struct {
    Code    string
    Message string
    Time    time.Time
}

func (e BankingError) Error() string {
    return fmt.Sprintf("[%s] %s at %v", e.Code, e.Message, e.Time)
}

// Account with enhanced error handling
type BankAccount struct {
    ID       string
    Owner    string
    Balance  float64
    Created  time.Time
    History  []Transaction
}

type Transaction struct {
    Type      string
    Amount    float64
    Timestamp time.Time
    Success   bool
}

// Banking operations with comprehensive error handling
func (a *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return BankingError{
            Code:    "INVALID_AMOUNT",
            Message: "withdrawal amount must be positive",
            Time:    time.Now(),
        }
    }
    
    if amount > a.Balance {
        return BankingError{
            Code:    "INSUFFICIENT_FUNDS",
            Message: fmt.Sprintf("insufficient funds: requested %.2f, available %.2f", amount, a.Balance),
            Time:    time.Now(),
        }
    }
    
    a.Balance -= amount
    a.History = append(a.History, Transaction{
        Type:      "WITHDRAWAL",
        Amount:    amount,
        Timestamp: time.Now(),
        Success:   true,
    })
    
    return nil
}

func (a *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return BankingError{
            Code:    "INVALID_AMOUNT",
            Message: "deposit amount must be positive",
            Time:    time.Now(),
        }
    }
    
    a.Balance += amount
    a.History = append(a.History, Transaction{
        Type:      "DEPOSIT",
        Amount:    amount,
        Timestamp: time.Now(),
        Success:   true,
    })
    
    return nil
}

func (a *BankAccount) Transfer(to *BankAccount, amount float64) error {
    if to == nil {
        return BankingError{
            Code:    "INVALID_ACCOUNT",
            Message: "destination account cannot be nil",
            Time:    time.Now(),
        }
    }
    
    if amount <= 0 {
        return BankingError{
            Code:    "INVALID_AMOUNT",
            Message: "transfer amount must be positive",
            Time:    time.Now(),
        }
    }
    
    // Withdraw from source account
    err := a.Withdraw(amount)
    if err != nil {
        return fmt.Errorf("transfer failed during withdrawal: %w", err)
    }
    
    // Deposit to destination account
    err = to.Deposit(amount)
    if err != nil {
        // Rollback the withdrawal
        a.Balance += amount
        return fmt.Errorf("transfer failed during deposit, rolled back: %w", err)
    }
    
    return nil
}

func (a BankAccount) GetTransactionHistory() []Transaction {
    return a.History
}

func (a BankAccount) GetBalance() float64 {
    return a.Balance
}

// Bank service with error handling
type BankService struct {
    accounts map[string]*BankAccount
}

func NewBankService() *BankService {
    return &BankService{
        accounts: make(map[string]*BankAccount),
    }
}

func (bs *BankService) CreateAccount(id, owner string, initialBalance float64) (*BankAccount, error) {
    if id == "" {
        return nil, BankingError{
            Code:    "INVALID_ID",
            Message: "account ID cannot be empty",
            Time:    time.Now(),
        }
    }
    
    if owner == "" {
        return nil, BankingError{
            Code:    "INVALID_OWNER",
            Message: "account owner cannot be empty",
            Time:    time.Now(),
        }
    }
    
    if initialBalance < 0 {
        return nil, BankingError{
            Code:    "INVALID_BALANCE",
            Message: "initial balance cannot be negative",
            Time:    time.Now(),
        }
    }
    
    if _, exists := bs.accounts[id]; exists {
        return nil, BankingError{
            Code:    "ACCOUNT_EXISTS",
            Message: fmt.Sprintf("account with ID %s already exists", id),
            Time:    time.Now(),
        }
    }
    
    account := &BankAccount{
        ID:      id,
        Owner:   owner,
        Balance: initialBalance,
        Created: time.Now(),
        History: []Transaction{},
    }
    
    bs.accounts[id] = account
    return account, nil
}

func (bs *BankService) GetAccount(id string) (*BankAccount, error) {
    if id == "" {
        return nil, BankingError{
            Code:    "INVALID_ID",
            Message: "account ID cannot be empty",
            Time:    time.Now(),
        }
    }
    
    account, exists := bs.accounts[id]
    if !exists {
        return nil, BankingError{
            Code:    "ACCOUNT_NOT_FOUND",
            Message: fmt.Sprintf("account with ID %s not found", id),
            Time:    time.Now(),
        }
    }
    
    return account, nil
}

func (bs *BankService) Transfer(fromID, toID string, amount float64) error {
    fromAccount, err := bs.GetAccount(fromID)
    if err != nil {
        return fmt.Errorf("failed to get source account: %w", err)
    }
    
    toAccount, err := bs.GetAccount(toID)
    if err != nil {
        return fmt.Errorf("failed to get destination account: %w", err)
    }
    
    return fromAccount.Transfer(toAccount, amount)
}

func main() {
    // Create bank service
    bank := NewBankService()
    
    // Create accounts
    alice, err := bank.CreateAccount("ALICE001", "Alice Smith", 1000.0)
    if err != nil {
        fmt.Printf("Failed to create Alice's account: %v\n", err)
        return
    }
    fmt.Printf("Created account for %s with balance $%.2f\n", alice.Owner, alice.GetBalance())
    
    bob, err := bank.CreateAccount("BOB001", "Bob Johnson", 500.0)
    if err != nil {
        fmt.Printf("Failed to create Bob's account: %v\n", err)
        return
    }
    fmt.Printf("Created account for %s with balance $%.2f\n", bob.Owner, bob.GetBalance())
    
    // Test successful operations
    fmt.Println("\n=== Successful Operations ===")
    
    err = alice.Withdraw(200.0)
    if err != nil {
        fmt.Printf("Withdrawal error: %v\n", err)
    } else {
        fmt.Printf("Alice withdrew $200.00. New balance: $%.2f\n", alice.GetBalance())
    }
    
    err = bob.Deposit(300.0)
    if err != nil {
        fmt.Printf("Deposit error: %v\n", err)
    } else {
        fmt.Printf("Bob deposited $300.00. New balance: $%.2f\n", bob.GetBalance())
    }
    
    // Test transfer
    err = bank.Transfer("ALICE001", "BOB001", 150.0)
    if err != nil {
        fmt.Printf("Transfer error: %v\n", err)
    } else {
        fmt.Printf("Transfer successful. Alice: $%.2f, Bob: $%.2f\n", 
            alice.GetBalance(), bob.GetBalance())
    }
    
    // Test error conditions
    fmt.Println("\n=== Error Conditions ===")
    
    // Invalid withdrawal
    err = alice.Withdraw(-100.0)
    if err != nil {
        fmt.Printf("Invalid withdrawal: %v\n", err)
    }
    
    // Insufficient funds
    err = alice.Withdraw(10000.0)
    if err != nil {
        fmt.Printf("Insufficient funds: %v\n", err)
    }
    
    // Invalid transfer
    err = bank.Transfer("ALICE001", "NONEXISTENT", 100.0)
    if err != nil {
        fmt.Printf("Invalid transfer: %v\n", err)
    }
    
    // Print transaction history
    fmt.Println("\n=== Transaction History ===")
    fmt.Printf("Alice's transactions: %d\n", len(alice.GetTransactionHistory()))
    fmt.Printf("Bob's transactions: %d\n", len(bob.GetTransactionHistory()))
}