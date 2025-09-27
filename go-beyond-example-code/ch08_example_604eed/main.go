package main

import (
    "fmt"
    "runtime"
)

// Pattern 1: Basic recovery
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Caught panic: %v\n", r)
        }
    }()
    
    fmt.Println("About to panic...")
    panic("Oh no!")
    fmt.Println("This won't be printed")
}

// Pattern 2: Recovery with stack trace
func safeFunctionWithStack() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Caught panic: %v\n", r)
            
            // Print stack trace
            buf := make([]byte, 1024)
            n := runtime.Stack(buf, false)
            fmt.Printf("Stack trace:\n%s\n", buf[:n])
        }
    }()
    
    fmt.Println("About to panic with stack trace...")
    panic("Detailed error")
}

// Pattern 3: Recovery with cleanup
func safeFunctionWithCleanup() {
    resource := "database connection"
    fmt.Printf("Acquired resource: %s\n", resource)
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic occurred, cleaning up resource: %s\n", resource)
            fmt.Printf("Panic: %v\n", r)
        }
    }()
    
    fmt.Println("About to panic...")
    panic("Resource cleanup needed")
}

// Pattern 4: Recovery with error return
func safeFunctionWithError() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("function panicked: %v", r)
        }
    }()
    
    fmt.Println("About to panic...")
    panic("Converting panic to error")
}

func main() {
    fmt.Println("=== Basic Recovery ===")
    safeFunction()
    fmt.Println("Program continues after basic recovery")
    
    fmt.Println("\n=== Recovery with Stack Trace ===")
    safeFunctionWithStack()
    fmt.Println("Program continues after stack trace recovery")
    
    fmt.Println("\n=== Recovery with Cleanup ===")
    safeFunctionWithCleanup()
    fmt.Println("Program continues after cleanup recovery")
    
    fmt.Println("\n=== Recovery with Error Return ===")
    result, err := safeFunctionWithError()
    if err != nil {
        fmt.Printf("Function returned error: %v\n", err)
    } else {
        fmt.Printf("Function result: %s\n", result)
    }
}