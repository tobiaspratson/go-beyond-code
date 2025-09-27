package main

import (
    "fmt"
    "strings"
)

type Contact struct {
    Name  string
    Phone string
    Email string
}

func main() {
    // Map to store contacts
    contacts := make(map[string]Contact)
    
    // Add contacts
    contacts["alice"] = Contact{
        Name:  "Alice Smith",
        Phone: "555-1234",
        Email: "alice@example.com",
    }
    
    contacts["bob"] = Contact{
        Name:  "Bob Johnson",
        Phone: "555-5678",
        Email: "bob@example.com",
    }
    
    // Search for contact
    searchName := "alice"
    if contact, exists := contacts[searchName]; exists {
        fmt.Printf("Found contact: %+v\n", contact)
    } else {
        fmt.Printf("Contact '%s' not found\n", searchName)
    }
    
    // List all contacts
    fmt.Println("\nAll contacts:")
    for key, contact := range contacts {
        fmt.Printf("%s: %s (%s)\n", key, contact.Name, contact.Phone)
    }
    
    // Search by partial name
    fmt.Println("\nSearching for 'Alice':")
    searchResults := searchContacts(contacts, "Alice")
    for _, contact := range searchResults {
        fmt.Printf("Found: %s (%s)\n", contact.Name, contact.Phone)
    }
    
    // Add new contact
    addContact(contacts, "charlie", Contact{
        Name:  "Charlie Brown",
        Phone: "555-9999",
        Email: "charlie@example.com",
    })
    
    fmt.Println("\nAfter adding Charlie:")
    for key, contact := range contacts {
        fmt.Printf("%s: %s\n", key, contact.Name)
    }
}

// Search contacts by partial name match
func searchContacts(contacts map[string]Contact, searchTerm string) []Contact {
    var results []Contact
    searchTerm = strings.ToLower(searchTerm)
    
    for _, contact := range contacts {
        if strings.Contains(strings.ToLower(contact.Name), searchTerm) {
            results = append(results, contact)
        }
    }
    
    return results
}

// Add contact with validation
func addContact(contacts map[string]Contact, key string, contact Contact) bool {
    if key == "" || contact.Name == "" {
        return false
    }
    
    contacts[key] = contact
    return true
}