package main

import "fmt"

// Base types with conflicting methods
type Logger struct {
    Name string
}

func (l Logger) Log(message string) {
    fmt.Printf("[%s] %s\n", l.Name, message)
}

type Debugger struct {
    Name string
}

func (d Debugger) Log(message string) {
    fmt.Printf("DEBUG [%s] %s\n", d.Name, message)
}

// Type with conflicting embedded methods
type Service struct {
    Logger
    Debugger
    Name string
}

// Override the conflicting method
func (s Service) Log(message string) {
    fmt.Printf("SERVICE [%s] %s\n", s.Name, message)
}

// Method that calls specific embedded methods
func (s Service) LogWithLogger(message string) {
    s.Logger.Log(message)  // Explicitly call Logger's Log
}

func (s Service) LogWithDebugger(message string) {
    s.Debugger.Log(message)  // Explicitly call Debugger's Log
}

func main() {
    service := Service{
        Logger:   Logger{Name: "MainLogger"},
        Debugger: Debugger{Name: "MainDebugger"},
        Name:     "MyService",
    }
    
    fmt.Println("=== Method Resolution ===")
    service.Log("This uses Service's Log method")
    service.LogWithLogger("This uses Logger's Log method")
    service.LogWithDebugger("This uses Debugger's Log method")
}