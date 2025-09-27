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

func (c Calculator) Sum(numbers ...int) int {
    sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func callMethod(obj interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
    val := reflect.ValueOf(obj)
    
    // Handle pointer methods
    if val.Kind() == reflect.Ptr {
        method := val.MethodByName(methodName)
        if method.IsValid() {
            return callReflectMethod(method, args...)
        }
    }
    
    // Handle value methods
    method := val.MethodByName(methodName)
    if !method.IsValid() {
        return nil, fmt.Errorf("method %s not found", methodName)
    }
    
    return callReflectMethod(method, args...)
}

func callReflectMethod(method reflect.Value, args ...interface{}) ([]interface{}, error) {
    // Convert args to reflect.Values
    argValues := make([]reflect.Value, len(args))
    for i, arg := range args {
        argValues[i] = reflect.ValueOf(arg)
    }
    
    // Call the method
    results := method.Call(argValues)
    
    // Convert results back to interface{}
    resultInterfaces := make([]interface{}, len(results))
    for i, result := range results {
        resultInterfaces[i] = result.Interface()
    }
    
    return resultInterfaces, nil
}

func callMethodWithValidation(obj interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
    val := reflect.ValueOf(obj)
    typ := reflect.TypeOf(obj)
    
    // Find the method
    var method reflect.Value
    var methodType reflect.Method
    
    // Check value methods first
    if val.Kind() != reflect.Ptr {
        method = val.MethodByName(methodName)
        if method.IsValid() {
            // Get method info
            for i := 0; i < typ.NumMethod(); i++ {
                if typ.Method(i).Name == methodName {
                    methodType = typ.Method(i)
                    break
                }
            }
        }
    }
    
    // Check pointer methods if value method not found
    if !method.IsValid() && val.Kind() == reflect.Ptr {
        method = val.MethodByName(methodName)
        if method.IsValid() {
            for i := 0; i < typ.NumMethod(); i++ {
                if typ.Method(i).Name == methodName {
                    methodType = typ.Method(i)
                    break
                }
            }
        }
    }
    
    if !method.IsValid() {
        return nil, fmt.Errorf("method %s not found", methodName)
    }
    
    // Validate argument count
    expectedArgs := methodType.Type.NumIn() - 1 // Subtract 1 for receiver
    if len(args) != expectedArgs {
        return nil, fmt.Errorf("expected %d arguments, got %d", expectedArgs, len(args))
    }
    
    // Validate argument types
    for i, arg := range args {
        expectedType := methodType.Type.In(i + 1) // +1 for receiver
        argType := reflect.TypeOf(arg)
        
        if !argType.AssignableTo(expectedType) && !argType.ConvertibleTo(expectedType) {
            return nil, fmt.Errorf("argument %d: expected %v, got %v", i, expectedType, argType)
        }
    }
    
    return callReflectMethod(method, args...)
}

func main() {
    calc := Calculator{Name: "MyCalculator"}
    
    fmt.Println("=== Basic Method Calls ===")
    
    // Call Add method
    results, err := callMethod(calc, "Add", 10, 5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Add(10, 5) = %v\n", results[0])
    }
    
    // Call Multiply method
    results, err = callMethod(calc, "Multiply", 3, 4)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Multiply(3, 4) = %v\n", results[0])
    }
    
    // Call Divide method
    results, err = callMethod(calc, "Divide", 10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Divide(10, 2) = %v, error: %v\n", results[0], results[1])
    }
    
    // Call GetName method
    results, err = callMethod(calc, "GetName")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("GetName() = %v\n", results[0])
    }
    
    fmt.Println("\n=== Pointer Method Calls ===")
    
    // Call SetName method (pointer method)
    results, err = callMethod(&calc, "SetName", "NewCalculator")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("SetName called successfully\n")
        fmt.Printf("Calculator name: %s\n", calc.Name)
    }
    
    fmt.Println("\n=== Variadic Method Calls ===")
    
    // Call Sum method with variadic arguments
    results, err = callMethod(calc, "Sum", 1, 2, 3, 4, 5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Sum(1, 2, 3, 4, 5) = %v\n", results[0])
    }
    
    fmt.Println("\n=== Method Calls with Validation ===")
    
    // Test with wrong argument count
    results, err = callMethodWithValidation(calc, "Add", 10)
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }
    
    // Test with wrong argument type
    results, err = callMethodWithValidation(calc, "Add", "10", 5)
    if err != nil {
        fmt.Printf("Expected error: %v\n", err)
    }
    
    // Test with correct arguments
    results, err = callMethodWithValidation(calc, "Add", 10, 5)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Add(10, 5) = %v\n", results[0])
    }
}