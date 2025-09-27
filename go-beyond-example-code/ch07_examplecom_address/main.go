package main

import "fmt"

type Student struct {
    Name     string
    Age      int
    GPA      float64
    IsActive bool
}

func main() {
    // Zero values for each type
    var student Student
    fmt.Printf("Zero values: %+v\n", student)
    fmt.Printf("Name: '%s' (empty string)\n", student.Name)
    fmt.Printf("Age: %d (zero)\n", student.Age)
    fmt.Printf("GPA: %.2f (zero)\n", student.GPA)
    fmt.Printf("IsActive: %t (false)\n", student.IsActive)
}