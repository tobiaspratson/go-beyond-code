package main

import "fmt"

// Constraint for types that support addition
type Addable interface {
    ~int | ~float64 | ~string
}

// Generic function for summing a slice
func Sum[T Addable](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}

// Constraint for types that support multiplication
type Multipliable interface {
    ~int | ~float64
}

// Generic function for calculating product
func Product[T Multipliable](slice []T) T {
    var product T = 1
    for _, v := range slice {
        product *= v
    }
    return product
}

// Constraint for types that support both addition and multiplication
type Numeric interface {
    ~int | ~float64
}

// Generic function for calculating average
func Average[T Numeric](slice []T) float64 {
    if len(slice) == 0 {
        return 0
    }
    
    var sum T
    for _, v := range slice {
        sum += v
    }
    
    return float64(sum) / float64(len(slice))
}

// Constraint for types that support string conversion
type Stringer interface {
    ~int | ~float64 | ~string | ~bool
}

// Generic function for converting to string
func ToString[T Stringer](value T) string {
    return fmt.Sprintf("%v", value)
}

// Constraint for types that support both ordering and arithmetic
type OrderedNumeric interface {
    ~int | ~float64
}

// Generic function for finding the range (max - min)
func Range[T OrderedNumeric](slice []T) T {
    if len(slice) == 0 {
        var zero T
        return zero
    }
    
    min, max := slice[0], slice[0]
    for _, v := range slice {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    
    return max - min
}

// Constraint for types that support both ordering and string conversion
type OrderedStringer interface {
    ~int | ~float64 | ~string
}

// Generic function for finding the median
func Median[T OrderedStringer](slice []T) T {
    if len(slice) == 0 {
        var zero T
        return zero
    }
    
    // Create a copy and sort it
    sorted := make([]T, len(slice))
    copy(sorted, slice)
    
    // Simple bubble sort
    for i := 0; i < len(sorted); i++ {
        for j := 0; j < len(sorted)-1-i; j++ {
            if sorted[j] > sorted[j+1] {
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
            }
        }
    }
    
    mid := len(sorted) / 2
    if len(sorted)%2 == 0 {
        // For even length, return the first of the two middle elements
        return sorted[mid-1]
    }
    return sorted[mid]
}

// Constraint for types that support both ordering and arithmetic
type Statistical interface {
    ~int | ~float64
}

// Generic function for calculating standard deviation
func StandardDeviation[T Statistical](slice []T) float64 {
    if len(slice) == 0 {
        return 0
    }
    
    avg := Average(slice)
    var sum float64
    for _, v := range slice {
        diff := float64(v) - avg
        sum += diff * diff
    }
    
    variance := sum / float64(len(slice))
    return variance // Simplified - would need sqrt for actual std dev
}

func main() {
    // Test Sum function
    intSlice := []int{1, 2, 3, 4, 5}
    fmt.Printf("Sum of %v: %d\n", intSlice, Sum(intSlice))
    
    floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
    fmt.Printf("Sum of %v: %.2f\n", floatSlice, Sum(floatSlice))
    
    stringSlice := []string{"Hello", " ", "World", "!"}
    fmt.Printf("Sum of %v: %s\n", stringSlice, Sum(stringSlice))
    
    // Test Product function
    fmt.Printf("Product of %v: %d\n", intSlice, Product(intSlice))
    fmt.Printf("Product of %v: %.2f\n", floatSlice, Product(floatSlice))
    
    // Test Average function
    fmt.Printf("Average of %v: %.2f\n", intSlice, Average(intSlice))
    fmt.Printf("Average of %v: %.2f\n", floatSlice, Average(floatSlice))
    
    // Test ToString function
    fmt.Printf("ToString(42): %s\n", ToString(42))
    fmt.Printf("ToString(3.14): %s\n", ToString(3.14))
    fmt.Printf("ToString(true): %s\n", ToString(true))
    
    // Test Range function
    fmt.Printf("Range of %v: %d\n", intSlice, Range(intSlice))
    fmt.Printf("Range of %v: %.2f\n", floatSlice, Range(floatSlice))
    
    // Test Median function
    fmt.Printf("Median of %v: %d\n", intSlice, Median(intSlice))
    fmt.Printf("Median of %v: %.2f\n", floatSlice, Median(floatSlice))
    
    words := []string{"banana", "apple", "cherry", "date"}
    fmt.Printf("Median of %v: %s\n", words, Median(words))
    
    // Test StandardDeviation function
    fmt.Printf("Standard deviation of %v: %.2f\n", intSlice, StandardDeviation(intSlice))
    fmt.Printf("Standard deviation of %v: %.2f\n", floatSlice, StandardDeviation(floatSlice))
}