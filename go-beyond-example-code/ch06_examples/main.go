package main

import (
    "fmt"
    "strconv"
    "sort"
)

type Student struct {
    Name   string
    Grades []float64
}

func (s *Student) AddGrade(grade float64) {
    s.Grades = append(s.Grades, grade)
}

func (s *Student) GetAverage() float64 {
    if len(s.Grades) == 0 {
        return 0.0
    }
    
    sum := 0.0
    for _, grade := range s.Grades {
        sum += grade
    }
    return sum / float64(len(s.Grades))
}

func (s *Student) GetHighest() float64 {
    if len(s.Grades) == 0 {
        return 0.0
    }
    
    highest := s.Grades[0]
    for _, grade := range s.Grades {
        if grade > highest {
            highest = grade
        }
    }
    return highest
}

func (s *Student) GetLowest() float64 {
    if len(s.Grades) == 0 {
        return 0.0
    }
    
    lowest := s.Grades[0]
    for _, grade := range s.Grades {
        if grade < lowest {
            lowest = grade
        }
    }
    return lowest
}

func (s *Student) GetSortedGrades() []float64 {
    sorted := make([]float64, len(s.Grades))
    copy(sorted, s.Grades)
    sort.Float64s(sorted)
    return sorted
}

func (s *Student) GetGradeDistribution() map[string]int {
    distribution := map[string]int{
        "A": 0, "B": 0, "C": 0, "D": 0, "F": 0,
    }
    
    for _, grade := range s.Grades {
        switch {
        case grade >= 90:
            distribution["A"]++
        case grade >= 80:
            distribution["B"]++
        case grade >= 70:
            distribution["C"]++
        case grade >= 60:
            distribution["D"]++
        default:
            distribution["F"]++
        }
    }
    
    return distribution
}

func main() {
    var students []Student
    var input string
    
    fmt.Println("=== Student Grade Management System ===")
    
    for {
        fmt.Print("\nEnter student name (or 'quit' to exit): ")
        fmt.Scanln(&input)
        
        if input == "quit" {
            break
        }
        
        student := Student{Name: input}
        
        fmt.Printf("Enter grades for %s (type 'done' when finished):\n", student.Name)
        
        for {
            fmt.Print("Grade: ")
            fmt.Scanln(&input)
            
            if input == "done" {
                break
            }
            
            if grade, err := strconv.ParseFloat(input, 64); err == nil {
                if grade >= 0 && grade <= 100 {
                    student.AddGrade(grade)
                } else {
                    fmt.Println("Grade must be between 0 and 100!")
                }
            } else {
                fmt.Println("Invalid grade! Please enter a number.")
            }
        }
        
        if len(student.Grades) > 0 {
            students = append(students, student)
        }
    }
    
    if len(students) == 0 {
        fmt.Println("No students entered.")
        return
    }
    
    // Display results for each student
    for _, student := range students {
        fmt.Printf("\n=== %s's Grade Report ===\n", student.Name)
        fmt.Printf("Grades: %v\n", student.Grades)
        fmt.Printf("Sorted: %v\n", student.GetSortedGrades())
        fmt.Printf("Average: %.2f\n", student.GetAverage())
        fmt.Printf("Highest: %.2f\n", student.GetHighest())
        fmt.Printf("Lowest: %.2f\n", student.GetLowest())
        
        distribution := student.GetGradeDistribution()
        fmt.Printf("Grade Distribution: %v\n", distribution)
    }
    
    // Class statistics
    fmt.Println("\n=== Class Statistics ===")
    var allGrades []float64
    for _, student := range students {
        allGrades = append(allGrades, student.Grades...)
    }
    
    if len(allGrades) > 0 {
        classSum := 0.0
        for _, grade := range allGrades {
            classSum += grade
        }
        classAverage := classSum / float64(len(allGrades))
        
        sort.Float64s(allGrades)
        fmt.Printf("Total students: %d\n", len(students))
        fmt.Printf("Total grades: %d\n", len(allGrades))
        fmt.Printf("Class average: %.2f\n", classAverage)
        fmt.Printf("Class highest: %.2f\n", allGrades[len(allGrades)-1])
        fmt.Printf("Class lowest: %.2f\n", allGrades[0])
    }
}