package main

import "fmt"

// Employee interface
type Employee interface {
    GetName() string
    GetID() string
    CalculateSalary() float64
    GetInfo() string
}

// Base employee struct
type BaseEmployee struct {
    Name     string
    ID       string
    BaseSalary float64
}

func (e BaseEmployee) GetName() string {
    return e.Name
}

func (e BaseEmployee) GetID() string {
    return e.ID
}

// Full-time employee
type FullTimeEmployee struct {
    BaseEmployee
    Bonus float64
}

func (f FullTimeEmployee) CalculateSalary() float64 {
    return f.BaseSalary + f.Bonus
}

func (f FullTimeEmployee) GetInfo() string {
    return fmt.Sprintf("Full-time: %s (ID: %s) - Salary: $%.2f", 
        f.Name, f.ID, f.CalculateSalary())
}

// Part-time employee
type PartTimeEmployee struct {
    BaseEmployee
    HoursWorked float64
    HourlyRate  float64
}

func (p PartTimeEmployee) CalculateSalary() float64 {
    return p.HoursWorked * p.HourlyRate
}

func (p PartTimeEmployee) GetInfo() string {
    return fmt.Sprintf("Part-time: %s (ID: %s) - Salary: $%.2f (%.1f hours)", 
        p.Name, p.ID, p.CalculateSalary(), p.HoursWorked)
}

// Commission employee
type CommissionEmployee struct {
    BaseEmployee
    SalesAmount float64
    CommissionRate float64
}

func (c CommissionEmployee) CalculateSalary() float64 {
    return c.BaseSalary + (c.SalesAmount * c.CommissionRate)
}

func (c CommissionEmployee) GetInfo() string {
    return fmt.Sprintf("Commission: %s (ID: %s) - Salary: $%.2f (Sales: $%.2f)", 
        c.Name, c.ID, c.CalculateSalary(), c.SalesAmount)
}

// Company with employee management
type Company struct {
    Name      string
    Employees []Employee
}

func (c *Company) AddEmployee(employee Employee) {
    c.Employees = append(c.Employees, employee)
    fmt.Printf("Added employee: %s\n", employee.GetInfo())
}

func (c Company) GetTotalPayroll() float64 {
    total := 0.0
    for _, emp := range c.Employees {
        total += emp.CalculateSalary()
    }
    return total
}

func (c Company) ListEmployees() {
    fmt.Printf("\n=== %s Employees ===\n", c.Name)
    for i, emp := range c.Employees {
        fmt.Printf("%d. %s\n", i+1, emp.GetInfo())
    }
}

func (c Company) GetPayrollStats() {
    if len(c.Employees) == 0 {
        fmt.Println("No employees")
        return
    }
    
    total := c.GetTotalPayroll()
    average := total / float64(len(c.Employees))
    
    fmt.Printf("\n=== Payroll Statistics ===\n")
    fmt.Printf("Total employees: %d\n", len(c.Employees))
    fmt.Printf("Total payroll: $%.2f\n", total)
    fmt.Printf("Average salary: $%.2f\n", average)
}

func main() {
    // Create company
    company := Company{Name: "TechCorp Inc."}
    
    // Add different types of employees
    company.AddEmployee(FullTimeEmployee{
        BaseEmployee: BaseEmployee{
            Name:      "Alice Johnson",
            ID:        "EMP001",
            BaseSalary: 75000,
        },
        Bonus: 5000,
    })
    
    company.AddEmployee(PartTimeEmployee{
        BaseEmployee: BaseEmployee{
            Name:      "Bob Smith",
            ID:        "EMP002",
            BaseSalary: 0, // Not used for part-time
        },
        HoursWorked: 20,
        HourlyRate:  25,
    })
    
    company.AddEmployee(CommissionEmployee{
        BaseEmployee: BaseEmployee{
            Name:      "Carol Davis",
            ID:        "EMP003",
            BaseSalary: 30000,
        },
        SalesAmount:    100000,
        CommissionRate: 0.05, // 5% commission
    })
    
    company.AddEmployee(FullTimeEmployee{
        BaseEmployee: BaseEmployee{
            Name:      "David Wilson",
            ID:        "EMP004",
            BaseSalary: 80000,
        },
        Bonus: 3000,
    })
    
    // Display all employees
    company.ListEmployees()
    
    // Show payroll statistics
    company.GetPayrollStats()
}