package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func main() {
    // Working with different languages
    text := "Hello 世界 🌍"
    
    fmt.Printf("Text: %s\n", text)
    fmt.Printf("Byte length: %d\n", len(text))
    fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))
    
    // Iterate over runes (characters)
    fmt.Print("Characters: ")
    for i, r := range text {
        fmt.Printf("%d:'%c' ", i, r)
    }
    fmt.Println()
    
    // Unicode categories
    for _, r := range text {
        if unicode.IsLetter(r) {
            fmt.Printf("'%c' is a letter\n", r)
        } else if unicode.IsDigit(r) {
            fmt.Printf("'%c' is a digit\n", r)
        } else if unicode.IsSpace(r) {
            fmt.Printf("'%c' is whitespace\n", r)
        } else if unicode.IsSymbol(r) {
            fmt.Printf("'%c' is a symbol\n", r)
        }
    }
}