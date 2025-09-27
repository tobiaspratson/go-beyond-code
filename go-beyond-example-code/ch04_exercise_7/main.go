package main

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

func isValidNumber(s string) bool {
    s = strings.TrimSpace(s)
    if s == "" {
        return false
    }
    
    // Check for valid characters
    for _, r := range s {
        if !unicode.IsDigit(r) && r != '.' && r != '-' && r != '+' {
            return false
        }
    }
    
    // Try to parse as float
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func convertToNumber(s string) (float64, error) {
    if !isValidNumber(s) {
        return 0, fmt.Errorf("invalid number format: %s", s)
    }
    
    return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func main() {
    testInputs := []string{
        "123",
        "123.45",
        "-67.89",
        "+100",
        " 42 ",
        "abc",
        "12.34.56",
        "",
        "   ",
    }
    
    for _, input := range testInputs {
        if isValidNumber(input) {
            if num, err := convertToNumber(input); err == nil {
                fmt.Printf("'%s' -> %.2f (valid)\n", input, num)
            }
        } else {
            fmt.Printf("'%s' -> invalid\n", input)
        }
    }
}