package main

import "fmt"

// Pack RGB into a single 32-bit integer
func packRGB(r, g, b uint8) uint32 {
	return uint32(r)<<16 | uint32(g)<<8 | uint32(b)
}

// Unpack RGB from a 32-bit integer
func unpackRGB(color uint32) (uint8, uint8, uint8) {
	r := uint8(color >> 16)
	g := uint8(color >> 8)
	b := uint8(color)
	return r, g, b
}

func main() {

	// Example colors
	red := packRGB(255, 0, 0)
	green := packRGB(0, 255, 0)
	blue := packRGB(0, 0, 255)

	fmt.Printf("Red: 0x%06X\n", red)
	fmt.Printf("Green: 0x%06X\n", green)
	fmt.Printf("Blue: 0x%06X\n", blue)

	// Unpack colors
	r, g, b := unpackRGB(red)
	fmt.Printf("Red unpacked: R=%d, G=%d, B=%d\n", r, g, b)

	// Mix colors (simple average)
	mixed := (red + green) / 2
	mr, mg, mb := unpackRGB(mixed)
	fmt.Printf("Mixed (red+green): R=%d, G=%d, B=%d\n", mr, mg, mb)
}
