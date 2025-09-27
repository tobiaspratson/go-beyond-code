package main

import "fmt"

func main() {
    // Different array types
    var intArray [3]int
    var stringArray [3]string
    var boolArray [3]bool
    var floatArray [3]float64
    
    fmt.Printf("int array: %v\n", intArray)      // [0 0 0]
    fmt.Printf("string array: %v\n", stringArray) // [  ]
    fmt.Printf("bool array: %v\n", boolArray)    // [false false false]
    fmt.Printf("float array: %v\n", floatArray)  // [0 0 0]
    
    // Array of structs
    type Person struct {
        Name string
        Age  int
    }
    
    var people [2]Person
    fmt.Printf("People: %+v\n", people)  // [{Name: Age:0} {Name: Age:0}]
}