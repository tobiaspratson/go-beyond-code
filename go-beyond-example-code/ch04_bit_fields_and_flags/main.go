package main

import "fmt"

// Count set flags
func countSetFlags(f uint32) int {
	count := 0
	for f != 0 {
		count += int(f & 1)
		f >>= 1
	}
	return count
}

func main() {
	// Define bit flags
	const (
		FlagA = 1 << iota // 1
		FlagB             // 2
		FlagC             // 4
		FlagD             // 8
		FlagE             // 16
	)

	// Set multiple flags
	flags := FlagA | FlagC | FlagE
	fmt.Printf("Flags: %b (binary), %d (decimal)\n", flags, flags)

	// Check if specific flags are set
	fmt.Printf("FlagA set: %t\n", (flags&FlagA) != 0)
	fmt.Printf("FlagB set: %t\n", (flags&FlagB) != 0)
	fmt.Printf("FlagC set: %t\n", (flags&FlagC) != 0)

	// Toggle flags
	flags ^= FlagB // Toggle FlagB
	fmt.Printf("After toggling FlagB: %b\n", flags)

	// Clear all flags
	flags = 0
	fmt.Printf("All flags cleared: %b\n", flags)

	// Set all flags
	flags = FlagA | FlagB | FlagC | FlagD | FlagE
	fmt.Printf("All flags set: %b\n", flags)

	fmt.Printf("Number of set flags: %d\n", countSetFlags(uint32(flags)))
}
