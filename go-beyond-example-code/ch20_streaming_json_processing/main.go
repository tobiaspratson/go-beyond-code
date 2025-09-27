package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

type LogEntry struct {
    Timestamp string `json:"timestamp"`
    Level     string `json:"level"`
    Message   string `json:"message"`
    Service   string `json:"service"`
}

func processJSONStream(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    lineNumber := 0
    
    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        
        var entry LogEntry
        err := json.Unmarshal([]byte(line), &entry)
        if err != nil {
            fmt.Printf("Warning: failed to parse line %d: %v\n", lineNumber, err)
            continue
        }
        
        fmt.Printf("Entry %d: [%s] %s - %s (%s)\n", 
            lineNumber, entry.Level, entry.Timestamp, 
            entry.Message, entry.Service)
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    
    return nil
}

func main() {
    err := processJSONStream("logs.jsonl") // JSON Lines format
    if err != nil {
        fmt.Printf("Error processing JSON: %v\n", err)
    }
}