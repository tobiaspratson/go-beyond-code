package main

import "fmt"

func main() {
    // Array and slice pointers
    arr := [5]int{1, 2, 3, 4, 5}
    slice := []int{10, 20, 30, 40, 50}
    
    // Point to array elements
    ptr1 := &arr[0]
    ptr2 := &arr[2]
    
    fmt.Printf("arr[0] = %d, address = %p\n", *ptr1, ptr1)
    fmt.Printf("arr[2] = %d, address = %p\n", *ptr2, ptr2)
    
    // Point to slice elements
    slicePtr := &slice[1]
    fmt.Printf("slice[1] = %d, address = %p\n", *slicePtr, slicePtr)
    
    // Modify through pointers
    *ptr1 = 100
    *slicePtr = 200
    
    fmt.Printf("After modification:\n")
    fmt.Printf("arr: %v\n", arr)
    fmt.Printf("slice: %v\n", slice)
}