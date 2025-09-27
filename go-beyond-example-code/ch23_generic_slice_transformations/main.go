package main

import "fmt"

// Generic function to map over a slice
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Generic function to filter a slice
func Filter[T any](slice []T, fn func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

// Generic function to reduce a slice
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

// Generic function to reverse a slice
func Reverse[T any](slice []T) []T {
    result := make([]T, len(slice))
    for i, v := range slice {
        result[len(slice)-1-i] = v
    }
    return result
}

func main() {
    // Test Map function
    numbers := []int{1, 2, 3, 4, 5}
    doubled := Map(numbers, func(x int) int { return x * 2 })
    fmt.Printf("Original: %v\n", numbers)
    fmt.Printf("Doubled: %v\n", doubled)
    
    // Test Map with different types
    strings := []string{"hello", "world", "golang"}
    lengths := Map(strings, func(s string) int { return len(s) })
    fmt.Printf("Strings: %v\n", strings)
    fmt.Printf("Lengths: %v\n", lengths)
    
    // Test Filter function
    evenNumbers := Filter(numbers, func(x int) bool { return x%2 == 0 })
    fmt.Printf("Even numbers: %v\n", evenNumbers)
    
    longStrings := Filter(strings, func(s string) bool { return len(s) > 4 })
    fmt.Printf("Long strings: %v\n", longStrings)
    
    // Test Reduce function
    sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
    fmt.Printf("Sum of %v: %d\n", numbers, sum)
    
    product := Reduce(numbers, 1, func(acc, x int) int { return acc * x })
    fmt.Printf("Product of %v: %d\n", numbers, product)
    
    // Test Reverse function
    reversed := Reverse(numbers)
    fmt.Printf("Original: %v\n", numbers)
    fmt.Printf("Reversed: %v\n", reversed)
}