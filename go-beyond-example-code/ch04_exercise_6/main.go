package main

import "fmt"

// Test 3: Safe comparison function
func nearlyEqual(a, b, epsilon float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}

func main() {
	// Test 1: Simple addition
	a := 0.1
	b := 0.2
	c := a + b
	fmt.Printf("0.1 + 0.2 = %.17f\n", c)
	fmt.Printf("Is 0.1 + 0.2 == 0.3? %t\n", c == 0.3)

	// Test 2: Accumulation error
	var sum float64
	for i := 0; i < 10; i++ {
		sum += 0.1
	}
	fmt.Printf("0.1 added 10 times = %.17f\n", sum)

	fmt.Printf("0.1 + 0.2 ≈ 0.3? %t\n", nearlyEqual(c, 0.3, 1e-9))
	fmt.Printf("Sum ≈ 1.0? %t\n", nearlyEqual(sum, 1.0, 1e-9))
}
