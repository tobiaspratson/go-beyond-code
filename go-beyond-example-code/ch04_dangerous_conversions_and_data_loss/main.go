package main

import "fmt"

func main() {
    // Demonstrate data loss in conversions
    var large int64 = 3000000000  // 3 billion
    var small int32 = int32(large)  // This will overflow!
    
    fmt.Printf("Original int64: %d\n", large)
    fmt.Printf("Converted to int32: %d (data lost!)\n", small)
    
    // Negative to unsigned conversion
    var negative int32 = -100
    var unsigned uint32 = uint32(negative)  // This becomes a large positive number!
    
    fmt.Printf("Negative int32: %d\n", negative)
    fmt.Printf("As uint32: %d (wrapped around!)\n", unsigned)
    
    // Safe conversion with bounds checking
    if large <= 2147483647 && large >= -2147483648 {
        safeInt32 := int32(large)
        fmt.Printf("Safe conversion: %d\n", safeInt32)
    } else {
        fmt.Println("Conversion would lose data!")
    }
}