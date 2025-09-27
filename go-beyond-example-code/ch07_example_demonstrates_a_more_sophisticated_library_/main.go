package main

import "fmt"
import "strings"

// Base interface for all library items
type LibraryItem interface {
    GetID() string
    GetTitle() string
    GetInfo() string
    IsAvailable() bool
    Rent() bool
    Return() bool
}

// Base struct for common item properties
type BaseItem struct {
    ID        string
    Title     string
    IsRented  bool
    RentedBy  string
    RentedAt  string
}

func (b *BaseItem) GetID() string {
    return b.ID
}

func (b *BaseItem) GetTitle() string {
    return b.Title
}

func (b BaseItem) IsAvailable() bool {
    return !b.IsRented
}

func (b *BaseItem) Rent() bool {
    if !b.IsRented {
        b.IsRented = true
        return true
    }
    return false
}

func (b *BaseItem) Return() bool {
    if b.IsRented {
        b.IsRented = false
        b.RentedBy = ""
        b.RentedAt = ""
        return true
    }
    return false
}

// Book with specific properties
type Book struct {
    BaseItem
    Author    string
    Pages     int
    Genre     string
    Publisher string
}

func (b Book) GetInfo() string {
    status := "Available"
    if b.IsRented {
        status = fmt.Sprintf("Rented by %s", b.RentedBy)
    }
    return fmt.Sprintf("Book: '%s' by %s (%d pages, %s) - %s", 
        b.Title, b.Author, b.Pages, b.Genre, status)
}

func (b *Book) Rent() bool {
    if b.BaseItem.Rent() {
        fmt.Printf("Book '%s' has been rented\n", b.Title)
        return true
    } else {
        fmt.Printf("Book '%s' is already rented\n", b.Title)
        return false
    }
}

func (b *Book) Return() bool {
    if b.BaseItem.Return() {
        fmt.Printf("Book '%s' has been returned\n", b.Title)
        return true
    } else {
        fmt.Printf("Book '%s' is not currently rented\n", b.Title)
        return false
    }
}

// DVD with specific properties
type DVD struct {
    BaseItem
    Director  string
    Duration  int // in minutes
    Rating    string
    Year      int
}

func (d DVD) GetInfo() string {
    status := "Available"
    if d.IsRented {
        status = fmt.Sprintf("Rented by %s", d.RentedBy)
    }
    return fmt.Sprintf("DVD: '%s' by %s (%d min, %s) - %s", 
        d.Title, d.Director, d.Duration, d.Rating, status)
}

func (d *DVD) Rent() bool {
    if d.BaseItem.Rent() {
        fmt.Printf("DVD '%s' has been rented\n", d.Title)
        return true
    } else {
        fmt.Printf("DVD '%s' is already rented\n", d.Title)
        return false
    }
}

func (d *DVD) Return() bool {
    if d.BaseItem.Return() {
        fmt.Printf("DVD '%s' has been returned\n", d.Title)
        return true
    } else {
        fmt.Printf("DVD '%s' is not currently rented\n", d.Title)
        return false
    }
}

// User management
type User struct {
    ID       string
    Name     string
    Email    string
    RentedItems []LibraryItem
}

func (u *User) RentItem(item LibraryItem) bool {
    if item.Rent() {
        u.RentedItems = append(u.RentedItems, item)
        return true
    }
    return false
}

func (u *User) ReturnItem(item LibraryItem) bool {
    if item.Return() {
        // Remove from user's rented items
        for i, rentedItem := range u.RentedItems {
            if rentedItem.GetID() == item.GetID() {
                u.RentedItems = append(u.RentedItems[:i], u.RentedItems[i+1:]...)
                break
            }
        }
        return true
    }
    return false
}

func (u User) GetRentedItems() []LibraryItem {
    return u.RentedItems
}

// Enhanced library system
type Library struct {
    Name  string
    Items []LibraryItem
    Users []User
}

func (l *Library) AddItem(item LibraryItem) {
    l.Items = append(l.Items, item)
    fmt.Printf("Added item: %s\n", item.GetInfo())
}

func (l *Library) AddUser(user User) {
    l.Users = append(l.Users, user)
    fmt.Printf("Added user: %s (%s)\n", user.Name, user.Email)
}

func (l Library) FindItem(id string) LibraryItem {
    for _, item := range l.Items {
        if item.GetID() == id {
            return item
        }
    }
    return nil
}

func (l Library) FindUser(id string) *User {
    for i := range l.Users {
        if l.Users[i].ID == id {
            return &l.Users[i]
        }
    }
    return nil
}

func (l Library) SearchItems(query string) []LibraryItem {
    var results []LibraryItem
    query = strings.ToLower(query)
    
    for _, item := range l.Items {
        if strings.Contains(strings.ToLower(item.GetTitle()), query) {
            results = append(results, item)
        }
    }
    return results
}

func (l Library) ListAvailableItems() {
    fmt.Println("\n=== Available Items ===")
    for i, item := range l.Items {
        if item.IsAvailable() {
            fmt.Printf("%d. %s\n", i+1, item.GetInfo())
        }
    }
}

func (l Library) ListAllItems() {
    fmt.Printf("\n=== %s Library Items ===\n", l.Name)
    for i, item := range l.Items {
        fmt.Printf("%d. %s\n", i+1, item.GetInfo())
    }
}

func (l Library) GetStatistics() {
    totalItems := len(l.Items)
    availableItems := 0
    rentedItems := 0
    
    for _, item := range l.Items {
        if item.IsAvailable() {
            availableItems++
        } else {
            rentedItems++
        }
    }
    
    fmt.Printf("\n=== Library Statistics ===\n")
    fmt.Printf("Total items: %d\n", totalItems)
    fmt.Printf("Available: %d\n", availableItems)
    fmt.Printf("Rented: %d\n", rentedItems)
    fmt.Printf("Total users: %d\n", len(l.Users))
}

func main() {
    // Create library
    library := Library{Name: "Central Library"}
    
    // Add books
    book1 := &Book{
        BaseItem: BaseItem{
            ID:   "BK001",
            Title: "The Go Programming Language",
        },
        Author:    "Alan Donovan",
        Pages:     380,
        Genre:     "Programming",
        Publisher: "Addison-Wesley",
    }
    
    book2 := &Book{
        BaseItem: BaseItem{
            ID:   "BK002",
            Title: "Effective Go",
        },
        Author:    "Google",
        Pages:     150,
        Genre:     "Programming",
        Publisher: "Google",
    }
    
    // Add DVDs
    dvd1 := &DVD{
        BaseItem: BaseItem{
            ID:   "DV001",
            Title: "The Matrix",
        },
        Director: "The Wachowskis",
        Duration: 136,
        Rating:   "R",
        Year:     1999,
    }
    
    // Add items to library
    library.AddItem(book1)
    library.AddItem(book2)
    library.AddItem(dvd1)
    
    // Add users
    user1 := User{
        ID:    "USR001",
        Name:  "Alice",
        Email: "alice@example.com",
    }
    
    user2 := User{
        ID:    "USR002",
        Name:  "Bob",
        Email: "bob@example.com",
    }
    
    library.AddUser(user1)
    library.AddUser(user2)
    
    // List all items
    library.ListAllItems()
    
    // Search for items
    fmt.Println("\n=== Search Results for 'Go' ===")
    results := library.SearchItems("Go")
    for i, item := range results {
        fmt.Printf("%d. %s\n", i+1, item.GetInfo())
    }
    
    // User operations
    fmt.Println("\n=== User Operations ===")
    alice := library.FindUser("USR001")
    if alice != nil {
        alice.RentItem(book1)
        alice.RentItem(dvd1)
    }
    
    bob := library.FindUser("USR002")
    if bob != nil {
        bob.RentItem(book2)
    }
    
    // Show updated status
    library.ListAllItems()
    
    // Show statistics
    library.GetStatistics()
    
    // Return items
    fmt.Println("\n=== Returning Items ===")
    if alice != nil {
        alice.ReturnItem(book1)
    }
    
    // Show final status
    library.ListAvailableItems()
}