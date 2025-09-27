package main

import (
    "fmt"
    "runtime"
    "time"
)

// Custom panic type
type CustomPanic struct {
    Message string
    Code    int
    Time    time.Time
}

func (p CustomPanic) Error() string {
    return fmt.Sprintf("custom panic [%d] at %v: %s", p.Code, p.Time, p.Message)
}

// Panic with custom type
func customPanicExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
            
            // Type assertion to get custom panic details
            if cp, ok := r.(CustomPanic); ok {
                fmt.Printf("Custom panic details - Code: %d, Message: %s\n", 
                    cp.Code, cp.Message)
            }
        }
    }()
    
    panic(CustomPanic{
        Message: "Something went wrong",
        Code:    500,
        Time:    time.Now(),
    })
}

// Nested panic recovery
func outerFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Outer function recovered: %v\n", r)
        }
    }()
    
    fmt.Println("Outer function calling inner...")
    innerFunction()
    fmt.Println("Outer function completed")
}

func innerFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Inner function recovered: %v\n", r)
            // Re-panic to propagate to outer function
            panic(fmt.Sprintf("re-panicking: %v", r))
        }
    }()
    
    fmt.Println("Inner function about to panic...")
    panic("Inner function panic")
}

// Panic in goroutine
func panicInGoroutine() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Main goroutine recovered: %v\n", r)
        }
    }()
    
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Goroutine recovered: %v\n", r)
            }
        }()
        
        fmt.Println("Goroutine about to panic...")
        panic("Goroutine panic")
    }()
    
    // Give goroutine time to panic
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Main goroutine continues")
}

// Panic with multiple defer functions
func multipleDeferExample() {
    fmt.Println("Setting up multiple defer functions...")
    
    defer func() {
        fmt.Println("Defer 1: First defer function")
    }()
    
    defer func() {
        fmt.Println("Defer 2: Second defer function")
    }()
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Defer 3: Recovered panic: %v\n", r)
        }
    }()
    
    defer func() {
        fmt.Println("Defer 4: Fourth defer function")
    }()
    
    fmt.Println("About to panic...")
    panic("Multiple defer test")
}

func main() {
    fmt.Println("=== Custom Panic Type ===")
    customPanicExample()
    fmt.Println("Program continues after custom panic")
    
    fmt.Println("\n=== Nested Panic Recovery ===")
    outerFunction()
    fmt.Println("Program continues after nested panic")
    
    fmt.Println("\n=== Panic in Goroutine ===")
    panicInGoroutine()
    fmt.Println("Program continues after goroutine panic")
    
    fmt.Println("\n=== Multiple Defer Functions ===")
    multipleDeferExample()
    fmt.Println("Program continues after multiple defer test")
}