package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 42
    var y string = "Hello"
    var z bool = true
    
    // Get type information using reflect.TypeOf()
    fmt.Printf("Type of x: %v\n", reflect.TypeOf(x))
    fmt.Printf("Type of y: %v\n", reflect.TypeOf(y))
    fmt.Printf("Type of z: %v\n", reflect.TypeOf(z))
    
    // Get value information using reflect.ValueOf()
    fmt.Printf("Value of x: %v\n", reflect.ValueOf(x))
    fmt.Printf("Value of y: %v\n", reflect.ValueOf(y))
    fmt.Printf("Value of z: %v\n", reflect.ValueOf(z))
    
    // Let's also examine the kind of each type
    fmt.Printf("Kind of x: %v\n", reflect.TypeOf(x).Kind())
    fmt.Printf("Kind of y: %v\n", reflect.TypeOf(y).Kind())
    fmt.Printf("Kind of z: %v\n", reflect.TypeOf(z).Kind())
}