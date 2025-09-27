package main

import (
    "fmt"
    "reflect"
)

type Calculator struct {
    Name string
}

func (c Calculator) Add(a, b int) int {
    return a + b
}

func (c Calculator) Multiply(a, b int) int {
    return a * b
}

func (c Calculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func (c Calculator) GetName() string {
    return c.Name
}

func (c *Calculator) SetName(name string) {
    c.Name = name
}

func methodInvocation() {
    fmt.Println("=== Method Invocation ===")
    
    calc := Calculator{Name: "MyCalculator"}
    val := reflect.ValueOf(calc)
    
    // Get all methods
    fmt.Printf("Methods on %T:\n", calc)
    for i := 0; i < val.NumMethod(); i++ {
        method := val.Type().Method(i)
        fmt.Printf("  %s: %v\n", method.Name, method.Type)
    }
    
    // Call a method by name
    addMethod := val.MethodByName("Add")
    if addMethod.IsValid() {
        args := []reflect.Value{
            reflect.ValueOf(10),
            reflect.ValueOf(5),
        }
        results := addMethod.Call(args)
        fmt.Printf("Add(10, 5) = %v\n", results[0].Int())
    }
    
    // Call a method that returns multiple values
    divideMethod := val.MethodByName("Divide")
    if divideMethod.IsValid() {
        args := []reflect.Value{
            reflect.ValueOf(10),
            reflect.ValueOf(2),
        }
        results := divideMethod.Call(args)
        fmt.Printf("Divide(10, 2) = %v, error: %v\n", results[0].Int(), results[1].Interface())
    }
    
    // Call a method that returns an error
    args := []reflect.Value{
        reflect.ValueOf(10),
        reflect.ValueOf(0),
    }
    results := divideMethod.Call(args)
    if results[1].Interface() != nil {
        fmt.Printf("Divide(10, 0) error: %v\n", results[1].Interface())
    }
}

func pointerMethodInvocation() {
    fmt.Println("\n=== Pointer Method Invocation ===")
    
    calc := &Calculator{Name: "MyCalculator"}
    val := reflect.ValueOf(calc)
    
    // Call a pointer method
    setNameMethod := val.MethodByName("SetName")
    if setNameMethod.IsValid() {
        args := []reflect.Value{
            reflect.ValueOf("NewName"),
        }
        setNameMethod.Call(args)
        fmt.Printf("After SetName: %+v\n", calc)
    }
    
    // Call a getter method
    getNameMethod := val.MethodByName("GetName")
    if getNameMethod.IsValid() {
        results := getNameMethod.Call([]reflect.Value{})
        fmt.Printf("GetName() = %v\n", results[0].String())
    }
}

func main() {
    methodInvocation()
    pointerMethodInvocation()
}