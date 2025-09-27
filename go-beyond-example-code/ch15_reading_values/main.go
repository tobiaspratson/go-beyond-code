package main

import (
    "fmt"
    "reflect"
)

func readValue(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    fmt.Printf("=== Value Inspection ===\n")
    fmt.Printf("Value: %v\n", val.Interface())
    fmt.Printf("Type: %v\n", typ)
    fmt.Printf("Kind: %v\n", val.Kind())
    fmt.Printf("Can set: %v\n", val.CanSet())
    fmt.Printf("Can address: %v\n", val.CanAddr())
    fmt.Printf("Is valid: %v\n", val.IsValid())
    fmt.Printf("Is nil: %v\n", val.IsNil())
    
    // Show specific value methods based on kind
    switch val.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Printf("Int value: %d\n", val.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        fmt.Printf("Uint value: %d\n", val.Uint())
    case reflect.Float32, reflect.Float64:
        fmt.Printf("Float value: %f\n", val.Float())
    case reflect.String:
        fmt.Printf("String value: %s\n", val.String())
        fmt.Printf("String length: %d\n", val.Len())
    case reflect.Bool:
        fmt.Printf("Bool value: %t\n", val.Bool())
    case reflect.Slice, reflect.Array:
        fmt.Printf("Length: %d\n", val.Len())
        fmt.Printf("Capacity: %d\n", val.Cap())
        // Show elements
        for i := 0; i < val.Len() && i < 3; i++ {
            fmt.Printf("  [%d]: %v\n", i, val.Index(i).Interface())
        }
        if val.Len() > 3 {
            fmt.Printf("  ... and %d more\n", val.Len()-3)
        }
    case reflect.Map:
        fmt.Printf("Map length: %d\n", val.Len())
        keys := val.MapKeys()
        for i, key := range keys {
            if i >= 3 {
                fmt.Printf("  ... and %d more\n", len(keys)-3)
                break
            }
            value := val.MapIndex(key)
            fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
        }
    case reflect.Ptr:
        fmt.Printf("Pointer value: %p\n", val.Pointer())
        if !val.IsNil() {
            fmt.Printf("Dereferenced: %v\n", val.Elem().Interface())
        }
    case reflect.Struct:
        fmt.Printf("Struct fields: %d\n", val.NumField())
        for i := 0; i < val.NumField(); i++ {
            field := val.Field(i)
            fieldType := typ.Field(i)
            fmt.Printf("  %s: %v (%v)\n", fieldType.Name, field.Interface(), field.Type())
        }
    }
    fmt.Println("======================\n")
}

func main() {
    // Test with different value types
    readValue(42)
    readValue("Hello World")
    readValue(3.14159)
    readValue(true)
    readValue([]int{1, 2, 3, 4, 5})
    readValue(map[string]int{"a": 1, "b": 2, "c": 3})
    
    // Test with pointer
    x := 42
    readValue(&x)
    
    // Test with nil pointer
    var nilPtr *int
    readValue(nilPtr)
    
    // Test with struct
    type Person struct {
        Name string
        Age  int
    }
    readValue(Person{Name: "Alice", Age: 25})
}