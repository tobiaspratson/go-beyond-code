package main

import (
    "fmt"
    "math"
    "time"
)

// Check if a number is prime (naive approach)
func IsPrimeNaive(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    for i := 3; i < n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Check if a number is prime (optimized)
func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    // Only check odd divisors up to sqrt(n)
    for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Generate prime numbers up to n (naive approach)
func GeneratePrimesNaive(n int) []int {
    var primes []int
    for i := 2; i <= n; i++ {
        if IsPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

// Sieve of Eratosthenes - O(n log log n)
func SieveOfEratosthenes(n int) []int {
    if n < 2 {
        return []int{}
    }
    
    isPrime := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        isPrime[i] = true
    }
    
    // Mark multiples of each prime as composite
    for i := 2; i*i <= n; i++ {
        if isPrime[i] {
            for j := i * i; j <= n; j += i {
                isPrime[j] = false
            }
        }
    }
    
    var primes []int
    for i := 2; i <= n; i++ {
        if isPrime[i] {
            primes = append(primes, i)
        }
    }
    return primes
}

// Segmented sieve for large ranges
func SegmentedSieve(low, high int) []int {
    if low < 2 {
        low = 2
    }
    if high < low {
        return []int{}
    }
    
    // Find primes up to sqrt(high)
    limit := int(math.Sqrt(float64(high))) + 1
    basePrimes := SieveOfEratosthenes(limit)
    
    // Create array for the range [low, high]
    isPrime := make([]bool, high-low+1)
    for i := range isPrime {
        isPrime[i] = true
    }
    
    // Mark multiples of base primes
    for _, p := range basePrimes {
        start := ((low + p - 1) / p) * p
        if start < p*p {
            start = p * p
        }
        
        for j := start; j <= high; j += p {
            isPrime[j-low] = false
        }
    }
    
    var primes []int
    for i, isP := range isPrime {
        if isP {
            primes = append(primes, low+i)
        }
    }
    return primes
}

// Miller-Rabin primality test (probabilistic)
func MillerRabin(n int, k int) bool {
    if n < 2 {
        return false
    }
    if n == 2 || n == 3 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    // Write n-1 as d * 2^r
    d := n - 1
    r := 0
    for d%2 == 0 {
        d /= 2
        r++
    }
    
    // Witness loop
    for i := 0; i < k; i++ {
        a := 2 + (i % (n-4)) // Random base
        x := modExp(a, d, n)
        
        if x == 1 || x == n-1 {
            continue
        }
        
        for j := 0; j < r-1; j++ {
            x = (x * x) % n
            if x == n-1 {
                break
            }
        }
        
        if x != n-1 {
            return false
        }
    }
    return true
}

// Modular exponentiation
func modExp(base, exp, mod int) int {
    result := 1
    base = base % mod
    
    for exp > 0 {
        if exp%2 == 1 {
            result = (result * base) % mod
        }
        exp = exp >> 1
        base = (base * base) % mod
    }
    return result
}

func main() {
    fmt.Println("=== Prime Number Algorithms ===")
    
    // Test individual numbers
    testNumbers := []int{17, 18, 97, 100, 101, 997, 1000}
    fmt.Println("\n--- Prime Testing ---")
    for _, n := range testNumbers {
        fmt.Printf("Is %d prime? %t\n", n, IsPrime(n))
    }
    
    // Performance comparison
    fmt.Println("\n--- Performance Comparison ---")
    n := 10000
    
    start := time.Now()
    primesNaive := GeneratePrimesNaive(n)
    naiveTime := time.Since(start)
    
    start = time.Now()
    primesSieve := SieveOfEratosthenes(n)
    sieveTime := time.Since(start)
    
    fmt.Printf("Naive approach: %d primes in %v\n", len(primesNaive), naiveTime)
    fmt.Printf("Sieve approach: %d primes in %v\n", len(primesSieve), sieveTime)
    fmt.Printf("Speedup: %.2fx\n", float64(naiveTime)/float64(sieveTime))
    
    // Segmented sieve
    fmt.Println("\n--- Segmented Sieve ---")
    low, high := 1000000, 1000100
    segmentedPrimes := SegmentedSieve(low, high)
    fmt.Printf("Primes between %d and %d: %v\n", low, high, segmentedPrimes)
    
    // Miller-Rabin test
    fmt.Println("\n--- Miller-Rabin Test ---")
    largeNumbers := []int{1009, 1013, 1019, 1021, 1024}
    for _, num := range largeNumbers {
        isPrime := MillerRabin(num, 5)
        fmt.Printf("Miller-Rabin(%d): %t\n", num, isPrime)
    }
}