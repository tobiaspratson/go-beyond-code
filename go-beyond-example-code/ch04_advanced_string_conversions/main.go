package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Different number bases
    binaryStr := "1010"
    hexStr := "FF"
    octalStr := "777"
    
    // Parse from different bases
    binaryNum, _ := strconv.ParseInt(binaryStr, 2, 64)  // Base 2
    hexNum, _ := strconv.ParseInt(hexStr, 16, 64)       // Base 16
    octalNum, _ := strconv.ParseInt(octalStr, 8, 64)     // Base 8
    
    fmt.Printf("Binary '%s' = %d\n", binaryStr, binaryNum)
    fmt.Printf("Hex '%s' = %d\n", hexStr, hexNum)
    fmt.Printf("Octal '%s' = %d\n", octalStr, octalNum)
    
    // Convert to different bases
    num := 255
    binary := strconv.FormatInt(int64(num), 2)
    hex := strconv.FormatInt(int64(num), 16)
    octal := strconv.FormatInt(int64(num), 8)
    
    fmt.Printf("%d in binary: %s\n", num, binary)
    fmt.Printf("%d in hex: %s\n", num, hex)
    fmt.Printf("%d in octal: %s\n", num, octal)
    
    // Boolean conversions
    boolStr := "true"
    boolVal, err := strconv.ParseBool(boolStr)
    if err != nil {
        fmt.Printf("Error parsing bool: %v\n", err)
    } else {
        fmt.Printf("String '%s' as bool: %t\n", boolStr, boolVal)
    }
    
    // Convert bool to string
    boolToStr := strconv.FormatBool(true)
    fmt.Printf("Bool true as string: '%s'\n", boolToStr)
}