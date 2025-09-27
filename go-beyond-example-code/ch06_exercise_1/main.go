package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
    
    // Calculate sum, average, min, max
    sum := 0
    min := numbers[0]
    max := numbers[0]
    
    for _, num := range numbers {
        sum += num
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    
    average := float64(sum) / float64(len(numbers))
    
    fmt.Printf("Numbers: %v\n", numbers)
    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Average: %.2f\n", average)
    fmt.Printf("Min: %d\n", min)
    fmt.Printf("Max: %d\n", max)
}