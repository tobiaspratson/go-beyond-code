package main

import (
	"fmt"
	"math"
	"time"
)

// Integer vs float operations
func benchmarkOperations() {
	// Integer operations
	start := time.Now()
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i * i
	}
	intTime := time.Since(start)

	// Float operations
	start = time.Now()
	floatSum := 0.0
	for i := 0; i < 10000000; i++ {
		floatSum += float64(i) * float64(i)
	}
	floatTime := time.Since(start)

	fmt.Printf("Integer operations time: %v\n", intTime)
	fmt.Printf("Float operations time: %v\n", floatTime)
	fmt.Printf("Integer sum: %d\n", sum)
	fmt.Printf("Float sum: %.0f\n", floatSum)
}

func main() {
	// Benchmark different power calculations
	base := 2.0
	exponent := 10.0

	// Method 1: math.Pow
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = math.Pow(base, exponent)
	}
	powTime := time.Since(start)

	// Method 2: Repeated multiplication (for integer exponents)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		result := 1.0
		for j := 0; j < int(exponent); j++ {
			result *= base
		}
		_ = result
	}
	multTime := time.Since(start)

	fmt.Printf("math.Pow time: %v\n", powTime)
	fmt.Printf("Repeated multiplication time: %v\n", multTime)

	benchmarkOperations()
}
