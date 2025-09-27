package main

import (
    "fmt"
    "reflect"
)

func main() {
    // Different types with same kind
    var a int = 42
    var b int32 = 42
    var c int64 = 42
    
    fmt.Printf("Type of a: %v, Kind: %v\n", reflect.TypeOf(a), reflect.TypeOf(a).Kind())
    fmt.Printf("Type of b: %v, Kind: %v\n", reflect.TypeOf(b), reflect.TypeOf(b).Kind())
    fmt.Printf("Type of c: %v, Kind: %v\n", reflect.TypeOf(c), reflect.TypeOf(c).Kind())
    
    // Pointer example
    ptr := &a
    fmt.Printf("Type of ptr: %v, Kind: %v\n", reflect.TypeOf(ptr), reflect.TypeOf(ptr).Kind())
    fmt.Printf("Type of *ptr: %v, Kind: %v\n", reflect.TypeOf(ptr).Elem(), reflect.TypeOf(ptr).Elem().Kind())
}