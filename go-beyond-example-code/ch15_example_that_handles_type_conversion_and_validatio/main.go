package main

import (
    "fmt"
    "reflect"
)

func safeModifyValue(target interface{}, newValue interface{}) error {
    targetVal := reflect.ValueOf(target)
    
    // Check if target is a pointer
    if targetVal.Kind() != reflect.Ptr {
        return fmt.Errorf("target must be a pointer")
    }
    
    // Get the element the pointer points to
    elem := targetVal.Elem()
    
    // Check if it's settable
    if !elem.CanSet() {
        return fmt.Errorf("target value is not settable")
    }
    
    // Get the new value
    newVal := reflect.ValueOf(newValue)
    
    // Check if types are directly assignable
    if newVal.Type().AssignableTo(elem.Type()) {
        elem.Set(newVal)
        return nil
    }
    
    // Check if types are convertible
    if newVal.Type().ConvertibleTo(elem.Type()) {
        converted := newVal.Convert(elem.Type())
        elem.Set(converted)
        return nil
    }
    
    return fmt.Errorf("cannot assign %v to %v", newVal.Type(), elem.Type())
}

func modifyWithConversion() {
    fmt.Println("=== Value Modification with Type Conversion ===")
    
    // Test direct assignment
    var x int = 42
    err := safeModifyValue(&x, 100)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("x is now: %d\n", x)
    }
    
    // Test type conversion
    var y int32 = 42
    err = safeModifyValue(&y, int64(100))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("y is now: %d\n", y)
    }
    
    // Test string conversion
    var s string = "Hello"
    err = safeModifyValue(&s, "World")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("s is now: %s\n", s)
    }
    
    // Test float conversion
    var f float64 = 3.14
    err = safeModifyValue(&f, float32(2.71))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("f is now: %f\n", f)
    }
    
    // This should fail
    err = safeModifyValue(&x, "not a number")
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }
}

func main() {
    modifyWithConversion()
}