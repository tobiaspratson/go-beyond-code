package main

import "fmt"

// Count number of set bits
func countBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	// Bit manipulation examples
	num := 0b1010 // 10 in binary

	fmt.Printf("Original: %d (%b)\n", num, num)

	// Set bit at position 0 (make it 1)
	setBit := num | (1 << 0) // OR with 1
	fmt.Printf("Set bit 0: %d (%b)\n", setBit, setBit)

	// Clear bit at position 1 (make it 0)
	clearBit := num &^ (1 << 1) // AND NOT with 1
	fmt.Printf("Clear bit 1: %d (%b)\n", clearBit, clearBit)

	// Toggle bit at position 2
	toggleBit := num ^ (1 << 2) // XOR with 1
	fmt.Printf("Toggle bit 2: %d (%b)\n", toggleBit, toggleBit)

	// Check if bit at position 3 is set
	isSet := (num & (1 << 3)) != 0
	fmt.Printf("Bit 3 is set: %t\n", isSet)

	fmt.Printf("Number of set bits in %d: %d\n", num, countBits(num))
}
