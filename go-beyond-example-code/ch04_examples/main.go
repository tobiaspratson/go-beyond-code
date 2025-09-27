package main

import "fmt"

func main() {
    // Different integer types with real-world context
    var small int8 = 127        // Perfect for small counters, flags
    var medium int16 = 32000    // Good for medium-sized numbers
    var large int64 = 9223372036854775807  // For very large values
    
    // Unsigned integers (only positive) - great for counts, IDs
    var positive uint32 = 4294967295  // User IDs, file sizes
    
    // Platform-dependent int - most common choice
    var general int = 42
    
    // Zero values - Go initializes to zero
    var uninitialized int
    var uninitializedUint uint
    
    fmt.Printf("Small: %d (type: %T)\n", small, small)
    fmt.Printf("Medium: %d (type: %T)\n", medium, medium)
    fmt.Printf("Large: %d (type: %T)\n", large, large)
    fmt.Printf("Positive: %d (type: %T)\n", positive, positive)
    fmt.Printf("General: %d (type: %T)\n", general, general)
    fmt.Printf("Uninitialized int: %d\n", uninitialized)
    fmt.Printf("Uninitialized uint: %d\n", uninitializedUint)
    
    // Show the ranges
    fmt.Printf("\nint8 range: %d to %d\n", -128, 127)
    fmt.Printf("uint8 range: %d to %d\n", 0, 255)
    fmt.Printf("int32 range: %d to %d\n", -2147483648, 2147483647)
}