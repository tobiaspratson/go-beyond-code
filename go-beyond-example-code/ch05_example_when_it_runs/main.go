package main

import "fmt"

func main() {
    fmt.Println("Understanding loop execution:")
    
    // Step 1: Initialization (runs once)
    i := 1
    fmt.Printf("1. Initialization: i = %d\n", i)
    
    // Step 2: Check condition
    for i <= 3 {
        fmt.Printf("2. Condition check: i (%d) <= 3? %t\n", i, i <= 3)
        
        // Step 3: Loop body
        fmt.Printf("3. Loop body: processing i = %d\n", i)
        
        // Step 4: Post statement (runs after body)
        i++
        fmt.Printf("4. Post statement: i is now %d\n", i)
        fmt.Println("---")
    }
    
    fmt.Printf("5. Loop finished: i = %d\n", i)
}