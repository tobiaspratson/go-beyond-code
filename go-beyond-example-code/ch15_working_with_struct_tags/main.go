package main

import (
    "fmt"
    "reflect"
    "strings"
)

type User struct {
    ID       int    `json:"id" db:"user_id" validate:"required"`
    Username string `json:"username" db:"username" validate:"required,min=3"`
    Email    string `json:"email" db:"email" validate:"required,email"`
    Password string `json:"-" db:"password_hash" validate:"required,min=8"`
    Age      int    `json:"age,omitempty" db:"age" validate:"min=18"`
    Active   bool   `json:"active" db:"is_active"`
}

func parseStructTags() {
    user := User{
        ID:       1,
        Username: "johndoe",
        Email:    "john@example.com",
        Password: "secret123",
        Age:      25,
        Active:   true,
    }
    
    val := reflect.ValueOf(user)
    typ := reflect.TypeOf(user)
    
    fmt.Printf("=== Struct Tag Parsing ===\n")
    fmt.Printf("User: %+v\n\n", user)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        fmt.Printf("Field: %s\n", fieldType.Name)
        fmt.Printf("  Value: %v\n", field.Interface())
        
        // Parse JSON tag
        jsonTag := fieldType.Tag.Get("json")
        if jsonTag != "" {
            fmt.Printf("  JSON tag: %s\n", jsonTag)
            
            // Parse JSON tag options
            jsonParts := strings.Split(jsonTag, ",")
            jsonName := jsonParts[0]
            var options []string
            if len(jsonParts) > 1 {
                options = jsonParts[1:]
            }
            
            fmt.Printf("    JSON name: %s\n", jsonName)
            if len(options) > 0 {
                fmt.Printf("    JSON options: %v\n", options)
            }
        }
        
        // Parse DB tag
        dbTag := fieldType.Tag.Get("db")
        if dbTag != "" {
            fmt.Printf("  DB tag: %s\n", dbTag)
        }
        
        // Parse validate tag
        validateTag := fieldType.Tag.Get("validate")
        if validateTag != "" {
            fmt.Printf("  Validate tag: %s\n", validateTag)
            
            // Parse validation rules
            rules := strings.Split(validateTag, ",")
            for _, rule := range rules {
                rule = strings.TrimSpace(rule)
                fmt.Printf("    Rule: %s\n", rule)
            }
        }
        
        fmt.Println()
    }
}

func main() {
    parseStructTags()
}