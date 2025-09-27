package main

import "fmt"

type Base struct {
    Name string
}

func (b Base) Method() {
    fmt.Println("Base method called")
}

func (b Base) CommonMethod() {
    fmt.Println("Base common method")
}

type Middle struct {
    Base  // Embed Base
    Value int
}

func (m Middle) Method() {
    fmt.Println("Middle method called")
}

func (m Middle) MiddleSpecific() {
    fmt.Println("Middle specific method")
}

type Top struct {
    Middle  // Embed Middle (which embeds Base)
    Status  string
}

func (t Top) Method() {
    fmt.Println("Top method called")
}

func (t Top) TopSpecific() {
    fmt.Println("Top specific method")
}

func main() {
    top := Top{
        Middle: Middle{
            Base:  Base{Name: "Test"},
            Value: 42,
        },
        Status: "Active",
    }
    
    // Method resolution order: Top -> Middle -> Base
    top.Method()                    // Calls Top.Method()
    top.CommonMethod()              // Calls Base.CommonMethod() (inherited)
    top.MiddleSpecific()            // Calls Middle.MiddleSpecific() (inherited)
    top.TopSpecific()               // Calls Top.TopSpecific()
    
    // Access promoted fields
    fmt.Printf("Name: %s\n", top.Name)    // From Base
    fmt.Printf("Value: %d\n", top.Value)  // From Middle
    fmt.Printf("Status: %s\n", top.Status) // From Top
}