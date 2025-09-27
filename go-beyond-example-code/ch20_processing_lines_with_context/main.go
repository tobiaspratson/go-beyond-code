package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type LineProcessor struct {
    lineNumber int
    totalLines int
    wordCount  int
    charCount  int
}

func (lp *LineProcessor) ProcessLine(line string) {
    lp.lineNumber++
    lp.charCount += len(line)
    
    words := strings.Fields(line)
    lp.wordCount += len(words)
    
    // Process specific line types
    if strings.HasPrefix(line, "#") {
        fmt.Printf("Comment on line %d: %s\n", lp.lineNumber, line)
    } else if strings.Contains(line, "TODO") {
        fmt.Printf("TODO found on line %d: %s\n", lp.lineNumber, line)
    }
}

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    processor := &LineProcessor{}
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        processor.ProcessLine(line)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("File statistics:\n")
    fmt.Printf("  Total lines: %d\n", processor.lineNumber)
    fmt.Printf("  Total words: %d\n", processor.wordCount)
    fmt.Printf("  Total characters: %d\n", processor.charCount)
}