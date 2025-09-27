package main

import (
    "fmt"
    "strings"
)

func main() {
    sentence := "The quick brown fox jumps over the lazy dog"
    words := strings.Fields(sentence)
    
    fmt.Printf("Sentence: %s\n", sentence)
    fmt.Printf("Number of words: %d\n", len(words))
    fmt.Printf("Words: %v\n", words)
    
    // Count word lengths
    for i, word := range words {
        fmt.Printf("Word %d: '%s' (length: %d)\n", i+1, word, len(word))
    }
}