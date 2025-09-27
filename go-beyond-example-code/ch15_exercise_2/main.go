package main

import (
    "fmt"
    "reflect"
    "regexp"
    "strconv"
    "strings"
)

type User struct {
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"min=18,max=100"`
    Password string `validate:"required,min=8,max=128"`
    Score    float64 `validate:"min=0,max=100"`
    Active   bool   `validate:"required"`
}

type ValidationError struct {
    Field   string
    Message string
}

func (ve ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", ve.Field, ve.Message)
}

func validateStruct(obj interface{}) []ValidationError {
    var errors []ValidationError
    val := reflect.ValueOf(obj)
    t := reflect.TypeOf(obj)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := t.Field(i)
        
        validateTag := fieldType.Tag.Get("validate")
        if validateTag == "" {
            continue
        }
        
        rules := strings.Split(validateTag, ",")
        for _, rule := range rules {
            rule = strings.TrimSpace(rule)
            
            if err := validateField(field, fieldType.Name, rule); err != nil {
                errors = append(errors, *err)
            }
        }
    }
    
    return errors
}

func validateField(field reflect.Value, fieldName, rule string) *ValidationError {
    switch {
    case rule == "required":
        if isEmpty(field) {
            return &ValidationError{Field: fieldName, Message: "is required"}
        }
    case strings.HasPrefix(rule, "min="):
        if err := validateMin(field, fieldName, rule); err != nil {
            return err
        }
    case strings.HasPrefix(rule, "max="):
        if err := validateMax(field, fieldName, rule); err != nil {
            return err
        }
    case rule == "email":
        if err := validateEmail(field, fieldName); err != nil {
            return err
        }
    }
    
    return nil
}

func isEmpty(field reflect.Value) bool {
    switch field.Kind() {
    case reflect.String:
        return field.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return field.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return field.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return field.Float() == 0
    case reflect.Bool:
        return !field.Bool()
    case reflect.Slice, reflect.Array:
        return field.Len() == 0
    case reflect.Map:
        return field.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return field.IsNil()
    default:
        return false
    }
}

func validateMin(field reflect.Value, fieldName, rule string) *ValidationError {
    minVal, err := strconv.ParseFloat(strings.TrimPrefix(rule, "min="), 64)
    if err != nil {
        return &ValidationError{Field: fieldName, Message: "invalid min rule"}
    }
    
    switch field.Kind() {
    case reflect.String:
        if len(field.String()) < int(minVal) {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %d characters", int(minVal))}
        }
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        if float64(field.Int()) < minVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %g", minVal)}
        }
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        if float64(field.Uint()) < minVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %g", minVal)}
        }
    case reflect.Float32, reflect.Float64:
        if field.Float() < minVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %g", minVal)}
        }
    }
    
    return nil
}

func validateMax(field reflect.Value, fieldName, rule string) *ValidationError {
    maxVal, err := strconv.ParseFloat(strings.TrimPrefix(rule, "max="), 64)
    if err != nil {
        return &ValidationError{Field: fieldName, Message: "invalid max rule"}
    }
    
    switch field.Kind() {
    case reflect.String:
        if len(field.String()) > int(maxVal) {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %d characters", int(maxVal))}
        }
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        if float64(field.Int()) > maxVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %g", maxVal)}
        }
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        if float64(field.Uint()) > maxVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %g", maxVal)}
        }
    case reflect.Float32, reflect.Float64:
        if field.Float() > maxVal {
            return &ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %g", maxVal)}
        }
    }
    
    return nil
}

func validateEmail(field reflect.Value, fieldName string) *ValidationError {
    if field.Kind() != reflect.String {
        return &ValidationError{Field: fieldName, Message: "email validation only applies to strings"}
    }
    
    email := field.String()
    if email == "" {
        return nil // Let required rule handle empty strings
    }
    
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return &ValidationError{Field: fieldName, Message: "must be a valid email address"}
    }
    
    return nil
}

func main() {
    user := User{
        Name:     "A", // Too short
        Email:    "invalid-email", // Invalid email
        Age:      16, // Too young
        Password: "123", // Too short
        Score:    105.5, // Too high
        Active:   false, // Required but false
    }
    
    errors := validateStruct(user)
    if len(errors) == 0 {
        fmt.Println("Validation passed!")
    } else {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("  %s\n", err.Error())
        }
    }
    
    // Test with valid data
    validUser := User{
        Name:     "John Doe",
        Email:    "john@example.com",
        Age:      25,
        Password: "securepassword123",
        Score:    85.5,
        Active:   true,
    }
    
    errors = validateStruct(validUser)
    if len(errors) == 0 {
        fmt.Println("\nValid user passed validation!")
    } else {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("  %s\n", err.Error())
        }
    }
}