package main

import "fmt"

// Function to find maximum value in array
func findMax(arr [5]int) int {
    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }
    return max
}

// Function to find minimum value in array
func findMin(arr [5]int) int {
    min := arr[0]
    for _, value := range arr {
        if value < min {
            min = value
        }
    }
    return min
}

// Function to calculate sum
func calculateSum(arr [5]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}

// Function to check if array contains value
func contains(arr [5]int, target int) bool {
    for _, value := range arr {
        if value == target {
            return true
        }
    }
    return false
}

// Function to reverse array (returns new array)
func reverse(arr [5]int) [5]int {
    var result [5]int
    for i, value := range arr {
        result[len(arr)-1-i] = value
    }
    return result
}

func main() {
    numbers := [5]int{10, 25, 5, 40, 15}
    
    fmt.Printf("Original array: %v\n", numbers)
    fmt.Printf("Maximum: %d\n", findMax(numbers))
    fmt.Printf("Minimum: %d\n", findMin(numbers))
    fmt.Printf("Sum: %d\n", calculateSum(numbers))
    fmt.Printf("Contains 25: %t\n", contains(numbers, 25))
    fmt.Printf("Contains 99: %t\n", contains(numbers, 99))
    fmt.Printf("Reversed: %v\n", reverse(numbers))
}