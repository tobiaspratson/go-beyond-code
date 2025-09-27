package main

import "fmt"

func main() {
    // Understanding string length vs character count
    english := "Hello"           // 5 characters, 5 bytes
    chinese := "你好"             // 2 characters, 6 bytes (3 bytes per character)
    emoji := "Hello 👋"          // 7 characters, 10 bytes (emoji is 4 bytes)
    
    fmt.Printf("English: '%s' - %d bytes, %d runes\n", 
        english, len(english), len([]rune(english)))
    fmt.Printf("Chinese: '%s' - %d bytes, %d runes\n", 
        chinese, len(chinese), len([]rune(chinese)))
    fmt.Printf("Emoji: '%s' - %d bytes, %d runes\n", 
        emoji, len(emoji), len([]rune(emoji)))
    
    // String indexing returns bytes, not characters!
    fmt.Printf("First byte of '你好': %d (%c)\n", chinese[0], chinese[0])
    fmt.Printf("First rune of '你好': %c\n", []rune(chinese)[0])
}