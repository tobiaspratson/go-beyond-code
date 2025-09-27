package main

import "fmt"

// Set bit at position pos
func setBit(n int, pos int) int {
    return n | (1 << pos)
}

// Clear bit at position pos
func clearBit(n int, pos int) int {
    return n &^ (1 << pos)
}

// Toggle bit at position pos
func toggleBit(n int, pos int) int {
    return n ^ (1 << pos)
}

// Check if bit at position pos is set
func isBitSet(n int, pos int) bool {
    return (n & (1 << pos)) != 0
}

// Count number of set bits
func countSetBits(n int) int {
    count := 0
    for n != 0 {
        count += n & 1
        n >>= 1
    }
    return count
}

// Check if number is power of 2
func isPowerOfTwo(n int) bool {
    return n > 0 && (n&(n-1)) == 0
}

func main() {
    num := 42  // 101010 in binary
    
    fmt.Printf("Original number: %d (%b)\n", num, num)
    fmt.Printf("Set bit 0: %d (%b)\n", setBit(num, 0), setBit(num, 0))
    fmt.Printf("Clear bit 1: %d (%b)\n", clearBit(num, 1), clearBit(num, 1))
    fmt.Printf("Toggle bit 2: %d (%b)\n", toggleBit(num, 2), toggleBit(num, 2))
    fmt.Printf("Bit 3 is set: %t\n", isBitSet(num, 3))
    fmt.Printf("Number of set bits: %d\n", countSetBits(num))
    fmt.Printf("Is power of 2: %t\n", isPowerOfTwo(num))
}