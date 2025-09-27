package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strings"
    "time"
)

type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
    Service   string
}

type LogAnalyzer struct {
    entries     []LogEntry
    levelCounts map[string]int
    serviceCounts map[string]int
    errorMessages []string
}

func NewLogAnalyzer() *LogAnalyzer {
    return &LogAnalyzer{
        levelCounts:   make(map[string]int),
        serviceCounts: make(map[string]int),
    }
}

func (la *LogAnalyzer) ParseLogLine(line string) *LogEntry {
    // Example log format: "2023-01-01 12:00:00 [INFO] service-name: Message"
    re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (\w+): (.+)`)
    matches := re.FindStringSubmatch(line)
    
    if len(matches) != 5 {
        return nil
    }
    
    timestamp, err := time.Parse("2006-01-02 15:04:05", matches[1])
    if err != nil {
        return nil
    }
    
    return &LogEntry{
        Timestamp: timestamp,
        Level:     matches[2],
        Message:   matches[4],
        Service:   matches[3],
    }
}

func (la *LogAnalyzer) ProcessEntry(entry *LogEntry) {
    la.entries = append(la.entries, *entry)
    la.levelCounts[entry.Level]++
    la.serviceCounts[entry.Service]++
    
    if entry.Level == "ERROR" {
        la.errorMessages = append(la.errorMessages, entry.Message)
    }
}

func (la *LogAnalyzer) PrintAnalysis() {
    fmt.Printf("Log Analysis Results:\n")
    fmt.Printf("Total entries: %d\n", len(la.entries))
    
    fmt.Printf("\nLevel distribution:\n")
    for level, count := range la.levelCounts {
        fmt.Printf("  %s: %d\n", level, count)
    }
    
    fmt.Printf("\nService distribution:\n")
    for service, count := range la.serviceCounts {
        fmt.Printf("  %s: %d\n", service, count)
    }
    
    if len(la.errorMessages) > 0 {
        fmt.Printf("\nError messages (%d total):\n", len(la.errorMessages))
        for i, msg := range la.errorMessages {
            if i >= 5 { // Show only first 5
                fmt.Printf("  ... and %d more\n", len(la.errorMessages)-5)
                break
            }
            fmt.Printf("  %s\n", msg)
        }
    }
}

func analyzeLogFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    analyzer := NewLogAnalyzer()
    scanner := bufio.NewScanner(file)
    lineNumber := 0
    
    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        
        entry := analyzer.ParseLogLine(line)
        if entry != nil {
            analyzer.ProcessEntry(entry)
        } else {
            fmt.Printf("Warning: could not parse line %d: %s\n", lineNumber, line)
        }
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    
    analyzer.PrintAnalysis()
    return nil
}

func main() {
    err := analyzeLogFile("application.log")
    if err != nil {
        fmt.Printf("Error analyzing log: %v\n", err)
    }
}