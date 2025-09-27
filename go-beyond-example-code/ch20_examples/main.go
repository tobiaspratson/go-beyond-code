package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "time"
)

type SearchOptions struct {
    CaseSensitive bool
    WholeWord     bool
    Regex         bool
    FilePattern   string
    ExcludeDirs   []string
    MaxDepth      int
    ShowLineNum   bool
    ShowContext   int
}

type SearchResult struct {
    File        string
    LineNumber  int
    Line        string
    Context     []string
}

func searchInFile(filename, searchTerm string, options SearchOptions) []SearchResult {
    file, err := os.Open(filename)
    if err != nil {
        return nil
    }
    defer file.Close()
    
    var results []SearchResult
    scanner := bufio.NewScanner(file)
    lineNumber := 0
    var contextLines []string
    
    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        
        // Store context lines
        contextLines = append(contextLines, line)
        if len(contextLines) > options.ShowContext*2+1 {
            contextLines = contextLines[1:]
        }
        
        var matches bool
        
        if options.Regex {
            re, err := regexp.Compile(searchTerm)
            if err != nil {
                continue
            }
            matches = re.MatchString(line)
        } else {
            searchLine := line
            searchTerm := searchTerm
            
            if !options.CaseSensitive {
                searchLine = strings.ToLower(searchLine)
                searchTerm = strings.ToLower(searchTerm)
            }
            
            if options.WholeWord {
                matches = strings.Contains(" "+searchLine+" ", " "+searchTerm+" ")
            } else {
                matches = strings.Contains(searchLine, searchTerm)
            }
        }
        
        if matches {
            result := SearchResult{
                File:       filename,
                LineNumber: lineNumber,
                Line:       line,
            }
            
            if options.ShowContext > 0 {
                result.Context = make([]string, len(contextLines))
                copy(result.Context, contextLines)
            }
            
            results = append(results, result)
        }
    }
    
    return results
}

func shouldSearchFile(path string, options SearchOptions) bool {
    if options.FilePattern == "" {
        return true
    }
    
    matched, err := filepath.Match(options.FilePattern, filepath.Base(path))
    if err != nil {
        return false
    }
    
    return matched
}

func searchInDirectory(dir, searchTerm string, options SearchOptions) []SearchResult {
    var allResults []SearchResult
    depth := 0
    
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        // Check depth limit
        if options.MaxDepth > 0 {
            relPath, _ := filepath.Rel(dir, path)
            depth = strings.Count(relPath, string(filepath.Separator))
            if depth > options.MaxDepth {
                if info.IsDir() {
                    return filepath.SkipDir
                }
                return nil
            }
        }
        
        // Skip excluded directories
        if info.IsDir() {
            for _, excludeDir := range options.ExcludeDirs {
                if strings.Contains(path, excludeDir) {
                    return filepath.SkipDir
                }
            }
            return nil
        }
        
        // Check file pattern
        if !shouldSearchFile(path, options) {
            return nil
        }
        
        // Search in file
        results := searchInFile(path, searchTerm, options)
        allResults = append(allResults, results...)
        
        return nil
    })
    
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
    }
    
    return allResults
}

func printResults(results []SearchResult, options SearchOptions) {
    if len(results) == 0 {
        fmt.Println("No matches found.")
        return
    }
    
    fmt.Printf("Found %d matches:\n\n", len(results))
    
    for _, result := range results {
        fmt.Printf("File: %s", result.File)
        if options.ShowLineNum {
            fmt.Printf(" (line %d)", result.LineNumber)
        }
        fmt.Println()
        
        if options.ShowContext > 0 && len(result.Context) > 0 {
            fmt.Println("Context:")
            for i, contextLine := range result.Context {
                marker := "  "
                if i == len(result.Context)/2 {
                    marker = "> "
                }
                fmt.Printf("%s%s\n", marker, contextLine)
            }
        } else {
            fmt.Printf("  %s\n", result.Line)
        }
        fmt.Println()
    }
}

func main() {
    var options SearchOptions
    
    flag.BoolVar(&options.CaseSensitive, "case", false, "Case sensitive search")
    flag.BoolVar(&options.WholeWord, "word", false, "Match whole words only")
    flag.BoolVar(&options.Regex, "regex", false, "Use regular expressions")
    flag.StringVar(&options.FilePattern, "file", "", "File pattern to search (e.g., '*.go')")
    flag.StringVar(&options.ExcludeDirs, "exclude", "", "Comma-separated directories to exclude")
    flag.IntVar(&options.MaxDepth, "depth", 0, "Maximum directory depth (0 = unlimited)")
    flag.BoolVar(&options.ShowLineNum, "line", true, "Show line numbers")
    flag.IntVar(&options.ShowContext, "context", 0, "Number of context lines to show")
    
    flag.Parse()
    
    if flag.NArg() < 2 {
        fmt.Println("Usage: go run main.go [options] <directory> <search_term>")
        flag.PrintDefaults()
        return
    }
    
    dir := flag.Arg(0)
    searchTerm := flag.Arg(1)
    
    if options.ExcludeDirs != "" {
        options.ExcludeDirs = strings.Split(options.ExcludeDirs, ",")
    }
    
    start := time.Now()
    results := searchInDirectory(dir, searchTerm, options)
    duration := time.Since(start)
    
    printResults(results, options)
    fmt.Printf("Search completed in %v\n", duration)
}