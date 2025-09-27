package main

import "fmt"

type Employee struct {
    ID       int
    Name     string
    Position string
    Salary   float64
}

func main() {
    employees := make(map[int]Employee)
    
    // Add employees
    employees[1] = Employee{1, "Alice", "Developer", 75000}
    employees[2] = Employee{2, "Bob", "Manager", 85000}
    employees[3] = Employee{3, "Charlie", "Designer", 65000}
    
    // Search by ID
    searchID := 2
    if emp, exists := employees[searchID]; exists {
        fmt.Printf("Found employee: %+v\n", emp)
    } else {
        fmt.Printf("Employee with ID %d not found\n", searchID)
    }
    
    // List all employees
    fmt.Println("\nAll employees:")
    for _, emp := range employees {
        fmt.Printf("ID: %d, Name: %s, Position: %s, Salary: $%.2f\n",
            emp.ID, emp.Name, emp.Position, emp.Salary)
    }
}