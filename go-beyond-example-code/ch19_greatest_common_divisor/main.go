package main

import "fmt"

// Euclidean algorithm for GCD
func GCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Extended Euclidean algorithm
func ExtendedGCD(a, b int) (int, int, int) {
    if a == 0 {
        return b, 0, 1
    }
    
    gcd, x1, y1 := ExtendedGCD(b%a, a)
    x := y1 - (b/a)*x1
    y := x1
    
    return gcd, x, y
}

// Least Common Multiple
func LCM(a, b int) int {
    return a * b / GCD(a, b)
}

func main() {
    a, b := 48, 18
    
    gcd := GCD(a, b)
    fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
    
    lcm := LCM(a, b)
    fmt.Printf("LCM(%d, %d) = %d\n", a, b, lcm)
    
    gcd, x, y := ExtendedGCD(a, b)
    fmt.Printf("Extended GCD(%d, %d) = %d, x = %d, y = %d\n", a, b, gcd, x, y)
    fmt.Printf("Verification: %d*%d + %d*%d = %d\n", a, x, b, y, a*x+b*y)
}