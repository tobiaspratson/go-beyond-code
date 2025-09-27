package main

import "fmt"

// Define a struct - this is like creating a blueprint
type Person struct {
    Name    string  // Field name and type
    Age     int     // Each field has a name and type
    Email   string
    Address string
}

func main() {
    // Create a person instance using struct literal
    person1 := Person{
        Name:    "Alice",
        Age:     25,
        Email:   "alice@example.com",
        Address: "123 Main St",
    }
    
    // Access fields using dot notation
    fmt.Printf("Person: %+v\n", person1)
    fmt.Printf("Name: %s, Age: %d\n", person1.Name, person1.Age)
    
    // You can also access individual fields
    fmt.Printf("Email: %s\n", person1.Email)
    fmt.Printf("Address: %s\n", person1.Address)
}