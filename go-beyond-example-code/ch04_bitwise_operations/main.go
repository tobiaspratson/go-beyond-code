package main

import "fmt"

func main() {
    a, b := 5, 3  // Binary: 101, 011
    
    fmt.Printf("a = %d (%b)\n", a, a)
    fmt.Printf("b = %d (%b)\n", b, b)
    fmt.Printf("a & b = %d (%b)  // AND\n", a&b, a&b)
    fmt.Printf("a | b = %d (%b)  // OR\n", a|b, a|b)
    fmt.Printf("a ^ b = %d (%b)  // XOR\n", a^b, a^b)
    fmt.Printf("a << 1 = %d (%b)  // Left shift\n", a<<1, a<<1)
    fmt.Printf("a >> 1 = %d (%b)  // Right shift\n", a>>1, a>>1)
    fmt.Printf("^a = %d (%b)  // NOT\n", ^a, ^a)
}