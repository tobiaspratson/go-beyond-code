package main

import (
    "fmt"
    "reflect"
    "strings"
)

// User struct with various field types
type User struct {
    ID       int    `json:"id" validate:"required"`
    Name     string `json:"name" validate:"required,min=2"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"min=18"`
    Active   bool   `json:"active"`
    Tags     []string `json:"tags"`
}

// Field validator using reflection
type FieldValidator struct {
    rules map[string]string
}

func NewFieldValidator() *FieldValidator {
    return &FieldValidator{
        rules: make(map[string]string),
    }
}

func (fv *FieldValidator) ValidateStruct(obj interface{}) []string {
    var errors []string
    val := reflect.ValueOf(obj)
    typ := reflect.TypeOf(obj)
    
    // Handle pointer types
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
        typ = typ.Elem()
    }
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        // Get validation rules from struct tags
        validateTag := fieldType.Tag.Get("validate")
        if validateTag == "" {
            continue
        }
        
        // Validate field based on rules
        fieldErrors := fv.validateField(field, fieldType.Name, validateTag)
        errors = append(errors, fieldErrors...)
    }
    
    return errors
}

func (fv *FieldValidator) validateField(field reflect.Value, fieldName, rules string) []string {
    var errors []string
    
    ruleList := strings.Split(rules, ",")
    for _, rule := range ruleList {
        rule = strings.TrimSpace(rule)
        
        switch rule {
        case "required":
            if field.IsZero() {
                errors = append(errors, fmt.Sprintf("%s is required", fieldName))
            }
        case "min=2":
            if field.Kind() == reflect.String && len(field.String()) < 2 {
                errors = append(errors, fmt.Sprintf("%s must be at least 2 characters", fieldName))
            }
        case "min=18":
            if field.Kind() == reflect.Int && field.Int() < 18 {
                errors = append(errors, fmt.Sprintf("%s must be at least 18", fieldName))
            }
        case "email":
            if field.Kind() == reflect.String && !strings.Contains(field.String(), "@") {
                errors = append(errors, fmt.Sprintf("%s must be a valid email", fieldName))
            }
        }
    }
    
    return errors
}

// Dynamic field setter
func (fv *FieldValidator) SetField(obj interface{}, fieldName string, value interface{}) error {
    val := reflect.ValueOf(obj)
    
    // Handle pointer types
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    
    field := val.FieldByName(fieldName)
    if !field.IsValid() {
        return fmt.Errorf("field %s not found", fieldName)
    }
    
    if !field.CanSet() {
        return fmt.Errorf("field %s cannot be set", fieldName)
    }
    
    // Convert value to the correct type
    valueVal := reflect.ValueOf(value)
    if !valueVal.Type().AssignableTo(field.Type()) {
        return fmt.Errorf("cannot assign %s to field %s of type %s", 
            valueVal.Type(), fieldName, field.Type())
    }
    
    field.Set(valueVal)
    return nil
}

// Get all field names and values
func (fv *FieldValidator) GetFieldInfo(obj interface{}) map[string]interface{} {
    info := make(map[string]interface{})
    val := reflect.ValueOf(obj)
    typ := reflect.TypeOf(obj)
    
    // Handle pointer types
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
        typ = typ.Elem()
    }
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)
        
        // Get JSON tag for field name
        jsonTag := fieldType.Tag.Get("json")
        if jsonTag == "" {
            jsonTag = fieldType.Name
        }
        
        info[jsonTag] = field.Interface()
    }
    
    return info
}

func main() {
    validator := NewFieldValidator()
    
    // Test with valid user
    user := User{
        ID:     1,
        Name:   "John Doe",
        Email:  "john@example.com",
        Age:    25,
        Active: true,
        Tags:   []string{"admin", "user"},
    }
    
    fmt.Println("=== Field Information ===")
    fieldInfo := validator.GetFieldInfo(user)
    for name, value := range fieldInfo {
        fmt.Printf("%s: %v\n", name, value)
    }
    
    fmt.Println("\n=== Validation ===")
    errors := validator.ValidateStruct(user)
    if len(errors) == 0 {
        fmt.Println("Validation passed!")
    } else {
        for _, err := range errors {
            fmt.Printf("Error: %s\n", err)
        }
    }
    
    fmt.Println("\n=== Dynamic Field Setting ===")
    // Create a new user and set fields dynamically
    newUser := &User{}
    
    fields := map[string]interface{}{
        "ID":     2,
        "Name":   "Jane Smith",
        "Email":  "jane@example.com",
        "Age":    30,
        "Active": true,
        "Tags": []string{"user"},
    }
    
    for fieldName, value := range fields {
        err := validator.SetField(newUser, fieldName, value)
        if err != nil {
            fmt.Printf("Error setting %s: %v\n", fieldName, err)
        } else {
            fmt.Printf("Set %s to %v\n", fieldName, value)
        }
    }
    
    fmt.Printf("\nFinal user: %+v\n", newUser)
}