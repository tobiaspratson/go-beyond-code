package main

import "fmt"

// Base struct
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s and I'm %d years old", p.Name, p.Age)
}

// Embedded struct
type Employee struct {
    Person  // Embedded struct
    ID      int
    Salary  float64
}

func (e Employee) Work() string {
    return fmt.Sprintf("%s is working (ID: %d, Salary: $%.2f)", e.Name, e.ID, e.Salary)
}

func main() {
    // Create an employee
    emp := Employee{
        Person: Person{
            Name: "Alice",
            Age:  30,
        },
        ID:     1001,
        Salary: 75000.0,
    }
    
    // Access embedded fields directly
    fmt.Printf("Name: %s\n", emp.Name)
    fmt.Printf("Age: %d\n", emp.Age)
    fmt.Printf("ID: %d\n", emp.ID)
    fmt.Printf("Salary: $%.2f\n", emp.Salary)
    
    // Call embedded methods
    fmt.Println(emp.Greet())
    fmt.Println(emp.Work())
}