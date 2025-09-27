package main

import "fmt"

func main() {
    // Decimal (base 10)
    decimal := 42
    
    // Binary (base 2) - prefix 0b
    binary := 0b101010  // 42 in binary
    
    // Octal (base 8) - prefix 0
    octal := 052  // 42 in octal
    
    // Hexadecimal (base 16) - prefix 0x
    hex := 0x2A  // 42 in hexadecimal
    
    fmt.Printf("Decimal: %d\n", decimal)
    fmt.Printf("Binary: %b (decimal: %d)\n", binary, binary)
    fmt.Printf("Octal: %o (decimal: %d)\n", octal, octal)
    fmt.Printf("Hex: %x (decimal: %d)\n", hex, hex)
    fmt.Printf("Hex uppercase: %X\n", hex)
}