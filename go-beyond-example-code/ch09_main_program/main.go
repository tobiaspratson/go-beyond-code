package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    text := "The quick brown fox jumps over the lazy dog. The fox is quick!"
    
    // Case-insensitive word frequency
    wordCount := make(map[string]int)
    words := strings.FieldsFunc(text, func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
    })
    
    for _, word := range words {
        word = strings.ToLower(word)
        wordCount[word]++
    }
    
    // Display results
    fmt.Println("Case-insensitive word frequencies:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
    
    // Find words that appear more than once
    fmt.Println("\nWords appearing more than once:")
    for word, count := range wordCount {
        if count > 1 {
            fmt.Printf("%s: %d times\n", word, count)
        }
    }
}