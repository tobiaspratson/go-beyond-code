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

func main() {
    p := &Person{
        Name:  "Alice",
        Age:   25,
        Email: "alice@example.com",
    }
    
    v := reflect.ValueOf(p).Elem() // Get the struct value
    
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
    
    fmt.Printf("Modified person: %+v\n", p)
}