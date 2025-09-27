package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String to number with error handling
    validStr := "123"
    invalidStr := "abc"
    
    // Valid conversion
    num, err := strconv.Atoi(validStr)
    if err != nil {
        fmt.Printf("Error converting '%s': %v\n", validStr, err)
    } else {
        fmt.Printf("String '%s' converted to number: %d\n", validStr, num)
    }
    
    // Invalid conversion
    num, err = strconv.Atoi(invalidStr)
    if err != nil {
        fmt.Printf("Error converting '%s': %v\n", invalidStr, err)
    } else {
        fmt.Printf("String '%s' converted to number: %d\n", invalidStr, num)
    }
    
    // Number to string
    number := 456
    stringNum := strconv.Itoa(number)
    fmt.Printf("Number %d converted to string: '%s'\n", number, stringNum)
    
    // Float to string with formatting
    floatNum := 3.14159
    floatStr := strconv.FormatFloat(floatNum, 'f', 2, 64)  // 2 decimal places
    fmt.Printf("Float %f converted to string: '%s'\n", floatNum, floatStr)
    
    // Parse float from string
    floatStr2 := "3.14"
    parsedFloat, err := strconv.ParseFloat(floatStr2, 64)
    if err != nil {
        fmt.Printf("Error parsing float: %v\n", err)
    } else {
        fmt.Printf("Parsed float: %f\n", parsedFloat)
    }
}