package main

import "fmt"

func main() {
    text := "Hello, World!"
    charCount := make(map[rune]int)
    
    for _, char := range text {
        charCount[char]++
    }
    
    fmt.Printf("Character frequencies in '%s':\n", text)
    for char, count := range charCount {
        fmt.Printf("'%c': %d\n", char, count)
    }
}