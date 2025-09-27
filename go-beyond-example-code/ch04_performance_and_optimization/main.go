package main

import (
	"fmt"
	"time"
)

// Power of 2 check
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// Fast multiplication by powers of 2
func multiplyByPowerOfTwo(n, power int) int {
	return n << power // Equivalent to n * (2^power)
}

func divideByPowerOfTwo(n, power int) int {
	return n >> power // Equivalent to n / (2^power)
}

func main() {
	// Benchmark bit operations vs arithmetic
	const iterations = 10000000

	// Method 1: Using division and modulo
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = (i / 2) % 2 // Check if even
	}
	divTime := time.Since(start)

	// Method 2: Using bit operations
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = i & 1 // Check if even (bit 0)
	}
	bitTime := time.Since(start)

	fmt.Printf("Division method: %v\n", divTime)
	fmt.Printf("Bit operation method: %v\n", bitTime)
	fmt.Printf("Bit operations are %.1fx faster\n",
		float64(divTime)/float64(bitTime))

	// Test power of 2 function
	testNumbers := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 3, 5, 7, 9, 15}

	fmt.Printf("\nPower of 2 check:\n")
	for _, num := range testNumbers {
		fmt.Printf("%d is power of 2: %t\n", num, isPowerOfTwo(num))
	}

	fmt.Printf("\nFast operations:\n")
	fmt.Printf("5 * 8 = %d (using bits: %d)\n", 5*8, multiplyByPowerOfTwo(5, 3))
	fmt.Printf("40 / 8 = %d (using bits: %d)\n", 40/8, divideByPowerOfTwo(40, 3))
}
