package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 255
    
    // Convert to different bases
    binary := strconv.FormatInt(int64(num), 2)
    octal := strconv.FormatInt(int64(num), 8)
    hex := strconv.FormatInt(int64(num), 16)
    
    fmt.Printf("Decimal %d in different bases:\n", num)
    fmt.Printf("Binary: %s\n", binary)
    fmt.Printf("Octal: %s\n", octal)
    fmt.Printf("Hexadecimal: %s\n", hex)
    
    // Parse from different bases
    binaryStr := "11111111"
    octalStr := "377"
    hexStr := "FF"
    
    binaryNum, _ := strconv.ParseInt(binaryStr, 2, 64)
    octalNum, _ := strconv.ParseInt(octalStr, 8, 64)
    hexNum, _ := strconv.ParseInt(hexStr, 16, 64)
    
    fmt.Printf("\nParsing from different bases:\n")
    fmt.Printf("Binary '%s' = %d\n", binaryStr, binaryNum)
    fmt.Printf("Octal '%s' = %d\n", octalStr, octalNum)
    fmt.Printf("Hex '%s' = %d\n", hexStr, hexNum)
}