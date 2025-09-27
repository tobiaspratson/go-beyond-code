package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generic function for finding the maximum value
// T is the type parameter, constraints.Ordered is the constraint
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic function for finding the minimum value
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic function for swapping values
// Uses 'any' constraint - works with any type
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// Generic function for checking if a value exists in a slice
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Generic function for finding the index of a value in a slice
func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Generic function for reversing a slice
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Generic function for filtering a slice
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test Max function with different types
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))

	// Test Min function with different types
	fmt.Printf("Min(10, 20): %d\n", Min(10, 20))
	fmt.Printf("Min(3.14, 2.71): %.2f\n", Min(3.14, 2.71))
	fmt.Printf("Min(\"apple\", \"banana\"): %s\n", Min("apple", "banana"))

	// Test Swap function
	x, y := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
	Swap(&x, &y)
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)

	// Test Contains function
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains(numbers, 3): %t\n", Contains(numbers, 3))
	fmt.Printf("Contains(numbers, 6): %t\n", Contains(numbers, 6))

	words := []string{"apple", "banana", "cherry"}
	fmt.Printf("Contains(words, \"banana\"): %t\n", Contains(words, "banana"))
	fmt.Printf("Contains(words, \"grape\"): %t\n", Contains(words, "grape"))

	// Test IndexOf function
	fmt.Printf("IndexOf(numbers, 3): %d\n", IndexOf(numbers, 3))
	fmt.Printf("IndexOf(numbers, 6): %d\n", IndexOf(numbers, 6))

	// Test Reverse function
	original := []int{1, 2, 3, 4, 5}
	reversed := Reverse(original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test Filter function
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	longWords := Filter(words, func(s string) bool {
		return len(s) > 5
	})
	fmt.Printf("Long words: %v\n", longWords)
}
