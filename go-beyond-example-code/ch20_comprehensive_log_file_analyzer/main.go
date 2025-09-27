package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "path/filepath"
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
    Thread    string
    UserID    string
}

type LogStats struct {
    TotalEntries    int
    LevelCounts     map[string]int
    ServiceCounts  map[string]int
    HourlyCounts   map[int]int
    ErrorMessages  []string
    TopErrors      map[string]int
    ResponseTimes  []float64
    UserActivity   map[string]int
}

type LogAnalyzer struct {
    stats     LogStats
    patterns  map[string]*regexp.Regexp
    startTime time.Time
    endTime   time.Time
}

func NewLogAnalyzer() *LogAnalyzer {
    return &LogAnalyzer{
        stats: LogStats{
            LevelCounts:    make(map[string]int),
            ServiceCounts: make(map[string]int),
            HourlyCounts:  make(map[int]int),
            TopErrors:     make(map[string]int),
            UserActivity:  make(map[string]int),
        },
        patterns: make(map[string]*regexp.Regexp),
    }
}

func (la *LogAnalyzer) compilePatterns() {
    // Common log patterns
    patterns := map[string]string{
        "standard": `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (\w+): (.+)`,
        "apache":   `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (.+)`,
        "json":     `{"timestamp":"([^"]+)","level":"(\w+)","message":"([^"]+)","service":"([^"]+)"}`,
        "nginx":    `(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (.+)`,
    }
    
    for name, pattern := range patterns {
        re, err := regexp.Compile(pattern)
        if err != nil {
            fmt.Printf("Warning: failed to compile pattern %s: %v\n", name, err)
            continue
        }
        la.patterns[name] = re
    }
}

func (la *LogAnalyzer) parseLogLine(line string) *LogEntry {
    for patternName, re := range la.patterns {
        matches := re.FindStringSubmatch(line)
        if len(matches) < 4 {
            continue
        }
        
        var timestamp time.Time
        var err error
        
        switch patternName {
        case "standard":
            timestamp, err = time.Parse("2006-01-02 15:04:05", matches[1])
            if err != nil {
                continue
            }
            return &LogEntry{
                Timestamp: timestamp,
                Level:     matches[2],
                Service:   matches[3],
                Message:   matches[4],
            }
        case "apache":
            timestamp, err = time.Parse("2006-01-02 15:04:05", matches[1])
            if err != nil {
                continue
            }
            return &LogEntry{
                Timestamp: timestamp,
                Level:     matches[2],
                Message:   matches[3],
            }
        case "json":
            timestamp, err = time.Parse(time.RFC3339, matches[1])
            if err != nil {
                continue
            }
            return &LogEntry{
                Timestamp: timestamp,
                Level:     matches[2],
                Message:   matches[3],
                Service:   matches[4],
            }
        }
    }
    
    return nil
}

func (la *LogAnalyzer) processEntry(entry *LogEntry) {
    la.stats.TotalEntries++
    la.stats.LevelCounts[entry.Level]++
    la.stats.ServiceCounts[entry.Service]++
    
    hour := entry.Timestamp.Hour()
    la.stats.HourlyCounts[hour]++
    
    if entry.Level == "ERROR" || entry.Level == "FATAL" {
        la.stats.ErrorMessages = append(la.stats.ErrorMessages, entry.Message)
        la.stats.TopErrors[entry.Message]++
    }
    
    // Extract user activity if present
    if entry.UserID != "" {
        la.stats.UserActivity[entry.UserID]++
    }
    
    // Track time range
    if la.startTime.IsZero() || entry.Timestamp.Before(la.startTime) {
        la.startTime = entry.Timestamp
    }
    if la.endTime.IsZero() || entry.Timestamp.After(la.endTime) {
        la.endTime = entry.Timestamp
    }
}

func (la *LogAnalyzer) analyzeFile(filename string) error {
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
        
        entry := la.parseLogLine(line)
        if entry != nil {
            la.processEntry(entry)
        }
    }
    
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    
    return nil
}

func (la *LogAnalyzer) printAnalysis() {
    fmt.Printf("Log Analysis Results:\n")
    fmt.Printf("====================\n\n")
    
    fmt.Printf("Time Range: %s to %s\n", 
        la.startTime.Format("2006-01-02 15:04:05"),
        la.endTime.Format("2006-01-02 15:04:05"))
    fmt.Printf("Total entries: %d\n\n", la.stats.TotalEntries)
    
    // Level distribution
    fmt.Printf("Level Distribution:\n")
    var levels []string
    for level := range la.stats.LevelCounts {
        levels = append(levels, level)
    }
    sort.Strings(levels)
    
    for _, level := range levels {
        count := la.stats.LevelCounts[level]
        percentage := float64(count) / float64(la.stats.TotalEntries) * 100
        fmt.Printf("  %s: %d (%.1f%%)\n", level, count, percentage)
    }
    
    // Service distribution
    if len(la.stats.ServiceCounts) > 0 {
        fmt.Printf("\nService Distribution:\n")
        var services []string
        for service := range la.stats.ServiceCounts {
            services = append(services, service)
        }
        sort.Strings(services)
        
        for _, service := range services {
            count := la.stats.ServiceCounts[service]
            percentage := float64(count) / float64(la.stats.TotalEntries) * 100
            fmt.Printf("  %s: %d (%.1f%%)\n", service, count, percentage)
        }
    }
    
    // Hourly distribution
    fmt.Printf("\nHourly Distribution:\n")
    for hour := 0; hour < 24; hour++ {
        count := la.stats.HourlyCounts[hour]
        if count > 0 {
            percentage := float64(count) / float64(la.stats.TotalEntries) * 100
            fmt.Printf("  %02d:00: %d (%.1f%%)\n", hour, count, percentage)
        }
    }
    
    // Top errors
    if len(la.stats.TopErrors) > 0 {
        fmt.Printf("\nTop Error Messages:\n")
        type errorCount struct {
            message string
            count   int
        }
        
        var errors []errorCount
        for msg, count := range la.stats.TopErrors {
            errors = append(errors, errorCount{msg, count})
        }
        
        sort.Slice(errors, func(i, j int) bool {
            return errors[i].count > errors[j].count
        })
        
        for i, err := range errors {
            if i >= 10 {
                break
            }
            fmt.Printf("  %d. %s (%d times)\n", i+1, err.message, err.count)
        }
    }
    
    // User activity
    if len(la.stats.UserActivity) > 0 {
        fmt.Printf("\nUser Activity:\n")
        type userCount struct {
            user  string
            count int
        }
        
        var users []userCount
        for user, count := range la.stats.UserActivity {
            users = append(users, userCount{user, count})
        }
        
        sort.Slice(users, func(i, j int) bool {
            return users[i].count > users[j].count
        })
        
        for i, user := range users {
            if i >= 10 {
                break
            }
            fmt.Printf("  %s: %d entries\n", user.user, user.count)
        }
    }
}

func main() {
    var (
        pattern = flag.String("pattern", "standard", "Log pattern (standard, apache, json, nginx)")
        verbose = flag.Bool("verbose", false, "Verbose output")
    )
    
    flag.Parse()
    
    if flag.NArg() < 1 {
        fmt.Println("Usage: go run main.go [options] <log_file>")
        flag.PrintDefaults()
        return
    }
    
    filename := flag.Arg(0)
    
    analyzer := NewLogAnalyzer()
    analyzer.compilePatterns()
    
    if *verbose {
        fmt.Printf("Analyzing log file: %s\n", filename)
        fmt.Printf("Using pattern: %s\n\n", *pattern)
    }
    
    start := time.Now()
    err := analyzer.analyzeFile(filename)
    if err != nil {
        fmt.Printf("Error analyzing log file: %v\n", err)
        return
    }
    duration := time.Since(start)
    
    analyzer.printAnalysis()
    
    if *verbose {
        fmt.Printf("\nAnalysis completed in %v\n", duration)
    }
}