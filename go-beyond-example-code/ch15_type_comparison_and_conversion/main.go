package main

import (
    "fmt"
    "reflect"
)

func typeComparison() {
    var a int = 42
    var b int32 = 42
    var c int64 = 42
    var d float64 = 42.0
    
    // Get types
    typeA := reflect.TypeOf(a)
    typeB := reflect.TypeOf(b)
    typeC := reflect.TypeOf(c)
    typeD := reflect.TypeOf(d)
    
    fmt.Printf("Type A: %v, Kind: %v\n", typeA, typeA.Kind())
    fmt.Printf("Type B: %v, Kind: %v\n", typeB, typeB.Kind())
    fmt.Printf("Type C: %v, Kind: %v\n", typeC, typeC.Kind())
    fmt.Printf("Type D: %v, Kind: %v\n", typeD, typeD.Kind())
    
    // Compare types
    fmt.Printf("A == B: %v\n", typeA == typeB)
    fmt.Printf("A == C: %v\n", typeA == typeC)
    fmt.Printf("A == D: %v\n", typeA == typeD)
    
    // Compare kinds
    fmt.Printf("A.Kind() == B.Kind(): %v\n", typeA.Kind() == typeB.Kind())
    fmt.Printf("A.Kind() == C.Kind(): %v\n", typeA.Kind() == typeC.Kind())
    fmt.Printf("A.Kind() == D.Kind(): %v\n", typeA.Kind() == typeD.Kind())
    
    // Check assignability
    fmt.Printf("B assignable to A: %v\n", typeB.AssignableTo(typeA))
    fmt.Printf("A assignable to B: %v\n", typeA.AssignableTo(typeB))
    fmt.Printf("C assignable to A: %v\n", typeC.AssignableTo(typeA))
    
    // Check convertibility
    fmt.Printf("B convertible to A: %v\n", typeB.ConvertibleTo(typeA))
    fmt.Printf("A convertible to B: %v\n", typeA.ConvertibleTo(typeB))
    fmt.Printf("C convertible to A: %v\n", typeC.ConvertibleTo(typeA))
    fmt.Printf("D convertible to A: %v\n", typeD.ConvertibleTo(typeA))
}

func main() {
    typeComparison()
}