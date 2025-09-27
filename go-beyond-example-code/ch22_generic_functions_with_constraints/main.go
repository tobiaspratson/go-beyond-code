package main

import "fmt"

// Constraint for types that support ordering
// The ~ symbol means "underlying type" - allows custom types based on these primitives
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 | ~string
}

// Generic function for finding maximum with ordering constraint
func MaxOrdered[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Generic function for sorting a slice
func Sort[T Ordered](slice []T) []T {
    result := make([]T, len(slice))
    copy(result, slice)
    
    // Simple bubble sort implementation
    for i := 0; i < len(result); i++ {
        for j := 0; j < len(result)-1-i; j++ {
            if result[j] > result[j+1] {
                result[j], result[j+1] = result[j+1], result[j]
            }
        }
    }
    
    return result
}

// Generic function for binary search (requires sorted slice)
func BinarySearch[T Ordered](slice []T, target T) int {
    left, right := 0, len(slice)-1
    
    for left <= right {
        mid := (left + right) / 2
        if slice[mid] == target {
            return mid
        } else if slice[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

// Generic function for finding the median of a slice
func Median[T Ordered](slice []T) T {
    if len(slice) == 0 {
        var zero T
        return zero
    }
    
    sorted := Sort(slice)
    mid := len(sorted) / 2
    
    if len(sorted)%2 == 0 {
        // For even length, return the average of two middle elements
        // Note: This assumes the type supports division, which Ordered doesn't guarantee
        return sorted[mid-1] // Simplified for this example
    }
    
    return sorted[mid]
}

// Generic function for removing duplicates from a slice
func Unique[T comparable](slice []T) []T {
    seen := make(map[T]bool)
    var result []T
    
    for _, v := range slice {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    
    return result
}

func main() {
    // Test MaxOrdered function with different types
    fmt.Printf("MaxOrdered(10, 20): %d\n", MaxOrdered(10, 20))
    fmt.Printf("MaxOrdered(3.14, 2.71): %.2f\n", MaxOrdered(3.14, 2.71))
    fmt.Printf("MaxOrdered(\"apple\", \"banana\"): %s\n", MaxOrdered("apple", "banana"))
    
    // Test Sort function with integers
    numbers := []int{5, 2, 8, 1, 9, 3}
    sortedNumbers := Sort(numbers)
    fmt.Printf("Original numbers: %v\n", numbers)
    fmt.Printf("Sorted numbers: %v\n", sortedNumbers)
    
    // Test Sort function with strings
    words := []string{"banana", "apple", "cherry", "date"}
    sortedWords := Sort(words)
    fmt.Printf("Original words: %v\n", words)
    fmt.Printf("Sorted words: %v\n", sortedWords)
    
    // Test Sort function with floats
    floats := []float64{3.14, 2.71, 1.41, 1.73}
    sortedFloats := Sort(floats)
    fmt.Printf("Original floats: %v\n", floats)
    fmt.Printf("Sorted floats: %v\n", sortedFloats)
    
    // Test BinarySearch function
    sortedNumbers = Sort(numbers)
    index := BinarySearch(sortedNumbers, 8)
    fmt.Printf("BinarySearch(sortedNumbers, 8): %d\n", index)
    
    index = BinarySearch(sortedNumbers, 6)
    fmt.Printf("BinarySearch(sortedNumbers, 6): %d\n", index)
    
    // Test Median function
    median := Median(numbers)
    fmt.Printf("Median of %v: %d\n", numbers, median)
    
    // Test Unique function
    duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
    unique := Unique(duplicates)
    fmt.Printf("Original: %v\n", duplicates)
    fmt.Printf("Unique: %v\n", unique)
    
    stringDuplicates := []string{"apple", "banana", "apple", "cherry", "banana"}
    uniqueStrings := Unique(stringDuplicates)
    fmt.Printf("Original strings: %v\n", stringDuplicates)
    fmt.Printf("Unique strings: %v\n", uniqueStrings)
}