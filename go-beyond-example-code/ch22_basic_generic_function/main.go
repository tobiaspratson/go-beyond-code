package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generic function that works with any ordered type
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Use with integers
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))

	// Use with floats
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))

	// Use with strings
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))
}
