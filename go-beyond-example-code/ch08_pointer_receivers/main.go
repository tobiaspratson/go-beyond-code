package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Value receiver - cannot modify original
func (p Person) SetAgeValue(age int) {
    p.Age = age
    fmt.Printf("Inside SetAgeValue: Age = %d\n", p.Age)
}

// Pointer receiver - can modify original
func (p *Person) SetAgePointer(age int) {
    p.Age = age
    fmt.Printf("Inside SetAgePointer: Age = %d\n", p.Age)
}

// Pointer receiver - can modify original
func (p *Person) SetName(name string) {
    p.Name = name
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    
    fmt.Printf("Original: %+v\n", person)
    
    // Value receiver
    person.SetAgeValue(30)
    fmt.Printf("After SetAgeValue: %+v\n", person)
    
    // Pointer receiver
    person.SetAgePointer(30)
    fmt.Printf("After SetAgePointer: %+v\n", person)
    
    // Pointer receiver
    person.SetName("Bob")
    fmt.Printf("After SetName: %+v\n", person)
}