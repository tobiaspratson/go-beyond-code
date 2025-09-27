package main

import (
    "fmt"
    "reflect"
)

func processValue(v interface{}) {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    fmt.Printf("Processing: %v (Type: %v, Kind: %v)\n", val.Interface(), typ, val.Kind())
    
    switch val.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        fmt.Printf("  Integer: %d\n", val.Int())
        fmt.Printf("  Size: %d bits\n", typ.Size()*8)
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        fmt.Printf("  Unsigned integer: %d\n", val.Uint())
        fmt.Printf("  Size: %d bits\n", typ.Size()*8)
    case reflect.Float32, reflect.Float64:
        fmt.Printf("  Float: %f\n", val.Float())
        fmt.Printf("  Size: %d bits\n", typ.Size()*8)
    case reflect.String:
        fmt.Printf("  String: %s\n", val.String())
        fmt.Printf("  Length: %d\n", val.Len())
    case reflect.Bool:
        fmt.Printf("  Boolean: %t\n", val.Bool())
    case reflect.Slice:
        fmt.Printf("  Slice with %d elements\n", val.Len())
        fmt.Printf("  Element type: %v\n", typ.Elem())
        // Show first few elements
        for i := 0; i < val.Len() && i < 3; i++ {
            fmt.Printf("    [%d]: %v\n", i, val.Index(i).Interface())
        }
        if val.Len() > 3 {
            fmt.Printf("    ... and %d more\n", val.Len()-3)
        }
    case reflect.Map:
        fmt.Printf("  Map with %d keys\n", val.Len())
        fmt.Printf("  Key type: %v\n", typ.Key())
        fmt.Printf("  Value type: %v\n", typ.Elem())
        // Show first few key-value pairs
        keys := val.MapKeys()
        for i, key := range keys {
            if i >= 3 {
                fmt.Printf("    ... and %d more\n", len(keys)-3)
                break
            }
            value := val.MapIndex(key)
            fmt.Printf("    %v: %v\n", key.Interface(), value.Interface())
        }
    case reflect.Ptr:
        fmt.Printf("  Pointer to %v\n", val.Elem().Type())
        if val.IsNil() {
            fmt.Printf("  Value: <nil>\n")
        } else {
            fmt.Printf("  Dereferenced value: %v\n", val.Elem().Interface())
        }
    case reflect.Struct:
        fmt.Printf("  Struct with %d fields\n", val.NumField())
        for i := 0; i < val.NumField(); i++ {
            field := val.Field(i)
            fieldType := typ.Field(i)
            fmt.Printf("    %s: %v (%v)\n", fieldType.Name, field.Interface(), field.Type())
        }
    case reflect.Array:
        fmt.Printf("  Array with %d elements\n", val.Len())
        fmt.Printf("  Element type: %v\n", typ.Elem())
        // Show first few elements
        for i := 0; i < val.Len() && i < 3; i++ {
            fmt.Printf("    [%d]: %v\n", i, val.Index(i).Interface())
        }
        if val.Len() > 3 {
            fmt.Printf("    ... and %d more\n", val.Len()-3)
        }
    case reflect.Interface:
        fmt.Printf("  Interface\n")
        if val.IsNil() {
            fmt.Printf("  Value: <nil>\n")
        } else {
            fmt.Printf("  Dynamic type: %v\n", val.Elem().Type())
            fmt.Printf("  Dynamic value: %v\n", val.Elem().Interface())
        }
    case reflect.Chan:
        fmt.Printf("  Channel\n")
        fmt.Printf("  Element type: %v\n", typ.Elem())
        fmt.Printf("  Direction: %v\n", typ.ChanDir())
    case reflect.Func:
        fmt.Printf("  Function\n")
        fmt.Printf("  NumIn: %d, NumOut: %d\n", typ.NumIn(), typ.NumOut())
    default:
        fmt.Printf("  Unknown type: %v\n", val.Type())
    }
    fmt.Println("---")
}

func main() {
    // Test with various types
    processValue(42)
    processValue(int32(42))
    processValue(int64(42))
    processValue("Hello")
    processValue(3.14)
    processValue(float32(3.14))
    processValue(true)
    processValue([]int{1, 2, 3, 4, 5})
    processValue(map[string]int{"a": 1, "b": 2, "c": 3})
    processValue([3]int{1, 2, 3})
    
    x := 42
    processValue(&x)
    
    var nilPtr *int
    processValue(nilPtr)
    
    // Test with struct
    type Person struct {
        Name string
        Age  int
    }
    processValue(Person{Name: "Alice", Age: 25})
    
    // Test with interface
    var iface interface{} = "Hello"
    processValue(iface)
    
    // Test with channel
    ch := make(chan int, 2)
    processValue(ch)
    
    // Test with function
    processValue(func(x int) int { return x * 2 })
}