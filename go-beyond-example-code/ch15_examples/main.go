package main

import (
    "fmt"
    "reflect"
    "strings"
)

type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email"`
    Active  bool   `json:"active"`
    private string // no tag
}

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
    Zip    string `json:"zip"`
}

type Employee struct {
    Person
    Address Address `json:"address"`
    Skills  []string `json:"skills"`
}

func toJSON(v interface{}) (string, error) {
    t := reflect.TypeOf(v)
    val := reflect.ValueOf(v)
    
    // Handle different kinds
    switch val.Kind() {
    case reflect.Struct:
        return structToJSON(t, val)
    case reflect.Slice:
        return sliceToJSON(val)
    case reflect.Map:
        return mapToJSON(val)
    case reflect.Ptr:
        if val.IsNil() {
            return "null", nil
        }
        return toJSON(val.Elem().Interface())
    default:
        return valueToJSON(val)
    }
}

func structToJSON(t reflect.Type, val reflect.Value) (string, error) {
    var fields []string
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := val.Field(i)
        
        // Skip unexported fields
        if field.PkgPath != "" {
            continue
        }
        
        // Get JSON tag
        jsonTag := field.Tag.Get("json")
        if jsonTag == "" {
            jsonTag = strings.ToLower(field.Name)
        }
        
        // Handle omitempty
        if strings.Contains(jsonTag, "omitempty") && isEmpty(value) {
            continue
        }
        
        // Remove omitempty from tag
        jsonTag = strings.Split(jsonTag, ",")[0]
        
        // Format the field
        fieldValue, err := valueToJSON(value)
        if err != nil {
            return "", err
        }
        
        fieldStr := fmt.Sprintf(`"%s":%s`, jsonTag, fieldValue)
        fields = append(fields, fieldStr)
    }
    
    return "{" + strings.Join(fields, ",") + "}", nil
}

func sliceToJSON(val reflect.Value) (string, error) {
    var elements []string
    
    for i := 0; i < val.Len(); i++ {
        element, err := valueToJSON(val.Index(i))
        if err != nil {
            return "", err
        }
        elements = append(elements, element)
    }
    
    return "[" + strings.Join(elements, ",") + "]", nil
}

func mapToJSON(val reflect.Value) (string, error) {
    var pairs []string
    
    for _, key := range val.MapKeys() {
        value := val.MapIndex(key)
        
        keyStr, err := valueToJSON(key)
        if err != nil {
            return "", err
        }
        
        valueStr, err := valueToJSON(value)
        if err != nil {
            return "", err
        }
        
        pair := fmt.Sprintf("%s:%s", keyStr, valueStr)
        pairs = append(pairs, pair)
    }
    
    return "{" + strings.Join(pairs, ",") + "}", nil
}

func valueToJSON(val reflect.Value) (string, error) {
    switch val.Kind() {
    case reflect.String:
        return fmt.Sprintf(`"%s"`, val.String()), nil
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return fmt.Sprintf("%d", val.Int()), nil
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return fmt.Sprintf("%d", val.Uint()), nil
    case reflect.Float32, reflect.Float64:
        return fmt.Sprintf("%g", val.Float()), nil
    case reflect.Bool:
        return fmt.Sprintf("%t", val.Bool()), nil
    case reflect.Slice, reflect.Array:
        return sliceToJSON(val)
    case reflect.Map:
        return mapToJSON(val)
    case reflect.Struct:
        return structToJSON(val.Type(), val)
    case reflect.Ptr:
        if val.IsNil() {
            return "null", nil
        }
        return valueToJSON(val.Elem())
    case reflect.Interface:
        if val.IsNil() {
            return "null", nil
        }
        return valueToJSON(val.Elem())
    default:
        return "null", nil
    }
}

func isEmpty(val reflect.Value) bool {
    switch val.Kind() {
    case reflect.String:
        return val.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return val.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return val.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return val.Float() == 0
    case reflect.Bool:
        return !val.Bool()
    case reflect.Slice, reflect.Array:
        return val.Len() == 0
    case reflect.Map:
        return val.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return val.IsNil()
    default:
        return false
    }
}

func main() {
    // Test with different types
    p := Person{
        Name:    "Alice",
        Age:     25,
        Email:   "alice@example.com",
        Active:  true,
        private: "secret",
    }
    
    json, err := toJSON(p)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Person JSON: %s\n", json)
    }
    
    // Test with nested struct
    emp := Employee{
        Person: Person{
            Name:   "John Doe",
            Age:    30,
            Email:  "john@example.com",
            Active: true,
        },
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        },
        Skills: []string{"Go", "Python", "JavaScript"},
    }
    
    json, err = toJSON(emp)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Employee JSON: %s\n", json)
    }
    
    // Test with slice
    people := []Person{p, emp.Person}
    json, err = toJSON(people)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("People JSON: %s\n", json)
    }
}