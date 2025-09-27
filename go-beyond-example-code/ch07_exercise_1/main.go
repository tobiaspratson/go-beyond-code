package main

import "fmt"

// Subject represents a course with grades
type Subject struct {
    Name   string
    Grades []float64
}

func (s *Subject) AddGrade(grade float64) {
    s.Grades = append(s.Grades, grade)
}

func (s Subject) Average() float64 {
    if len(s.Grades) == 0 {
        return 0.0
    }
    
    sum := 0.0
    for _, grade := range s.Grades {
        sum += grade
    }
    return sum / float64(len(s.Grades))
}

func (s Subject) GetInfo() string {
    return fmt.Sprintf("%s: %.2f (%.0f grades)", s.Name, s.Average(), float64(len(s.Grades)))
}

// Student with multiple subjects
type Student struct {
    Name      string
    ID        string
    Subjects  map[string]*Subject
}

func NewStudent(name, id string) *Student {
    return &Student{
        Name:     name,
        ID:       id,
        Subjects: make(map[string]*Subject),
    }
}

func (s *Student) AddSubject(subjectName string) {
    s.Subjects[subjectName] = &Subject{Name: subjectName}
}

func (s *Student) AddGrade(subjectName string, grade float64) bool {
    if subject, exists := s.Subjects[subjectName]; exists {
        subject.AddGrade(grade)
        return true
    }
    return false
}

func (s Student) GetSubjectAverage(subjectName string) float64 {
    if subject, exists := s.Subjects[subjectName]; exists {
        return subject.Average()
    }
    return 0.0
}

func (s Student) GetOverallAverage() float64 {
    if len(s.Subjects) == 0 {
        return 0.0
    }
    
    total := 0.0
    count := 0
    for _, subject := range s.Subjects {
        if len(subject.Grades) > 0 {
            total += subject.Average()
            count++
        }
    }
    
    if count == 0 {
        return 0.0
    }
    return total / float64(count)
}

func (s Student) GetInfo() string {
    return fmt.Sprintf("Student: %s (ID: %s), Overall Average: %.2f", 
        s.Name, s.ID, s.GetOverallAverage())
}

func (s Student) GetDetailedInfo() string {
    info := fmt.Sprintf("=== %s ===\n", s.GetInfo())
    for _, subject := range s.Subjects {
        if len(subject.Grades) > 0 {
            info += fmt.Sprintf("  %s\n", subject.GetInfo())
        }
    }
    return info
}

func main() {
    // Create student
    student := NewStudent("Alice Johnson", "STU001")
    
    // Add subjects
    student.AddSubject("Mathematics")
    student.AddSubject("Physics")
    student.AddSubject("Chemistry")
    student.AddSubject("English")
    
    // Add grades
    student.AddGrade("Mathematics", 85.5)
    student.AddGrade("Mathematics", 92.0)
    student.AddGrade("Mathematics", 78.5)
    
    student.AddGrade("Physics", 88.0)
    student.AddGrade("Physics", 91.5)
    student.AddGrade("Physics", 85.0)
    
    student.AddGrade("Chemistry", 82.0)
    student.AddGrade("Chemistry", 89.5)
    
    student.AddGrade("English", 90.0)
    student.AddGrade("English", 87.5)
    student.AddGrade("English", 93.0)
    
    // Display results
    fmt.Println(student.GetDetailedInfo())
    
    // Individual subject averages
    fmt.Println("Individual Subject Averages:")
    for subjectName := range student.Subjects {
        avg := student.GetSubjectAverage(subjectName)
        if avg > 0 {
            fmt.Printf("  %s: %.2f\n", subjectName, avg)
        }
    }
}