package main

import "fmt"

func main() {
    // String literals
    var name string = "Alice"
    var greeting string = `Hello, World!
This is a multi-line string.`
    
    // String concatenation
    var fullName = "Alice" + " " + "Smith"
    
    // String length
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Length: %d\n", len(name))
    fmt.Printf("Full name: %s\n", fullName)
    fmt.Printf("Greeting:\n%s\n", greeting)
}