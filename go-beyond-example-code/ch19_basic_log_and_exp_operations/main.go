package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Exponential Functions ===")
	// Natural exponential (e^x)
	fmt.Printf("Exp(1): %.6f (e^1)\n", math.Exp(1))
	fmt.Printf("Exp(2): %.6f (e^2)\n", math.Exp(2))
	fmt.Printf("Exp(0): %.6f (e^0 = 1)\n", math.Exp(0))
	fmt.Printf("Exp(-1): %.6f (e^-1)\n", math.Exp(-1))

	// Base 2 exponential
	fmt.Printf("Exp2(3): %.6f (2^3)\n", math.Exp2(3))
	fmt.Printf("Exp2(0): %.6f (2^0 = 1)\n", math.Exp2(0))
	fmt.Printf("Exp2(-1): %.6f (2^-1 = 0.5)\n", math.Exp2(-1))

	// Base 10 exponential
	fmt.Printf("Pow(10, 2): %.6f (10^2)\n", math.Pow(10, 2))
	fmt.Printf("Pow(10, 0): %.6f (10^0 = 1)\n", math.Pow(10, 0))
	fmt.Printf("Pow(10, -1): %.6f (10^-1 = 0.1)\n", math.Pow(10, -1))

	fmt.Println("\n=== Logarithmic Functions ===")
	// Natural logarithm (ln)
	fmt.Printf("Log(2.718): %.6f (ln(e) ≈ 1)\n", math.Log(2.718))
	fmt.Printf("Log(1): %.6f (ln(1) = 0)\n", math.Log(1))
	fmt.Printf("Log(10): %.6f (ln(10))\n", math.Log(10))

	// Base 2 logarithm
	fmt.Printf("Log2(8): %.6f (log₂(8) = 3)\n", math.Log2(8))
	fmt.Printf("Log2(1): %.6f (log₂(1) = 0)\n", math.Log2(1))
	fmt.Printf("Log2(0.5): %.6f (log₂(0.5) = -1)\n", math.Log2(0.5))

	// Base 10 logarithm
	fmt.Printf("Log10(100): %.6f (log₁₀(100) = 2)\n", math.Log10(100))
	fmt.Printf("Log10(1): %.6f (log₁₀(1) = 0)\n", math.Log10(1))
	fmt.Printf("Log10(0.1): %.6f (log₁₀(0.1) = -1)\n", math.Log10(0.1))

	fmt.Println("\n=== Logarithmic Identities ===")
	// Verify logarithmic identities
	x, y := 2.0, 3.0
	fmt.Printf("Log(%f) + Log(%f) = %.6f\n", x, y, math.Log(x)+math.Log(y))
	fmt.Printf("Log(%f × %f) = %.6f\n", x, y, math.Log(x*y))
	fmt.Printf("Are they equal? %t\n", math.Abs((math.Log(x)+math.Log(y))-math.Log(x*y)) < 1e-10)

	fmt.Printf("Log(%f) - Log(%f) = %.6f\n", y, x, math.Log(y)-math.Log(x))
	fmt.Printf("Log(%f / %f) = %.6f\n", y, x, math.Log(y/x))
	fmt.Printf("Are they equal? %t\n", math.Abs((math.Log(y)-math.Log(x))-math.Log(y/x)) < 1e-10)

	fmt.Println("\n=== Change of Base Formula ===")
	// Custom base logarithms using change of base formula
	base := 3.0
	value := 27.0
	customLog := math.Log(value) / math.Log(base)
	fmt.Printf("Log₃(%.1f) = %.6f\n", value, customLog)
	fmt.Printf("Verification: 3^%.6f = %.6f\n", customLog, math.Pow(base, customLog))

	// Multiple bases
	bases := []float64{2, 3, 5, 10}
	testValue := 100.0
	fmt.Printf("Log₁₀₀ in different bases:\n")
	for _, b := range bases {
		logValue := math.Log(testValue) / math.Log(b)
		fmt.Printf("Log_%.0f(%.0f) = %.6f\n", b, testValue, logValue)
	}
}
