package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Constraint for numeric types that support arithmetic operations
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Generic function for numeric operations
func Add[T Numeric](a, b T) T {
	return a + b
}

func Multiply[T Numeric](a, b T) T {
	return a * b
}

func Subtract[T Numeric](a, b T) T {
	return a - b
}

// Constraint for comparable types (built-in constraint)
type Comparable interface {
	comparable
}

// Generic function for finding maximum (works with any ordered type)
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic function for sorting
func Sort[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	// Simple bubble sort
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// Generic function for finding minimum
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic function for clamping a value between min and max
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func main() {
	// Test numeric operations
	fmt.Printf("Add(10, 20): %d\n", Add(10, 20))
	fmt.Printf("Add(3.14, 2.71): %.2f\n", Add(3.14, 2.71))
	fmt.Printf("Multiply(5, 6): %d\n", Multiply(5, 6))
	fmt.Printf("Multiply(2.5, 4.0): %.2f\n", Multiply(2.5, 4.0))
	fmt.Printf("Subtract(10, 3): %d\n", Subtract(10, 3))
	fmt.Printf("Subtract(5.5, 2.3): %.2f\n", Subtract(5.5, 2.3))

	// Test comparable operations
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))

	// Test ordering operations
	fmt.Printf("Min(10, 20): %d\n", Min(10, 20))
	fmt.Printf("Min(3.14, 2.71): %.2f\n", Min(3.14, 2.71))
	fmt.Printf("Min(\"apple\", \"banana\"): %s\n", Min("apple", "banana"))

	// Test clamping
	fmt.Printf("Clamp(15, 10, 20): %d\n", Clamp(15, 10, 20))
	fmt.Printf("Clamp(5, 10, 20): %d\n", Clamp(5, 10, 20))
	fmt.Printf("Clamp(25, 10, 20): %d\n", Clamp(25, 10, 20))

	// Test sorting
	numbers := []int{5, 2, 8, 1, 9, 3}
	sortedNumbers := Sort(numbers)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Sorted: %v\n", sortedNumbers)

	words := []string{"banana", "apple", "cherry", "date"}
	sortedWords := Sort(words)
	fmt.Printf("Original: %v\n", words)
	fmt.Printf("Sorted: %v\n", sortedWords)

	floats := []float64{3.14, 2.71, 1.41, 1.73}
	sortedFloats := Sort(floats)
	fmt.Printf("Original floats: %v\n", floats)
	fmt.Printf("Sorted floats: %v\n", sortedFloats)
}
