package main

import "fmt"

func swap(a, b *int) {
    *a, *b = *b, *a
}

func double(x *int) {
    *x *= 2
}

func main() {
    x, y := 10, 20
    fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
    
    swap(&x, &y)
    fmt.Printf("After swap: x = %d, y = %d\n", x, y)
    
    double(&x)
    fmt.Printf("After doubling x: x = %d\n", x)
}