package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

// Generate random number from normal distribution
func NormalRandom(mean, stdDev float64) float64 {
    // Box-Muller transform
    u1 := rand.Float64()
    u2 := rand.Float64()
    
    z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
    return z0*stdDev + mean
}

// Generate random number from exponential distribution
func ExponentialRandom(lambda float64) float64 {
    return -math.Log(1-rand.Float64()) / lambda
}

// Generate random permutation
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

func main() {
    rand.Seed(time.Now().UnixNano())
    
    // Normal distribution
    fmt.Println("Normal distribution (mean=0, std=1):")
    for i := 0; i < 5; i++ {
        fmt.Printf("%.4f ", NormalRandom(0, 1))
    }
    fmt.Println()
    
    // Exponential distribution
    fmt.Println("Exponential distribution (lambda=1):")
    for i := 0; i < 5; i++ {
        fmt.Printf("%.4f ", ExponentialRandom(1))
    }
    fmt.Println()
    
    // Random permutation
    fmt.Println("Random permutation of [0,1,2,3,4]:")
    perm := RandomPermutation(5)
    fmt.Printf("%v\n", perm)
}