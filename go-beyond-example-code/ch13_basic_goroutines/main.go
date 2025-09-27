package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello %s! (%d)\n", name, i+1)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Run function normally (sequential)
    fmt.Println("=== Sequential ===")
    sayHello("Alice")
    sayHello("Bob")
    
    // Run function as goroutine (concurrent)
    fmt.Println("\n=== Concurrent ===")
    go sayHello("Alice")
    go sayHello("Bob")
    
    // Wait a bit to see the output
    time.Sleep(1 * time.Second)
    fmt.Println("Main function completed")
}