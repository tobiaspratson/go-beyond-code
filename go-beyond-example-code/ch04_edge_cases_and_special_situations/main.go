package main

import (
	"fmt"
	"math"
)

// Safe integer operations
func safeAdd(a, b int64) (int64, bool) {
	if a > 0 && b > math.MaxInt64-a {
		return 0, false // Overflow
	}
	if a < 0 && b < math.MinInt64-a {
		return 0, false // Underflow
	}
	return a + b, true
}

func main() {
	// Division by zero with floats
	fmt.Printf("1.0 / 0.0 = %f (infinity)\n", 1.0/0.0)
	fmt.Printf("-1.0 / 0.0 = %f (negative infinity)\n", -1.0/0.0)
	fmt.Printf("0.0 / 0.0 = %f (NaN)\n", 0.0/0.0)

	// Very large numbers
	large := math.MaxFloat64
	fmt.Printf("Max float64: %.2e\n", large)
	fmt.Printf("Max float64 + 1: %.2e (no change due to precision)\n", large+1)

	// Very small numbers
	small := math.SmallestNonzeroFloat64
	fmt.Printf("Smallest float64: %.2e\n", small)
	fmt.Printf("Smallest float64 / 2: %.2e (underflow to 0)\n", small/2)

	// Overflow in integer operations
	var maxInt int64 = math.MaxInt64
	fmt.Printf("Max int64: %d\n", maxInt)
	fmt.Printf("Max int64 + 1: %d (overflow!)\n", maxInt+1)

	result, ok := safeAdd(maxInt, 1)
	if !ok {
		fmt.Println("Addition would overflow!")
	} else {
		fmt.Printf("Safe addition: %d\n", result)
	}
}
