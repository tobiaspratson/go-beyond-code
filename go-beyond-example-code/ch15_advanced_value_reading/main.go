package main

import (
    "fmt"
    "reflect"
)

func advancedValueReader(v interface{}) {
    val := reflect.ValueOf(v)
    
    if !val.IsValid() {
        fmt.Printf("Invalid value\n")
        return
    }
    
    fmt.Printf("Value: %v\n", val.Interface())
    fmt.Printf("Type: %v\n", val.Type())
    fmt.Printf("Kind: %v\n", val.Kind())
    
    // Handle nil values
    if val.IsNil() {
        fmt.Printf("Value is nil\n")
        return
    }
    
    // Handle zero values
    if val.IsZero() {
        fmt.Printf("Value is zero value\n")
    }
    
    // Show value using different methods
    switch val.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Printf("Int: %d\n", val.Int())
        fmt.Printf("Can convert to float: %v\n", val.CanConvert(reflect.TypeOf(float64(0))))
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        fmt.Printf("Uint: %d\n", val.Uint())
    case reflect.Float32, reflect.Float64:
        fmt.Printf("Float: %f\n", val.Float())
    case reflect.String:
        fmt.Printf("String: %s\n", val.String())
        fmt.Printf("Length: %d\n", val.Len())
        // Show runes
        runes := []rune(val.String())
        fmt.Printf("Runes: %v\n", runes)
    case reflect.Bool:
        fmt.Printf("Bool: %t\n", val.Bool())
    case reflect.Slice:
        fmt.Printf("Slice length: %d, capacity: %d\n", val.Len(), val.Cap())
        if val.Len() > 0 {
            fmt.Printf("First element: %v\n", val.Index(0).Interface())
            fmt.Printf("Last element: %v\n", val.Index(val.Len()-1).Interface())
        }
    case reflect.Map:
        fmt.Printf("Map length: %d\n", val.Len())
        if val.Len() > 0 {
            keys := val.MapKeys()
            firstKey := keys[0]
            firstValue := val.MapIndex(firstKey)
            fmt.Printf("First entry: %v -> %v\n", firstKey.Interface(), firstValue.Interface())
        }
    case reflect.Ptr:
        fmt.Printf("Pointer: %p\n", val.Pointer())
        if !val.IsNil() {
            elem := val.Elem()
            fmt.Printf("Points to: %v (%v)\n", elem.Interface(), elem.Type())
        }
    case reflect.Struct:
        fmt.Printf("Struct with %d fields\n", val.NumField())
        for i := 0; i < val.NumField(); i++ {
            field := val.Field(i)
            fieldType := val.Type().Field(i)
            fmt.Printf("  %s: %v\n", fieldType.Name, field.Interface())
        }
    }
    fmt.Println("---")
}

func main() {
    // Test with various values
    advancedValueReader(42)
    advancedValueReader("Hello")
    advancedValueReader(0) // zero value
    advancedValueReader("") // zero value
    advancedValueReader([]int{}) // empty slice
    advancedValueReader(map[string]int{}) // empty map
    
    // Test with nil
    var nilSlice []int
    advancedValueReader(nilSlice)
    
    var nilMap map[string]int
    advancedValueReader(nilMap)
    
    var nilPtr *int
    advancedValueReader(nilPtr)
}