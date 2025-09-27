package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int
    var fromBase, toBase int
    
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)
    fmt.Print("Enter source base (2-36): ")
    fmt.Scanln(&fromBase)
    fmt.Print("Enter target base (2-36): ")
    fmt.Scanln(&toBase)
    
    // Convert to string in source base
    numStr := strconv.FormatInt(int64(num), fromBase)
    fmt.Printf("Number %d in base %d: %s\n", num, fromBase, numStr)
    
    // Parse from source base and convert to target base
    parsed, err := strconv.ParseInt(numStr, fromBase, 64)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    result := strconv.FormatInt(parsed, toBase)
    fmt.Printf("Converted to base %d: %s\n", toBase, result)
}