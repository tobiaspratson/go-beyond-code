package main

import "fmt"

func main() {
    // Infinite loop (be careful!)
    count := 0
    for {
        count++
        fmt.Printf("Infinite count: %d\n", count)
        
        // Break out of infinite loop
        if count >= 5 {
            break
        }
    }
}