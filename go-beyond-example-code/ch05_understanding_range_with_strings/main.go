package main

import "fmt"

func main() {
    // Different ways to iterate over strings
    text := "Go语言"
    
    fmt.Println("=== Method 1: Range (recommended) ===")
    for i, char := range text {
        fmt.Printf("Position %d: %c\n", i, char)
    }
    
    fmt.Println("\n=== Method 2: Traditional loop (byte-by-byte) ===")
    for i := 0; i < len(text); i++ {
        fmt.Printf("Byte %d: %c (byte value: %d)\n", i, text[i], text[i])
    }
    
    fmt.Println("\n=== Method 3: Convert to rune slice ===")
    runes := []rune(text)
    for i, char := range runes {
        fmt.Printf("Rune %d: %c\n", i, char)
    }
}