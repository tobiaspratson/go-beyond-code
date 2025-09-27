package main

import (
    "fmt"
    "reflect"
)

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
    Zip    string `json:"zip"`
}

type Employee struct {
    Name    string  `json:"name" validate:"required"`
    Age     int     `json:"age" validate:"min=18"`
    Salary  float64 `json:"salary" validate:"min=0"`
    Address Address `json:"address"`
    Skills  []string `json:"skills"`
}

func advancedStructAccess() {
    emp := Employee{
        Name:   "John Doe",
        Age:    30,
        Salary: 75000.0,
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        },
        Skills: []string{"Go", "Python", "JavaScript"},
    }
    
    val := reflect.ValueOf(emp)
    typ := reflect.TypeOf(emp)
    
    fmt.Printf("=== Advanced Struct Access ===\n")
    fmt.Printf("Struct: %+v\n", emp)
    fmt.Printf("Type: %v\n", typ)
    fmt.Printf("Kind: %v\n", val.Kind())
    fmt.Printf("NumField: %d\n", val.NumField())
    
    fmt.Println("\n=== Field-by-Field Access ===")
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        fmt.Printf("Field %d: %s\n", i, fieldType.Name)
        fmt.Printf("  Type: %v\n", fieldType.Type)
        fmt.Printf("  Value: %v\n", field.Interface())
        fmt.Printf("  Kind: %v\n", field.Kind())
        fmt.Printf("  Can set: %v\n", field.CanSet())
        fmt.Printf("  Tag: %s\n", fieldType.Tag)
        
        // Handle nested struct
        if field.Kind() == reflect.Struct {
            fmt.Printf("  Nested struct fields:\n")
            for j := 0; j < field.NumField(); j++ {
                nestedField := field.Field(j)
                nestedFieldType := field.Type().Field(j)
                fmt.Printf("    %s: %v\n", nestedFieldType.Name, nestedField.Interface())
            }
        }
        
        // Handle slice fields
        if field.Kind() == reflect.Slice {
            fmt.Printf("  Slice length: %d\n", field.Len())
            for j := 0; j < field.Len(); j++ {
                fmt.Printf("    [%d]: %v\n", j, field.Index(j).Interface())
            }
        }
        
        fmt.Println()
    }
}

func fieldByNameAccess() {
    fmt.Println("=== Field by Name Access ===")
    
    emp := Employee{
        Name:   "Jane Smith",
        Age:    28,
        Salary: 80000.0,
        Address: Address{
            Street: "456 Oak Ave",
            City:   "San Francisco",
            Zip:    "94102",
        },
        Skills: []string{"Java", "Go", "Docker"},
    }
    
    val := reflect.ValueOf(emp)
    
    // Access fields by name
    nameField := val.FieldByName("Name")
    if nameField.IsValid() {
        fmt.Printf("Name: %v\n", nameField.String())
    }
    
    ageField := val.FieldByName("Age")
    if ageField.IsValid() {
        fmt.Printf("Age: %v\n", ageField.Int())
    }
    
    salaryField := val.FieldByName("Salary")
    if salaryField.IsValid() {
        fmt.Printf("Salary: %v\n", salaryField.Float())
    }
    
    // Access nested struct field
    addressField := val.FieldByName("Address")
    if addressField.IsValid() && addressField.Kind() == reflect.Struct {
        cityField := addressField.FieldByName("City")
        if cityField.IsValid() {
            fmt.Printf("City: %v\n", cityField.String())
        }
    }
    
    // Access slice field
    skillsField := val.FieldByName("Skills")
    if skillsField.IsValid() && skillsField.Kind() == reflect.Slice {
        fmt.Printf("Skills: ")
        for i := 0; i < skillsField.Len(); i++ {
            if i > 0 {
                fmt.Printf(", ")
            }
            fmt.Printf("%v", skillsField.Index(i).String())
        }
        fmt.Println()
    }
}

func main() {
    advancedStructAccess()
    fieldByNameAccess()
}