package main

import "fmt"

func main() {
    // Rune (Unicode code point) - int32 alias
    var letter rune = 'A'
    var unicode rune = '中'  // Chinese character
    var emoji rune = '🚀'    // Rocket emoji
    
    // Byte (ASCII character) - uint8 alias
    var ascii byte = 'B'
    
    // String indexing (returns byte)
    text := "Hello"
    firstChar := text[0]  // 'H' as byte
    
    fmt.Printf("Letter: %c (Unicode: %d)\n", letter, letter)
    fmt.Printf("Unicode: %c (Unicode: %d)\n", unicode, unicode)
    fmt.Printf("Emoji: %c (Unicode: %d)\n", emoji, emoji)
    fmt.Printf("ASCII: %c (ASCII: %d)\n", ascii, ascii)
    fmt.Printf("First char: %c\n", firstChar)
    
    // Converting between runes and bytes
    fmt.Printf("'A' as byte: %d\n", byte('A'))
    fmt.Printf("65 as rune: %c\n", rune(65))
}