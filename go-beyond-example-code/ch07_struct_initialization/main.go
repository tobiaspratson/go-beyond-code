package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Pages  int
    Price  float64
}

func main() {
    // Method 1: Named fields (RECOMMENDED)
    // This is the clearest and most maintainable way
    book1 := Book{
        Title:  "The Go Programming Language",
        Author: "Alan Donovan",
        Pages:  380,
        Price:  45.99,
    }
    
    // Method 2: Positional initialization (order matters!)
    // Use this only when the order is obvious and unlikely to change
    book2 := Book{
        "Effective Go",    // Title
        "Google",          // Author
        150,               // Pages
        29.99,             // Price
    }
    
    // Method 3: Zero value initialization
    // All fields get their zero values
    var book3 Book
    fmt.Printf("Zero values: %+v\n", book3)
    
    // Method 4: Partial initialization
    // Only specify the fields you want, others get zero values
    book4 := Book{
        Title: "Learning Go",
        Price: 39.99,
        // Author and Pages will be zero values
    }
    
    // Method 5: Using new() function
    // Returns a pointer to a zero-valued struct
    book5 := new(Book)
    book5.Title = "Go in Action"
    book5.Author = "William Kennedy"
    
    // Method 6: Address operator with literal
    // Creates a pointer to the struct
    book6 := &Book{
        Title:  "Concurrency in Go",
        Author: "Katherine Cox-Buday",
        Pages:  200,
        Price:  44.99,
    }
    
    fmt.Printf("Book 1: %+v\n", book1)
    fmt.Printf("Book 2: %+v\n", book2)
    fmt.Printf("Book 4: %+v\n", book4)
    fmt.Printf("Book 5: %+v\n", *book5)  // Dereference pointer
    fmt.Printf("Book 6: %+v\n", *book6) // Dereference pointer
}