package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strings"
)

type WordCount struct {
    Word  string
    Count int
}

type TextAnalyzer struct {
    wordCount    map[string]int
    lineCount    int
    charCount    int
    wordCounts   []WordCount
}

func NewTextAnalyzer() *TextAnalyzer {
    return &TextAnalyzer{
        wordCount: make(map[string]int),
    }
}

func (ta *TextAnalyzer) ProcessLine(line string) {
    ta.lineCount++
    ta.charCount += len(line)
    
    // Clean and split words
    words := ta.extractWords(line)
    for _, word := range words {
        if len(word) > 0 {
            ta.wordCount[word]++
        }
    }
}

func (ta *TextAnalyzer) extractWords(line string) []string {
    // Remove punctuation and convert to lowercase
    re := regexp.MustCompile(`[^\w\s]`)
    cleaned := re.ReplaceAllString(line, " ")
    
    words := strings.Fields(strings.ToLower(cleaned))
    return words
}

func (ta *TextAnalyzer) GenerateStats() {
    // Convert map to slice for sorting
    ta.wordCounts = make([]WordCount, 0, len(ta.wordCount))
    for word, count := range ta.wordCount {
        ta.wordCounts = append(ta.wordCounts, WordCount{Word: word, Count: count})
    }
    
    // Sort by count (descending)
    sort.Slice(ta.wordCounts, func(i, j int) bool {
        return ta.wordCounts[i].Count > ta.wordCounts[j].Count
    })
}

func (ta *TextAnalyzer) PrintStats() {
    fmt.Printf("Text Statistics:\n")
    fmt.Printf("  Lines: %d\n", ta.lineCount)
    fmt.Printf("  Characters: %d\n", ta.charCount)
    fmt.Printf("  Unique words: %d\n", len(ta.wordCount))
    
    totalWords := 0
    for _, count := range ta.wordCount {
        totalWords += count
    }
    fmt.Printf("  Total words: %d\n", totalWords)
    
    fmt.Printf("\nTop 10 most frequent words:\n")
    for i, wc := range ta.wordCounts {
        if i >= 10 {
            break
        }
        fmt.Printf("  %d. %s (%d times)\n", i+1, wc.Word, wc.Count)
    }
}

func analyzeTextFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    analyzer := NewTextAnalyzer()
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        analyzer.ProcessLine(line)
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    
    analyzer.GenerateStats()
    analyzer.PrintStats()
    
    return nil
}

func main() {
    err := analyzeTextFile("document.txt")
    if err != nil {
        fmt.Printf("Error analyzing text: %v\n", err)
    }
}