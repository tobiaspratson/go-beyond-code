package main

import (
    "fmt"
    "strconv"
)

func convertBase(number string, fromBase, toBase int) (string, error) {
    // Parse from source base
    num, err := strconv.ParseInt(number, fromBase, 64)
    if err != nil {
        return "", fmt.Errorf("invalid number in base %d: %s", fromBase, number)
    }
    
    // Convert to target base
    result := strconv.FormatInt(num, toBase)
    return result, nil
}

func main() {
    // Test conversions
    testCases := []struct {
        number   string
        fromBase int
        toBase   int
    }{
        {"1010", 2, 10},    // Binary to decimal
        {"255", 10, 16},    // Decimal to hex
        {"377", 8, 2},      // Octal to binary
        {"FF", 16, 10},    // Hex to decimal
        {"42", 10, 8},       // Decimal to octal
    }
    
    for _, tc := range testCases {
        result, err := convertBase(tc.number, tc.fromBase, tc.toBase)
        if err != nil {
            fmt.Printf("Error converting %s from base %d: %v\n", 
                tc.number, tc.fromBase, err)
        } else {
            fmt.Printf("%s (base %d) = %s (base %d)\n", 
                tc.number, tc.fromBase, result, tc.toBase)
        }
    }
}