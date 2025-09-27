package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name  string
    Age   int
    Email string
}

func fieldValueAccess() {
    p := Person{
        Name:  "Alice",
        Age:   25,
        Email: "alice@example.com",
    }
    
    v := reflect.ValueOf(p)
    t := reflect.TypeOf(p)
    
    fmt.Printf("=== Field Value Access ===\n")
    fmt.Printf("Person: %+v\n\n", p)
    
    // Access fields by index
    fmt.Println("Fields by index:")
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        fmt.Printf("  %s: %v (%v)\n", field.Name, value.Interface(), value.Type())
    }
    
    // Access specific field by name
    nameField := v.FieldByName("Name")
    if nameField.IsValid() {
        fmt.Printf("\nName field: %v\n", nameField.String())
    }
    
    ageField := v.FieldByName("Age")
    if ageField.IsValid() {
        fmt.Printf("Age field: %v\n", ageField.Int())
    }
    
    emailField := v.FieldByName("Email")
    if emailField.IsValid() {
        fmt.Printf("Email field: %v\n", emailField.String())
    }
}

func fieldValueModification() {
    fmt.Println("\n=== Field Value Modification ===")
    
    p := &Person{
        Name:  "Alice",
        Age:   25,
        Email: "alice@example.com",
    }
    
    fmt.Printf("Before: %+v\n", p)
    
    v := reflect.ValueOf(p).Elem() // Get the struct value (not pointer)
    
    // Modify Name field
    nameField := v.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("Bob")
    }
    
    // Modify Age field
    ageField := v.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(30)
    }
    
    // Modify Email field
    emailField := v.FieldByName("Email")
    if emailField.IsValid() && emailField.CanSet() {
        emailField.SetString("bob@example.com")
    }
    
    fmt.Printf("After: %+v\n", p)
}

func main() {
    fieldValueAccess()
    fieldValueModification()
}