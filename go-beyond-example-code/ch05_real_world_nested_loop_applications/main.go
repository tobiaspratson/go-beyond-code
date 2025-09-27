package main

import "fmt"

func main() {
    fmt.Println("=== 2D Array Processing ===")
    // Working with 2D data
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("Matrix:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%2d ", matrix[i][j])
        }
        fmt.Println()
    }
    
    fmt.Println("\nMatrix transpose:")
    for i := 0; i < len(matrix[0]); i++ {
        for j := 0; j < len(matrix); j++ {
            fmt.Printf("%2d ", matrix[j][i])
        }
        fmt.Println()
    }
    
    fmt.Println("\n=== String Processing ===")
    // Processing 2D string data
    words := [][]string{
        {"hello", "world"},
        {"go", "programming"},
        {"nested", "loops"},
    }
    
    for i, row := range words {
        fmt.Printf("Row %d: ", i+1)
        for j, word := range row {
            fmt.Printf("%s", word)
            if j < len(row)-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
    
    fmt.Println("\n=== Data Analysis ===")
    // Finding maximum value in 2D array
    scores := [][]int{
        {85, 92, 78},
        {91, 88, 95},
        {76, 89, 82},
    }
    
    maxScore := 0
    maxRow, maxCol := 0, 0
    
    for i := 0; i < len(scores); i++ {
        for j := 0; j < len(scores[i]); j++ {
            if scores[i][j] > maxScore {
                maxScore = scores[i][j]
                maxRow, maxCol = i, j
            }
        }
    }
    
    fmt.Printf("Maximum score: %d at position [%d][%d]\n", maxScore, maxRow, maxCol)
}