package main

import (
    "fmt"
    "reflect"
)

func inspectType(v interface{}) {
    t := reflect.TypeOf(v)
    v_val := reflect.ValueOf(v)
    
    fmt.Printf("Type: %v\n", t)
    fmt.Printf("Kind: %v\n", t.Kind())
    fmt.Printf("Value: %v\n", v_val)
    fmt.Printf("Can set: %v\n", v_val.CanSet())
    fmt.Printf("Can address: %v\n", v_val.CanAddr())
    fmt.Printf("Is nil: %v\n", v_val.IsNil())
    fmt.Printf("Is valid: %v\n", v_val.IsValid())
    fmt.Println("---")
}

func main() {
    // Different types
    inspectType(42)
    inspectType("Hello")
    inspectType(3.14)
    inspectType(true)
    
    // Pointer
    x := 42
    inspectType(&x)
    
    // Slice
    inspectType([]int{1, 2, 3})
    
    // Map
    inspectType(map[string]int{"a": 1, "b": 2})
    
    // Nil pointer
    var nilPtr *int
    inspectType(nilPtr)
}