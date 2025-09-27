package main

import "fmt"

func main() {
    // Create a 3x3 matrix
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("Matrix:")
    for i, row := range matrix {
        for j, val := range row {
            fmt.Printf("%d ", val)
            if j == len(row)-1 {
                fmt.Println()
            }
        }
    }
    
    // Calculate sum of all elements
    sum := 0
    for _, row := range matrix {
        for _, val := range row {
            sum += val
        }
    }
    fmt.Printf("Sum of all elements: %d\n", sum)
    
    // Calculate diagonal sum
    diagonalSum := 0
    for i := 0; i < len(matrix); i++ {
        diagonalSum += matrix[i][i]
    }
    fmt.Printf("Diagonal sum: %d\n", diagonalSum)
}