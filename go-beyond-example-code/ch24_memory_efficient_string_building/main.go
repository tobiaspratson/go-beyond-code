package main

import (
    "fmt"
    "strings"
    "time"
)

// Efficient string builder
type StringBuilder struct {
    builder strings.Builder
    parts   []string
}

func NewStringBuilder() *StringBuilder {
    return &StringBuilder{}
}

func (sb *StringBuilder) WriteString(s string) {
    sb.builder.WriteString(s)
}

func (sb *StringBuilder) WriteByte(c byte) {
    sb.builder.WriteByte(c)
}

func (sb *StringBuilder) WriteRune(r rune) {
    sb.builder.WriteRune(r)
}

func (sb *StringBuilder) String() string {
    return sb.builder.String()
}

func (sb *StringBuilder) Reset() {
    sb.builder.Reset()
}

func (sb *StringBuilder) Len() int {
    return sb.builder.Len()
}

func (sb *StringBuilder) Cap() int {
    return sb.builder.Cap()
}

// Alternative: Pre-allocated string builder
func (sb *StringBuilder) PreAllocate(size int) {
    sb.builder.Grow(size)
}

// Example: Building large strings efficiently
func buildStringInefficient(parts []string) string {
    result := ""
    for _, part := range parts {
        result += part
    }
    return result
}

func buildStringEfficient(parts []string) string {
    var builder strings.Builder
    builder.Grow(len(parts) * 10) // Pre-allocate capacity
    
    for _, part := range parts {
        builder.WriteString(part)
    }
    return builder.String()
}

func main() {
    fmt.Println("=== String Building Performance ===")
    
    // Create test data
    parts := make([]string, 1000)
    for i := 0; i < 1000; i++ {
        parts[i] = fmt.Sprintf("part-%d ", i)
    }
    
    // Test inefficient method
    start := time.Now()
    result1 := buildStringInefficient(parts)
    inefficientTime := time.Since(start)
    fmt.Printf("Inefficient method: %v (length: %d)\n", inefficientTime, len(result1))
    
    // Test efficient method
    start = time.Now()
    result2 := buildStringEfficient(parts)
    efficientTime := time.Since(start)
    fmt.Printf("Efficient method: %v (length: %d)\n", efficientTime, len(result2))
    
    fmt.Printf("Performance improvement: %.2fx faster\n", 
        float64(inefficientTime)/float64(efficientTime))
    
    // Test StringBuilder
    fmt.Println("\n=== StringBuilder Usage ===")
    sb := NewStringBuilder()
    sb.PreAllocate(10000) // Pre-allocate capacity
    
    for i := 0; i < 100; i++ {
        sb.WriteString(fmt.Sprintf("Item %d ", i))
        if i%10 == 0 {
            sb.WriteByte('\n')
        }
    }
    
    fmt.Printf("Builder length: %d, capacity: %d\n", sb.Len(), sb.Cap())
    fmt.Printf("First 100 chars: %s...\n", sb.String()[:100])
}