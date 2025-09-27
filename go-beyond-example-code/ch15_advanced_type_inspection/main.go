package main

import (
    "fmt"
    "reflect"
)

func detailedInspect(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    fmt.Printf("=== Detailed Type Inspection ===\n")
    fmt.Printf("Value: %v\n", val.Interface())
    fmt.Printf("Type: %v\n", typ)
    fmt.Printf("Kind: %v\n", typ.Kind())
    fmt.Printf("Size: %d bytes\n", typ.Size())
    fmt.Printf("Align: %d\n", typ.Align())
    fmt.Printf("FieldAlign: %d\n", typ.FieldAlign())
    fmt.Printf("NumMethod: %d\n", typ.NumMethod())
    fmt.Printf("PkgPath: %s\n", typ.PkgPath())
    fmt.Printf("String: %s\n", typ.String())
    
    // Show methods if any
    if typ.NumMethod() > 0 {
        fmt.Printf("Methods:\n")
        for i := 0; i < typ.NumMethod(); i++ {
            method := typ.Method(i)
            fmt.Printf("  %s: %v\n", method.Name, method.Type)
        }
    }
    
    // Special handling for different kinds
    switch typ.Kind() {
    case reflect.Slice, reflect.Array:
        fmt.Printf("Element type: %v\n", typ.Elem())
        fmt.Printf("Length: %d\n", val.Len())
    case reflect.Map:
        fmt.Printf("Key type: %v\n", typ.Key())
        fmt.Printf("Element type: %v\n", typ.Elem())
        fmt.Printf("Length: %d\n", val.Len())
    case reflect.Ptr:
        fmt.Printf("Element type: %v\n", typ.Elem())
        if !val.IsNil() {
            fmt.Printf("Dereferenced value: %v\n", val.Elem().Interface())
        }
    case reflect.Struct:
        fmt.Printf("NumField: %d\n", typ.NumField())
        for i := 0; i < typ.NumField(); i++ {
            field := typ.Field(i)
            fmt.Printf("  Field %d: %s (%v)\n", i, field.Name, field.Type)
        }
    }
    
    fmt.Println("================================\n")
}

func main() {
    // Test with different types
    detailedInspect(42)
    detailedInspect("Hello")
    detailedInspect([]int{1, 2, 3})
    detailedInspect(map[string]int{"a": 1})
    
    // Test with struct
    type Person struct {
        Name string
        Age  int
    }
    detailedInspect(Person{Name: "Alice", Age: 25})
}