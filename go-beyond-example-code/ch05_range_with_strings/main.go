package main

import "fmt"

func main() {
    text := "Hello"
    
    // Range over string (gives index and character)
    fmt.Println("=== Index and Character ===")
    for i, char := range text {
        fmt.Printf("Index %d: %c (Unicode: %d)\n", i, char, char)
    }
    
    // Range over string (index only)
    fmt.Println("\n=== Index Only ===")
    for i := range text {
        fmt.Printf("Index: %d\n", i)
    }
    
    // Range over string (character only)
    fmt.Println("\n=== Character Only ===")
    for _, char := range text {
        fmt.Printf("Character: %c\n", char)
    }
    
    // Working with Unicode characters
    fmt.Println("\n=== Unicode Handling ===")
    unicodeText := "Hello 世界"
    for i, char := range unicodeText {
        fmt.Printf("Index %d: %c (Unicode: %d, bytes: %d)\n", 
            i, char, char, len(string(char)))
    }
}