package main

import "fmt"

// Function to remove element at index
func removeAtIndex(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

// Function to insert element at index
func insertAtIndex(slice []int, index int, value int) []int {
    return append(slice[:index], append([]int{value}, slice[index:]...)...)
}

// Function to get every nth element
func getEveryNth(slice []int, n int) []int {
    var result []int
    for i := 0; i < len(slice); i += n {
        result = append(result, slice[i])
    }
    return result
}

// Function to reverse a slice
func reverseSlice(slice []int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[len(slice)-1-i] = v
    }
    return result
}

// Function to find all indices of a value
func findAllIndices(slice []int, value int) []int {
    var indices []int
    for i, v := range slice {
        if v == value {
            indices = append(indices, i)
        }
    }
    return indices
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Printf("Original: %v\n", numbers)
    
    // Remove element at index 2
    removed := removeAtIndex(numbers, 2)
    fmt.Printf("After removing index 2: %v\n", removed)
    
    // Insert element at index 3
    inserted := insertAtIndex(numbers, 3, 99)
    fmt.Printf("After inserting 99 at index 3: %v\n", inserted)
    
    // Get every 2nd element
    every2nd := getEveryNth(numbers, 2)
    fmt.Printf("Every 2nd element: %v\n", every2nd)
    
    // Reverse the slice
    reversed := reverseSlice(numbers)
    fmt.Printf("Reversed: %v\n", reversed)
    
    // Find all indices of value 5
    indices := findAllIndices([]int{1, 5, 3, 5, 5, 7}, 5)
    fmt.Printf("Indices of 5 in [1,5,3,5,5,7]: %v\n", indices)
}