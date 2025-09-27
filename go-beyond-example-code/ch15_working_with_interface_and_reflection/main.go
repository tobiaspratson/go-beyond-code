package main

import (
    "fmt"
    "reflect"
)

func examineValue(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    fmt.Printf("Value: %v\n", val.Interface())
    fmt.Printf("Type: %v\n", typ)
    fmt.Printf("Kind: %v\n", typ.Kind())
    fmt.Printf("Can set: %v\n", val.CanSet())
    fmt.Printf("Can address: %v\n", val.CanAddr())
    fmt.Println("---")
}

func main() {
    // Test with different types
    examineValue(42)
    examineValue("Hello World")
    examineValue(3.14159)
    examineValue(true)
    examineValue([]int{1, 2, 3})
    examineValue(map[string]int{"a": 1, "b": 2})
    
    // Test with pointer
    x := 42
    examineValue(&x)
}