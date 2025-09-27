package main

import "fmt"

// Generic linear search - O(n) time complexity
func LinearSearch[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

// Generic search with custom predicate
func SearchBy[T any](slice []T, predicate func(T) bool) int {
    for i, v := range slice {
        if predicate(v) {
            return i
        }
    }
    return -1
}

// Generic search for all matching elements
func SearchAll[T comparable](slice []T, target T) []int {
    var indices []int
    for i, v := range slice {
        if v == target {
            indices = append(indices, i)
        }
    }
    return indices
}

// Generic search with custom predicate for all matches
func SearchAllBy[T any](slice []T, predicate func(T) bool) []int {
    var indices []int
    for i, v := range slice {
        if predicate(v) {
            indices = append(indices, i)
        }
    }
    return indices
}

func main() {
    // Test linear search
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    index := LinearSearch(numbers, 5)
    fmt.Printf("Linear search for 5 in %v: %d\n", numbers, index)
    
    index = LinearSearch(numbers, 15)
    fmt.Printf("Linear search for 15 in %v: %d\n", numbers, index)
    
    // Test search with predicate
    index = SearchBy(numbers, func(x int) bool { return x > 7 })
    fmt.Printf("First number > 7 in %v: %d\n", numbers, index)
    
    // Test search all
    numbers = []int{1, 2, 3, 2, 4, 2, 5}
    indices := SearchAll(numbers, 2)
    fmt.Printf("All indices of 2 in %v: %v\n", numbers, indices)
    
    // Test search all with predicate
    indices = SearchAllBy(numbers, func(x int) bool { return x%2 == 0 })
    fmt.Printf("All even numbers in %v: %v\n", numbers, indices)
    
    // Test with strings
    words := []string{"apple", "banana", "cherry", "date", "elderberry"}
    index = LinearSearch(words, "cherry")
    fmt.Printf("Linear search for 'cherry' in %v: %d\n", words, index)
    
    // Test with custom types
    type Person struct {
        Name string
        Age  int
    }
    
    people := []Person{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 35},
        {"Diana", 28},
    }
    
    index = SearchBy(people, func(p Person) bool { return p.Age > 30 })
    fmt.Printf("First person older than 30: %+v\n", people[index])
}