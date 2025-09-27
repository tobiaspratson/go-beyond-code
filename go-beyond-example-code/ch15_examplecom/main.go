package main

import (
    "fmt"
    "reflect"
)

type Address struct {
    Street string
    City   string
    Zip    string
}

type Employee struct {
    Name    string
    Age     int
    Salary  float64
    Address Address
    Skills  []string
}

func advancedStructOperations() {
    emp := &Employee{
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
    
    val := reflect.ValueOf(emp).Elem()
    typ := reflect.TypeOf(emp).Elem()
    
    fmt.Printf("=== Advanced Struct Operations ===\n")
    fmt.Printf("Employee: %+v\n\n", emp)
    
    // Iterate through all fields
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        fmt.Printf("Field: %s\n", fieldType.Name)
        fmt.Printf("  Type: %v\n", fieldType.Type)
        fmt.Printf("  Value: %v\n", field.Interface())
        fmt.Printf("  Can set: %v\n", field.CanSet())
        
        // Handle different field types
        switch field.Kind() {
        case reflect.String:
            fmt.Printf("  String value: %s\n", field.String())
        case reflect.Int:
            fmt.Printf("  Int value: %d\n", field.Int())
        case reflect.Float64:
            fmt.Printf("  Float value: %f\n", field.Float())
        case reflect.Struct:
            fmt.Printf("  Struct fields:\n")
            for j := 0; j < field.NumField(); j++ {
                nestedField := field.Field(j)
                nestedFieldType := field.Type().Field(j)
                fmt.Printf("    %s: %v\n", nestedFieldType.Name, nestedField.Interface())
            }
        case reflect.Slice:
            fmt.Printf("  Slice length: %d\n", field.Len())
            for j := 0; j < field.Len(); j++ {
                fmt.Printf("    [%d]: %v\n", j, field.Index(j).Interface())
            }
        }
        
        fmt.Println()
    }
}

func modifyNestedStruct() {
    fmt.Println("=== Modifying Nested Struct ===")
    
    emp := &Employee{
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
    
    fmt.Printf("Before: %+v\n", emp)
    
    val := reflect.ValueOf(emp).Elem()
    
    // Modify nested struct field
    addressField := val.FieldByName("Address")
    if addressField.IsValid() && addressField.CanSet() {
        cityField := addressField.FieldByName("City")
        if cityField.IsValid() && cityField.CanSet() {
            cityField.SetString("Los Angeles")
        }
        
        zipField := addressField.FieldByName("Zip")
        if zipField.IsValid() && zipField.CanSet() {
            zipField.SetString("90210")
        }
    }
    
    // Modify slice field
    skillsField := val.FieldByName("Skills")
    if skillsField.IsValid() && skillsField.CanSet() {
        // Add a new skill
        newSkill := reflect.ValueOf("Kubernetes")
        newSkills := reflect.Append(skillsField, newSkill)
        skillsField.Set(newSkills)
    }
    
    fmt.Printf("After: %+v\n", emp)
}

func main() {
    advancedStructOperations()
    modifyNestedStruct()
}