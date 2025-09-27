package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string `json:"name" validate:"required"`
    Age     int    `json:"age" validate:"min=0,max=120"`
    Email   string `json:"email" validate:"email"`
    private string // lowercase = unexported
}

func basicStructInspection() {
    p := Person{
        Name:    "Alice",
        Age:     25,
        Email:   "alice@example.com",
        private: "secret",
    }
    
    // Get struct type
    t := reflect.TypeOf(p)
    fmt.Printf("Struct type: %v\n", t)
    fmt.Printf("Number of fields: %d\n", t.NumField())
    fmt.Printf("Package path: %s\n", t.PkgPath())
    fmt.Printf("Size: %d bytes\n", t.Size())
    
    fmt.Println("\n=== Field Information ===")
    // Iterate through fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field %d:\n", i)
        fmt.Printf("  Name: %s\n", field.Name)
        fmt.Printf("  Type: %v\n", field.Type)
        fmt.Printf("  Exported: %v\n", field.PkgPath == "")
        fmt.Printf("  Anonymous: %v\n", field.Anonymous)
        fmt.Printf("  Offset: %d\n", field.Offset)
        fmt.Printf("  Index: %v\n", field.Index)
        fmt.Printf("  Tag: %s\n", field.Tag)
        
        // Parse JSON tag
        jsonTag := field.Tag.Get("json")
        if jsonTag != "" {
            fmt.Printf("  JSON tag: %s\n", jsonTag)
        }
        
        // Parse validate tag
        validateTag := field.Tag.Get("validate")
        if validateTag != "" {
            fmt.Printf("  Validate tag: %s\n", validateTag)
        }
        
        fmt.Println()
    }
}

func main() {
    basicStructInspection()
}