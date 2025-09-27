package main

import "fmt"

func main() {
    // Create variables
    a := 10
    b := 20
    c := 30
    
    // Create pointers
    ptrA := &a
    ptrB := &b
    ptrC := &c
    
    fmt.Println("=== Memory Layout ===")
    fmt.Printf("Variable a: value=%d, address=%p\n", a, &a)
    fmt.Printf("Variable b: value=%d, address=%p\n", b, &b)
    fmt.Printf("Variable c: value=%d, address=%p\n", c, &c)
    fmt.Println()
    fmt.Printf("Pointer ptrA: points to=%p, value=%d\n", ptrA, *ptrA)
    fmt.Printf("Pointer ptrB: points to=%p, value=%d\n", ptrB, *ptrB)
    fmt.Printf("Pointer ptrC: points to=%p, value=%d\n", ptrC, *ptrC)
    
    // Chain pointers
    ptrToPtr := &ptrA
    fmt.Printf("Pointer to pointer: %p -> %p -> %d\n", &ptrA, ptrA, **ptrToPtr)
}