package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Employee struct {
    ID       int
    Name     string
    Position string
    Salary   float64
    Department string
}

func parseEmployee(record []string) (*Employee, error) {
    if len(record) != 5 {
        return nil, fmt.Errorf("invalid record length: expected 5, got %d", len(record))
    }
    
    id, err := strconv.Atoi(strings.TrimSpace(record[0]))
    if err != nil {
        return nil, fmt.Errorf("invalid ID: %v", err)
    }
    
    salary, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
    if err != nil {
        return nil, fmt.Errorf("invalid salary: %v", err)
    }
    
    return &Employee{
        ID:         id,
        Name:       strings.TrimSpace(record[1]),
        Position:   strings.TrimSpace(record[2]),
        Salary:     salary,
        Department: strings.TrimSpace(record[4]),
    }, nil
}

func processCSVFile(filename string) ([]Employee, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    reader.TrimLeadingSpace = true // Trim leading spaces
    
    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("failed to read CSV: %w", err)
    }
    
    var employees []Employee
    for i, record := range records {
        if i == 0 {
            // Skip header row
            continue
        }
        
        employee, err := parseEmployee(record)
        if err != nil {
            fmt.Printf("Warning: skipping row %d: %v\n", i+1, err)
            continue
        }
        
        employees = append(employees, *employee)
    }
    
    return employees, nil
}

func main() {
    employees, err := processCSVFile("employees.csv")
    if err != nil {
        fmt.Printf("Error processing CSV: %v\n", err)
        return
    }
    
    fmt.Printf("Processed %d employees:\n", len(employees))
    for _, emp := range employees {
        fmt.Printf("ID: %d, Name: %s, Position: %s, Salary: $%.2f, Dept: %s\n",
            emp.ID, emp.Name, emp.Position, emp.Salary, emp.Department)
    }
    
    // Calculate statistics
    totalSalary := 0.0
    deptCount := make(map[string]int)
    
    for _, emp := range employees {
        totalSalary += emp.Salary
        deptCount[emp.Department]++
    }
    
    fmt.Printf("\nStatistics:\n")
    fmt.Printf("Average salary: $%.2f\n", totalSalary/float64(len(employees)))
    fmt.Printf("Department distribution:\n")
    for dept, count := range deptCount {
        fmt.Printf("  %s: %d employees\n", dept, count)
    }
}