package main

import "fmt"

func main() {
    // Map of maps (2D map)
    students := make(map[string]map[string]int)
    
    // Add student data
    students["Alice"] = map[string]int{
        "math":    95,
        "science": 87,
        "english": 92,
    }
    
    students["Bob"] = map[string]int{
        "math":    78,
        "science": 85,
        "english": 90,
    }
    
    // Safe access to nested data
    if aliceScores, exists := students["Alice"]; exists {
        if mathScore, mathExists := aliceScores["math"]; mathExists {
            fmt.Printf("Alice's math score: %d\n", mathScore)
        } else {
            fmt.Println("Alice doesn't have a math score")
        }
    } else {
        fmt.Println("Alice not found")
    }
    
    // Iterate over nested map
    fmt.Println("\nAll student scores:")
    for student, subjects := range students {
        fmt.Printf("\n%s's scores:\n", student)
        for subject, score := range subjects {
            fmt.Printf("  %s: %d\n", subject, score)
        }
    }
    
    // Add new student with scores
    addStudentScores(students, "Charlie", map[string]int{
        "math":    88,
        "science": 92,
        "english": 85,
    })
    
    // Calculate average for each student
    fmt.Println("\nStudent averages:")
    for student, subjects := range students {
        avg := calculateAverage(subjects)
        fmt.Printf("%s: %.2f\n", student, avg)
    }
}

// Helper function to add student scores
func addStudentScores(students map[string]map[string]int, name string, scores map[string]int) {
    students[name] = scores
}

// Helper function to calculate average
func calculateAverage(scores map[string]int) float64 {
    if len(scores) == 0 {
        return 0.0
    }
    
    sum := 0
    for _, score := range scores {
        sum += score
    }
    
    return float64(sum) / float64(len(scores))
}