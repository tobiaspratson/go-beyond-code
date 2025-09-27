package main

import "fmt"

func main() {
    // Demonstrate overflow with int8
    var maxInt8 int8 = 127
    fmt.Printf("Max int8: %d\n", maxInt8)
    
    // This will overflow and wrap around
    var overflowed int8 = maxInt8 + 1
    fmt.Printf("127 + 1 = %d (overflowed!)\n", overflowed)
    
    // Demonstrate underflow with int8
    var minInt8 int8 = -128
    fmt.Printf("Min int8: %d\n", minInt8)
    
    // This will underflow and wrap around
    var underflowed int8 = minInt8 - 1
    fmt.Printf("-128 - 1 = %d (underflowed!)\n", underflowed)
    
    // Unsigned overflow
    var maxUint8 uint8 = 255
    fmt.Printf("Max uint8: %d\n", maxUint8)
    
    var unsignedOverflow uint8 = maxUint8 + 1
    fmt.Printf("255 + 1 = %d (wrapped to 0)\n", unsignedOverflow)
}