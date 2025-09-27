package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generic function to find the maximum value in a slice
func Max[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Generic function to find the minimum value in a slice
func Min[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Generic function to check if a slice contains a value
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Generic function to find the index of a value
func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Generic function to find all indices of a value
func AllIndicesOf[T comparable](slice []T, value T) []int {
	var indices []int
	for i, v := range slice {
		if v == value {
			indices = append(indices, i)
		}
	}
	return indices
}

// Generic function to count occurrences of a value
func Count[T comparable](slice []T, value T) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

// Generic function to remove duplicates from a slice
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
	// Test with integers
	intSlice := []int{5, 2, 8, 1, 9, 3, 8, 2}
	fmt.Printf("Max of %v: %d\n", intSlice, Max(intSlice))
	fmt.Printf("Min of %v: %d\n", intSlice, Min(intSlice))
	fmt.Printf("Contains 8: %t\n", Contains(intSlice, 8))
	fmt.Printf("Contains 6: %t\n", Contains(intSlice, 6))
	fmt.Printf("Index of 8: %d\n", IndexOf(intSlice, 8))
	fmt.Printf("All indices of 8: %v\n", AllIndicesOf(intSlice, 8))
	fmt.Printf("Count of 8: %d\n", Count(intSlice, 8))
	fmt.Printf("Unique values: %v\n", Unique(intSlice))

	// Test with strings
	stringSlice := []string{"apple", "banana", "cherry", "date", "banana"}
	fmt.Printf("Max of %v: %s\n", stringSlice, Max(stringSlice))
	fmt.Printf("Min of %v: %s\n", stringSlice, Min(stringSlice))
	fmt.Printf("Contains 'banana': %t\n", Contains(stringSlice, "banana"))
	fmt.Printf("Contains 'grape': %t\n", Contains(stringSlice, "grape"))
	fmt.Printf("Index of 'banana': %d\n", IndexOf(stringSlice, "banana"))
	fmt.Printf("All indices of 'banana': %v\n", AllIndicesOf(stringSlice, "banana"))
	fmt.Printf("Count of 'banana': %d\n", Count(stringSlice, "banana"))
	fmt.Printf("Unique values: %v\n", Unique(stringSlice))

	// Test with floats
	floatSlice := []float64{3.14, 2.71, 1.41, 0.57, 2.71}
	fmt.Printf("Max of %v: %.2f\n", floatSlice, Max(floatSlice))
	fmt.Printf("Min of %v: %.2f\n", floatSlice, Min(floatSlice))
	fmt.Printf("Unique values: %v\n", Unique(floatSlice))
}
