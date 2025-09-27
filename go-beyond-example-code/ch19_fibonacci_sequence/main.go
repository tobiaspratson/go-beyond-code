package main

import (
    "fmt"
    "math"
    "time"
)

// Recursive Fibonacci (exponential time complexity)
func FibonacciRecursive(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// Iterative Fibonacci (linear time complexity)
func FibonacciIterative(n int) int {
    if n <= 1 {
        return n
    }
    
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

// Fibonacci with memoization (linear time, linear space)
func FibonacciMemo(n int) int {
    memo := make(map[int]int)
    return fibonacciMemoHelper(n, memo)
}

func fibonacciMemoHelper(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, exists := memo[n]; exists {
        return val
    }
    
    memo[n] = fibonacciMemoHelper(n-1, memo) + fibonacciMemoHelper(n-2, memo)
    return memo[n]
}

// Fibonacci using matrix exponentiation (logarithmic time)
func FibonacciMatrix(n int) int {
    if n <= 1 {
        return n
    }
    
    // Matrix [[1,1],[1,0]]^n = [[F(n+1),F(n)],[F(n),F(n-1)]]
    result := matrixPower([][]int{{1, 1}, {1, 0}}, n-1)
    return result[0][0]
}

// Matrix multiplication
func matrixMultiply(a, b [][]int) [][]int {
    n := len(a)
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}

// Matrix exponentiation using binary exponentiation
func matrixPower(matrix [][]int, exp int) [][]int {
    n := len(matrix)
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
        result[i][i] = 1 // Identity matrix
    }
    
    base := matrix
    for exp > 0 {
        if exp%2 == 1 {
            result = matrixMultiply(result, base)
        }
        base = matrixMultiply(base, base)
        exp /= 2
    }
    return result
}

// Fibonacci using Binet's formula (constant time, but limited precision)
func FibonacciBinet(n int) int {
    if n <= 1 {
        return n
    }
    
    phi := (1 + math.Sqrt(5)) / 2
    psi := (1 - math.Sqrt(5)) / 2
    
    result := (math.Pow(phi, float64(n)) - math.Pow(psi, float64(n))) / math.Sqrt(5)
    return int(math.Round(result))
}

// Generate Fibonacci sequence up to n
func GenerateFibonacciSequence(n int) []int {
    if n < 0 {
        return []int{}
    }
    
    sequence := make([]int, n+1)
    for i := 0; i <= n; i++ {
        sequence[i] = FibonacciIterative(i)
    }
    return sequence
}

// Check if a number is a Fibonacci number
func IsFibonacci(num int) bool {
    if num < 0 {
        return false
    }
    
    // A number is Fibonacci if and only if one of (5*n^2 + 4) or (5*n^2 - 4) is a perfect square
    n := float64(num)
    val1 := 5*n*n + 4
    val2 := 5*n*n - 4
    
    sqrt1 := int(math.Sqrt(val1))
    sqrt2 := int(math.Sqrt(val2))
    
    return sqrt1*sqrt1 == int(val1) || sqrt2*sqrt2 == int(val2)
}

// Find the index of a Fibonacci number
func FibonacciIndex(fib int) int {
    if fib < 0 {
        return -1
    }
    if fib <= 1 {
        return fib
    }
    
    a, b := 0, 1
    index := 1
    
    for b < fib {
        a, b = b, a+b
        index++
    }
    
    if b == fib {
        return index
    }
    return -1
}

func main() {
    fmt.Println("=== Fibonacci Sequence Algorithms ===")
    
    // Generate sequence
    n := 15
    fmt.Printf("\n--- Fibonacci Sequence (first %d numbers) ---\n", n+1)
    for i := 0; i <= n; i++ {
        fib := FibonacciIterative(i)
        fmt.Printf("F(%2d) = %3d\n", i, fib)
    }
    
    // Performance comparison
    fmt.Println("\n--- Performance Comparison ---")
    testN := 35
    
    start := time.Now()
    recursive := FibonacciRecursive(testN)
    recursiveTime := time.Since(start)
    
    start = time.Now()
    iterative := FibonacciIterative(testN)
    iterativeTime := time.Since(start)
    
    start = time.Now()
    memoized := FibonacciMemo(testN)
    memoTime := time.Since(start)
    
    start = time.Now()
    matrix := FibonacciMatrix(testN)
    matrixTime := time.Since(start)
    
    start = time.Now()
    binet := FibonacciBinet(testN)
    binetTime := time.Since(start)
    
    fmt.Printf("F(%d) = %d\n", testN, iterative)
    fmt.Printf("Recursive: %v\n", recursiveTime)
    fmt.Printf("Iterative: %v\n", iterativeTime)
    fmt.Printf("Memoized: %v\n", memoTime)
    fmt.Printf("Matrix: %v\n", matrixTime)
    fmt.Printf("Binet: %v\n", binetTime)
    
    // Verify all methods give same result
    fmt.Printf("\nAll methods agree: %t\n", 
        recursive == iterative && iterative == memoized && 
        memoized == matrix && matrix == binet)
    
    // Fibonacci number properties
    fmt.Println("\n--- Fibonacci Number Properties ---")
    testNumbers := []int{0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 100, 144}
    
    for _, num := range testNumbers {
        isFib := IsFibonacci(num)
        if isFib {
            index := FibonacciIndex(num)
            fmt.Printf("%d is Fibonacci number F(%d)\n", num, index)
        } else {
            fmt.Printf("%d is not a Fibonacci number\n", num)
        }
    }
    
    // Golden ratio approximation
    fmt.Println("\n--- Golden Ratio Approximation ---")
    for i := 1; i <= 20; i++ {
        if i > 1 {
            ratio := float64(FibonacciIterative(i)) / float64(FibonacciIterative(i-1))
            fmt.Printf("F(%d)/F(%d) = %.10f\n", i, i-1, ratio)
        }
    }
}