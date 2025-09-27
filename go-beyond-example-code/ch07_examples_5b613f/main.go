package main

import "fmt"

// Base account interface
type Account interface {
    GetBalance() float64
    GetInfo() string
    Deposit(amount float64) bool
    Withdraw(amount float64) bool
}

// Base account struct with common functionality
type BaseAccount struct {
    AccountNumber string
    Balance       float64
    Owner         string
    CreatedAt     string
}

func (a *BaseAccount) Deposit(amount float64) bool {
    if amount > 0 {
        a.Balance += amount
        fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, a.Balance)
        return true
    } else {
        fmt.Println("Deposit amount must be positive")
        return false
    }
}

func (a *BaseAccount) GetBalance() float64 {
    return a.Balance
}

func (a BaseAccount) GetInfo() string {
    return fmt.Sprintf("Account %s (Owner: %s, Balance: $%.2f)", 
        a.AccountNumber, a.Owner, a.Balance)
}

// Checking account with overdraft protection
type CheckingAccount struct {
    BaseAccount
    OverdraftLimit float64
}

func (c *CheckingAccount) Withdraw(amount float64) bool {
    if amount <= 0 {
        fmt.Println("Withdrawal amount must be positive")
        return false
    }
    
    if amount <= c.Balance+c.OverdraftLimit {
        c.Balance -= amount
        fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, c.Balance)
        return true
    } else {
        fmt.Printf("Insufficient funds. Available: $%.2f (including overdraft)\n", 
            c.Balance+c.OverdraftLimit)
        return false
    }
}

func (c CheckingAccount) GetInfo() string {
    return fmt.Sprintf("Checking Account %s (Owner: %s, Balance: $%.2f, Overdraft: $%.2f)", 
        c.AccountNumber, c.Owner, c.Balance, c.OverdraftLimit)
}

// Savings account with interest
type SavingsAccount struct {
    BaseAccount
    InterestRate float64
}

func (s *SavingsAccount) Withdraw(amount float64) bool {
    if amount <= 0 {
        fmt.Println("Withdrawal amount must be positive")
        return false
    }
    
    if amount <= s.Balance {
        s.Balance -= amount
        fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, s.Balance)
        return true
    } else {
        fmt.Println("Insufficient funds")
        return false
    }
}

func (s *SavingsAccount) AddInterest() {
    interest := s.Balance * s.InterestRate
    s.Balance += interest
    fmt.Printf("Interest added: $%.2f. New balance: $%.2f\n", interest, s.Balance)
}

func (s SavingsAccount) GetInfo() string {
    return fmt.Sprintf("Savings Account %s (Owner: %s, Balance: $%.2f, Rate: %.2f%%)", 
        s.AccountNumber, s.Owner, s.Balance, s.InterestRate*100)
}

// Bank system
type Bank struct {
    Name     string
    Accounts []Account
}

func (b *Bank) AddAccount(account Account) {
    b.Accounts = append(b.Accounts, account)
    fmt.Printf("Added account to %s bank\n", b.Name)
}

func (b Bank) GetTotalDeposits() float64 {
    total := 0.0
    for _, account := range b.Accounts {
        total += account.GetBalance()
    }
    return total
}

func (b Bank) ListAccounts() {
    fmt.Printf("\n=== %s Bank Accounts ===\n", b.Name)
    for i, account := range b.Accounts {
        fmt.Printf("%d. %s\n", i+1, account.GetInfo())
    }
}

func (b Bank) ProcessInterest() {
    fmt.Println("\n=== Processing Interest ===")
    for _, account := range b.Accounts {
        if savings, ok := account.(*SavingsAccount); ok {
            savings.AddInterest()
        }
    }
}

func main() {
    // Create bank
    bank := Bank{Name: "First National Bank"}
    
    // Create different account types
    checking := &CheckingAccount{
        BaseAccount: BaseAccount{
            AccountNumber: "CHK001",
            Balance:       1000.00,
            Owner:         "Alice",
            CreatedAt:     "2024-01-01",
        },
        OverdraftLimit: 500.00,
    }
    
    savings := &SavingsAccount{
        BaseAccount: BaseAccount{
            AccountNumber: "SAV001",
            Balance:       5000.00,
            Owner:         "Alice",
            CreatedAt:     "2024-01-01",
        },
        InterestRate: 0.02, // 2% annual interest
    }
    
    // Add accounts to bank
    bank.AddAccount(checking)
    bank.AddAccount(savings)
    
    // List all accounts
    bank.ListAccounts()
    
    // Perform transactions
    fmt.Println("\n=== Transactions ===")
    checking.Deposit(500.00)
    checking.Withdraw(200.00)
    checking.Withdraw(1500.00) // Should use overdraft
    
    savings.Deposit(1000.00)
    savings.Withdraw(500.00)
    
    // Process interest for savings accounts
    bank.ProcessInterest()
    
    // Final account status
    fmt.Println("\n=== Final Status ===")
    bank.ListAccounts()
    fmt.Printf("Total deposits in bank: $%.2f\n", bank.GetTotalDeposits())
}