package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    text := "the quick brown fox jumps over the lazy dog the fox is quick"
    words := strings.Fields(text)
    
    // Count word frequencies
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    
    // Display frequencies
    fmt.Println("Word frequencies:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
    
    // Find most frequent word
    var mostFrequent string
    var maxCount int
    for word, count := range wordCount {
        if count > maxCount {
            maxCount = count
            mostFrequent = word
        }
    }
    fmt.Printf("\nMost frequent word: '%s' (%d times)\n", mostFrequent, maxCount)
    
    // Sort words by frequency
    fmt.Println("\nWords sorted by frequency:")
    sortedWords := sortByFrequency(wordCount)
    for _, word := range sortedWords {
        fmt.Printf("%s: %d\n", word, wordCount[word])
    }
}

// Sort words by frequency (descending)
func sortByFrequency(wordCount map[string]int) []string {
    var words []string
    for word := range wordCount {
        words = append(words, word)
    }
    
    sort.Slice(words, func(i, j int) bool {
        return wordCount[words[i]] > wordCount[words[j]]
    })
    
    return words
}