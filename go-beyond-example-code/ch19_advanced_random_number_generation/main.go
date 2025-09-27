package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Generate random number from normal distribution using Box-Muller transform
func NormalRandom(mean, stdDev float64) float64 {
	// Box-Muller transform
	u1 := rand.Float64()
	u2 := rand.Float64()

	// Avoid log(0)
	for u1 == 0 {
		u1 = rand.Float64()
	}

	z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0*stdDev + mean
}

// Generate random number from exponential distribution
func ExponentialRandom(lambda float64) float64 {
	u := rand.Float64()
	// Avoid log(0)
	for u == 0 || u == 1 {
		u = rand.Float64()
	}
	return -math.Log(1-u) / lambda
}

// Generate random number from uniform distribution in range [a, b]
func UniformRandom(a, b float64) float64 {
	return rand.Float64()*(b-a) + a
}

// Generate random permutation using Fisher-Yates shuffle
func RandomPermutation(n int) []int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}

	// Fisher-Yates shuffle
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		perm[i], perm[j] = perm[j], perm[i]
	}

	return perm
}

// Generate random sample without replacement
func RandomSample(population []int, sampleSize int) []int {
	if sampleSize > len(population) {
		sampleSize = len(population)
	}

	// Create a copy to avoid modifying original
	pop := make([]int, len(population))
	copy(pop, population)

	// Fisher-Yates shuffle for first sampleSize elements
	for i := 0; i < sampleSize; i++ {
		j := rand.Intn(len(pop)-i) + i
		pop[i], pop[j] = pop[j], pop[i]
	}

	return pop[:sampleSize]
}

// Generate random string
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Generate random password with specific requirements
func RandomPassword(length int, includeUppercase, includeNumbers, includeSymbols bool) string {
	charset := "abcdefghijklmnopqrstuvwxyz"

	if includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSymbols {
		charset += "!@#$%^&*()_+-=[]{}|;:,.<>?"
	}

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Monte Carlo simulation for estimating π
func EstimatePi(numSamples int) float64 {
	insideCircle := 0

	for i := 0; i < numSamples; i++ {
		x := rand.Float64()*2 - 1 // Random x in [-1, 1]
		y := rand.Float64()*2 - 1 // Random y in [-1, 1]

		if x*x+y*y <= 1 {
			insideCircle++
		}
	}

	return 4.0 * float64(insideCircle) / float64(numSamples)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Advanced Random Number Generation ===")

	// Normal distribution
	fmt.Println("\n--- Normal Distribution ---")
	mean := 0.0
	stdDev := 1.0
	fmt.Printf("Normal distribution (mean=%.1f, std=%.1f):\n", mean, stdDev)
	for i := 0; i < 10; i++ {
		fmt.Printf("%.4f ", NormalRandom(mean, stdDev))
	}
	fmt.Println()

	// Exponential distribution
	fmt.Println("\n--- Exponential Distribution ---")
	lambda := 1.0
	fmt.Printf("Exponential distribution (λ=%.1f):\n", lambda)
	for i := 0; i < 10; i++ {
		fmt.Printf("%.4f ", ExponentialRandom(lambda))
	}
	fmt.Println()

	// Random permutations
	fmt.Println("\n--- Random Permutations ---")
	perm := RandomPermutation(10)
	fmt.Printf("Random permutation of [0,1,2,3,4,5,6,7,8,9]: %v\n", perm)

	// Random sampling
	fmt.Println("\n--- Random Sampling ---")
	population := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sample := RandomSample(population, 5)
	fmt.Printf("Random sample of 5 from %v: %v\n", population, sample)

	// Random strings
	fmt.Println("\n--- Random Strings ---")
	fmt.Printf("Random string (10 chars): %s\n", RandomString(10))
	fmt.Printf("Random password: %s\n", RandomPassword(12, true, true, true))

	// Monte Carlo simulation
	fmt.Println("\n--- Monte Carlo Simulation ---")
	samples := 100000
	estimatedPi := EstimatePi(samples)
	fmt.Printf("Estimated π using %d samples: %.6f\n", samples, estimatedPi)
	fmt.Printf("Actual π: %.6f\n", math.Pi)
	fmt.Printf("Error: %.6f\n", math.Abs(estimatedPi-math.Pi))

	// Statistical analysis of random numbers
	fmt.Println("\n--- Statistical Analysis ---")
	numSamples := 10000
	sum := 0.0
	sumSquares := 0.0

	for i := 0; i < numSamples; i++ {
		val := rand.Float64()
		sum += val
		sumSquares += val * val
	}

	mean = sum / float64(numSamples)
	variance := sumSquares/float64(numSamples) - mean*mean
	stdDev = math.Sqrt(variance)

	fmt.Printf("Uniform distribution [0,1] statistics:\n")
	fmt.Printf("Mean: %.6f (expected: 0.5)\n", mean)
	fmt.Printf("Variance: %.6f (expected: 1/12 ≈ 0.0833)\n", variance)
	fmt.Printf("Standard deviation: %.6f\n", stdDev)
}
