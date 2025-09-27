package main

import (
    "fmt"
    "sort"
)

func main() {
    // Map to store student grades
    grades := make(map[string][]int)
    
    // Add grades for students
    grades["Alice"] = []int{95, 87, 92}
    grades["Bob"] = []int{78, 85, 90}
    grades["Charlie"] = []int{88, 91, 89}
    
    // Calculate averages
    averages := make(map[string]float64)
    for student, gradeList := range grades {
        sum := 0
        for _, grade := range gradeList {
            sum += grade
        }
        averages[student] = float64(sum) / float64(len(gradeList))
    }
    
    // Display results
    fmt.Println("Student Averages:")
    for student, avg := range averages {
        fmt.Printf("%s: %.2f\n", student, avg)
    }
    
    // Find highest average
    var topStudent string
    var topAverage float64
    for student, avg := range averages {
        if avg > topAverage {
            topAverage = avg
            topStudent = student
        }
    }
    fmt.Printf("\nTop student: %s with %.2f\n", topStudent, topAverage)
}