package main

import "fmt"

func main() {
    // Create a map with mixed data types
    person := map[string]interface{}{
        "name":    "Alice",
        "age":     25,
        "city":    "New York",
        "married": true,
    }
    
    // Access elements (returns interface{} type)
    name := person["name"]
    age := person["age"]
    
    fmt.Printf("Name: %v (type: %T)\n", name, name)
    fmt.Printf("Age: %v (type: %T)\n", age, age)
    
    // Access non-existent key (returns zero value)
    salary := person["salary"]
    fmt.Printf("Salary: %v (zero value, type: %T)\n", salary, salary)
    
    // Type assertion for specific types
    if nameStr, ok := name.(string); ok {
        fmt.Printf("Name as string: %s\n", nameStr)
    }
    
    if ageInt, ok := age.(int); ok {
        fmt.Printf("Age as int: %d\n", ageInt)
    }
}