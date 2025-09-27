package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type WalkStats struct {
    TotalFiles    int
    TotalDirs     int
    TotalSize     int64
    FileTypes     map[string]int
    LargestFile   string
    LargestSize   int64
}

func walkDirectory(root string) (*WalkStats, error) {
    stats := &WalkStats{
        FileTypes: make(map[string]int),
    }
    
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            // Log error but continue walking
            fmt.Printf("Error accessing %s: %v\n", path, err)
            return nil
        }
        
        if info.IsDir() {
            stats.TotalDirs++
        } else {
            stats.TotalFiles++
            stats.TotalSize += info.Size()
            
            // Track file types
            ext := strings.ToLower(filepath.Ext(path))
            if ext == "" {
                ext = "no extension"
            }
            stats.FileTypes[ext]++
            
            // Track largest file
            if info.Size() > stats.LargestSize {
                stats.LargestFile = path
                stats.LargestSize = info.Size()
            }
        }
        
        return nil
    })
    
    return stats, err
}

func main() {
    root := "." // Start from current directory
    stats, err := walkDirectory(root)
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
        return
    }
    
    fmt.Printf("Directory Statistics for %s:\n", root)
    fmt.Printf("Total files: %d\n", stats.TotalFiles)
    fmt.Printf("Total directories: %d\n", stats.TotalDirs)
    fmt.Printf("Total size: %d bytes\n", stats.TotalSize)
    fmt.Printf("Largest file: %s (%d bytes)\n", stats.LargestFile, stats.LargestSize)
    
    fmt.Println("\nFile types:")
    for ext, count := range stats.FileTypes {
        fmt.Printf("  %s: %d files\n", ext, count)
    }
}