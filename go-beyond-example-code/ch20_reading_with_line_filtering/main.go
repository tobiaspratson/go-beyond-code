package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readLinesWithFilter(filename string, filter func(string) bool) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var lines []string
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        if filter(line) {
            lines = append(lines, line)
        }
    }
    
    return lines, scanner.Err()
}

func main() {
    // Filter function for non-empty lines
    nonEmptyFilter := func(line string) bool {
        return strings.TrimSpace(line) != ""
    }
    
    // Filter function for lines containing specific text
    containsFilter := func(searchTerm string) func(string) bool {
        return func(line string) bool {
            return strings.Contains(strings.ToLower(line), strings.ToLower(searchTerm))
        }
    }
    
    // Read non-empty lines
    lines, err := readLinesWithFilter("example.txt", nonEmptyFilter)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("Found %d non-empty lines:\n", len(lines))
    for i, line := range lines {
        fmt.Printf("%d: %s\n", i+1, line)
    }
    
    // Read lines containing "error"
    errorLines, err := readLinesWithFilter("example.txt", containsFilter("error"))
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("\nFound %d lines containing 'error':\n", len(errorLines))
    for i, line := range errorLines {
        fmt.Printf("%d: %s\n", i+1, line)
    }
}