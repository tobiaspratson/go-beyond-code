package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords) // Split by words instead of lines
    
    wordCount := 0
    wordFreq := make(map[string]int)
    
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        // Remove punctuation
        word = strings.Trim(word, ".,!?;:\"'")
        
        if len(word) > 0 {
            wordCount++
            wordFreq[word]++
            fmt.Printf("Word %d: %s\n", wordCount, word)
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
    
    fmt.Printf("Total words: %d\n", wordCount)
    fmt.Printf("Unique words: %d\n", len(wordFreq))
}