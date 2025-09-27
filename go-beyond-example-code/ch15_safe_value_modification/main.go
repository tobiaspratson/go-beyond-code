package main

import (
    "fmt"
    "reflect"
)

func safeSetValue(v interface{}, newValue interface{}) error {
    val := reflect.ValueOf(v)
    
    // Check if it's a pointer
    if val.Kind() != reflect.Ptr {
        return fmt.Errorf("value must be a pointer")
    }
    
    // Get the element the pointer points to
    elem := val.Elem()
    
    // Check if it's settable
    if !elem.CanSet() {
        return fmt.Errorf("value is not settable")
    }
    
    // Get the new value
    newVal := reflect.ValueOf(newValue)
    
    // Check if types are compatible
    if !newVal.Type().AssignableTo(elem.Type()) {
        return fmt.Errorf("cannot assign %v to %v", newVal.Type(), elem.Type())
    }
    
    // Set the value
    elem.Set(newVal)
    return nil
}

func main() {
    // Test with different types
    var x int = 42
    err := safeSetValue(&x, 100)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("x is now: %d\n", x)
    }
    
    var s string = "Hello"
    err = safeSetValue(&s, "World")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("s is now: %s\n", s)
    }
    
    // This should fail
    err = safeSetValue(&x, "not a number")
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }
}