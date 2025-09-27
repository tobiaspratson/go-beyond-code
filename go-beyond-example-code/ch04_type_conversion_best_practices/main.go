package main

import (
	"fmt"
	"strconv"
)

// Safe conversion function
func safeInt32ToInt64(val int32) int64 {
	return int64(val) // Always safe
}

func safeInt64ToInt32(val int64) (int32, bool) {
	if val >= -2147483648 && val <= 2147483647 {
		return int32(val), true
	}
	return 0, false // Conversion would lose data
}

// Robust string to number conversion
func robustStringToInt(s string) (int, error) {
	// Try parsing as int first
	if val, err := strconv.Atoi(s); err == nil {
		return val, nil
	}

	// Try parsing as float and converting
	if val, err := strconv.ParseFloat(s, 64); err == nil {
		return int(val), nil
	}

	return 0, fmt.Errorf("cannot convert '%s' to int", s)
}

func main() {

	// Test safe conversions
	smallInt64 := int64(1000)
	largeInt64 := int64(3000000000)

	if val, ok := safeInt64ToInt32(smallInt64); ok {
		fmt.Printf("Safe conversion: %d -> %d\n", smallInt64, val)
	}

	if val, ok := safeInt64ToInt32(largeInt64); ok {
		fmt.Printf("Safe conversion: %d -> %d\n", largeInt64, val)
	} else {
		fmt.Printf("Unsafe conversion: %d would lose data\n", largeInt64)
	}

	// Test robust conversion
	testStrings := []string{"123", "45.67", "abc", "999999999999999999"}

	for _, s := range testStrings {
		if val, err := robustStringToInt(s); err == nil {
			fmt.Printf("'%s' -> %d\n", s, val)
		} else {
			fmt.Printf("'%s' -> Error: %v\n", s, err)
		}
	}
}
