package main

import (
    "fmt"
    "reflect"
)

// Service interface
type Service interface {
    Process(data string) string
    Validate(data string) bool
    GetInfo() map[string]interface{}
}

// Implementation
type MyService struct {
    Name    string
    Version string
}

func (s MyService) Process(data string) string {
    return fmt.Sprintf("Processed by %s v%s: %s", s.Name, s.Version, data)
}

func (s MyService) Validate(data string) bool {
    return len(data) > 0
}

func (s MyService) GetInfo() map[string]interface{} {
    return map[string]interface{}{
        "name":    s.Name,
        "version": s.Version,
        "type":    "MyService",
    }
}

// Dynamic service caller with enhanced error handling
type ServiceCaller struct {
    service Service
}

func NewServiceCaller(service Service) *ServiceCaller {
    return &ServiceCaller{service: service}
}

func (sc *ServiceCaller) CallMethod(methodName string, args ...interface{}) ([]interface{}, error) {
    val := reflect.ValueOf(sc.service)
    method := val.MethodByName(methodName)
    
    if !method.IsValid() {
        return nil, fmt.Errorf("method %s not found", methodName)
    }
    
    // Check if method is callable
    if !method.CanInterface() {
        return nil, fmt.Errorf("method %s is not callable", methodName)
    }
    
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

// Get available methods
func (sc *ServiceCaller) GetMethods() []string {
    val := reflect.ValueOf(sc.service)
    typ := reflect.TypeOf(sc.service)
    
    var methods []string
    for i := 0; i < val.NumMethod(); i++ {
        method := typ.Method(i)
        methods = append(methods, method.Name)
    }
    
    return methods
}

func main() {
    service := MyService{Name: "MyService", Version: "1.0.0"}
    caller := NewServiceCaller(service)
    
    fmt.Println("=== Available Methods ===")
    methods := caller.GetMethods()
    for _, method := range methods {
        fmt.Printf("- %s\n", method)
    }
    
    fmt.Println("\n=== Dynamic Method Calls ===")
    
    // Call Process method
    results, err := caller.CallMethod("Process", "Hello, World!")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Process result: %v\n", results[0])
    }
    
    // Call Validate method
    results, err = caller.CallMethod("Validate", "test")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Validate result: %v\n", results[0])
    }
    
    // Call GetInfo method
    results, err = caller.CallMethod("GetInfo")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Info result: %v\n", results[0])
    }
}