package main

import "fmt"

// Generic function to chunk a slice into smaller slices
func Chunk[T any](slice []T, size int) [][]T {
    if size <= 0 {
        return nil
    }
    
    var chunks [][]T
    for i := 0; i < len(slice); i += size {
        end := i + size
        if end > len(slice) {
            end = len(slice)
        }
        chunks = append(chunks, slice[i:end])
    }
    return chunks
}

// Generic function to flatten a slice of slices
func Flatten[T any](slices [][]T) []T {
    var result []T
    for _, slice := range slices {
        result = append(result, slice...)
    }
    return result
}

// Generic function to zip two slices together
func Zip[T, U any](slice1 []T, slice2 []U) []struct{ First T; Second U } {
    minLen := len(slice1)
    if len(slice2) < minLen {
        minLen = len(slice2)
    }
    
    result := make([]struct{ First T; Second U }, minLen)
    for i := 0; i < minLen; i++ {
        result[i] = struct{ First T; Second U }{slice1[i], slice2[i]}
    }
    return result
}

// Generic function to partition a slice based on a predicate
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
    var trueSlice, falseSlice []T
    for _, v := range slice {
        if predicate(v) {
            trueSlice = append(trueSlice, v)
        } else {
            falseSlice = append(falseSlice, v)
        }
    }
    return trueSlice, falseSlice
}

// Generic function to take the first n elements
func Take[T any](slice []T, n int) []T {
    if n <= 0 {
        return nil
    }
    if n > len(slice) {
        n = len(slice)
    }
    return slice[:n]
}

// Generic function to drop the first n elements
func Drop[T any](slice []T, n int) []T {
    if n <= 0 {
        return slice
    }
    if n >= len(slice) {
        return nil
    }
    return slice[n:]
}

func main() {
    // Test Chunk function
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    chunks := Chunk(numbers, 3)
    fmt.Printf("Chunked %v into chunks of 3: %v\n", numbers, chunks)
    
    // Test Flatten function
    nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
    flattened := Flatten(nested)
    fmt.Printf("Flattened %v: %v\n", nested, flattened)
    
    // Test Zip function
    names := []string{"Alice", "Bob", "Charlie"}
    ages := []int{25, 30, 35}
    zipped := Zip(names, ages)
    fmt.Printf("Zipped names and ages: %v\n", zipped)
    
    // Test Partition function
    numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evens, odds := Partition(numbers, func(x int) bool { return x%2 == 0 })
    fmt.Printf("Partitioned %v into evens: %v, odds: %v\n", numbers, evens, odds)
    
    // Test Take and Drop functions
    numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    firstThree := Take(numbers, 3)
    afterThree := Drop(numbers, 3)
    fmt.Printf("First 3 of %v: %v\n", numbers, firstThree)
    fmt.Printf("After dropping 3 from %v: %v\n", numbers, afterThree)
}