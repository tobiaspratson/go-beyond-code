package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    Country string `json:"country"`
    ZipCode string `json:"zipcode"`
}

type Employee struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Email     string  `json:"email"`
    Age       int     `json:"age"`
    Salary    float64 `json:"salary"`
    Address   Address `json:"address"`
    Skills    []string `json:"skills"`
    IsActive  bool    `json:"is_active"`
}

type Company struct {
    Name      string     `json:"company_name"`
    Founded   int        `json:"founded"`
    Employees []Employee `json:"employees"`
}

func processJSONFile(filename string) (*Company, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    
    var company Company
    err = json.Unmarshal(content, &company)
    if err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }
    
    return &company, nil
}

func main() {
    company, err := processJSONFile("company.json")
    if err != nil {
        fmt.Printf("Error processing JSON: %v\n", err)
        return
    }
    
    fmt.Printf("Company: %s (Founded: %d)\n", company.Name, company.Founded)
    fmt.Printf("Total employees: %d\n", len(company.Employees))
    
    totalSalary := 0.0
    activeCount := 0
    
    for _, emp := range company.Employees {
        if emp.IsActive {
            activeCount++
            totalSalary += emp.Salary
        }
        
        fmt.Printf("Employee: %s (%s) - $%.2f\n", 
            emp.Name, emp.Email, emp.Salary)
        fmt.Printf("  Address: %s, %s, %s %s\n", 
            emp.Address.Street, emp.Address.City, 
            emp.Address.Country, emp.Address.ZipCode)
        fmt.Printf("  Skills: %v\n", emp.Skills)
        fmt.Println()
    }
    
    if activeCount > 0 {
        fmt.Printf("Average salary of active employees: $%.2f\n", 
            totalSalary/float64(activeCount))
    }
}