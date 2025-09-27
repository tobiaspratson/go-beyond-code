package main

import "fmt"

func main() {
    // Method 1: Empty slice (nil slice)
    var numbers []int
    fmt.Printf("Empty slice: %v, length: %d, capacity: %d\n", 
        numbers, len(numbers), cap(numbers))
    fmt.Printf("Is nil: %t\n", numbers == nil)
    
    // Method 2: Slice with initial values
    colors := []string{"red", "green", "blue"}
    fmt.Printf("Colors: %v, length: %d, capacity: %d\n", 
        colors, len(colors), cap(colors))
    
    // Method 3: Slice from array
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4]  // [2 3 4]
    fmt.Printf("Slice from array: %v, length: %d, capacity: %d\n", 
        slice, len(slice), cap(slice))
    
    // Method 4: Make a slice with specific length and capacity
    dynamicSlice := make([]int, 3, 5)  // length 3, capacity 5
    fmt.Printf("Made slice: %v, length: %d, capacity: %d\n", 
        dynamicSlice, len(dynamicSlice), cap(dynamicSlice))
    
    // Method 5: Slice literal with capacity
    literalSlice := []int{1, 2, 3, 4, 5}
    fmt.Printf("Literal slice: %v, length: %d, capacity: %d\n", 
        literalSlice, len(literalSlice), cap(literalSlice))
}