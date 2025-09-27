package main

import "fmt"

func main() {
    // Count from 0 to 9
    for i := 0; i < 10; i++ {
        fmt.Printf("i = %d\n", i)
    }
    
    // Count backwards
    for i := 10; i > 0; i-- {
        fmt.Printf("Countdown: %d\n", i)
    }
    
    // Count by 2s
    for i := 0; i < 20; i += 2 {
        fmt.Printf("Even: %d\n", i)
    }
}