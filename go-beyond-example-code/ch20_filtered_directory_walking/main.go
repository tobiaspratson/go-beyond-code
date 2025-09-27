package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type WalkFilter struct {
    IncludeExtensions []string
    ExcludeExtensions []string
    MinSize          int64
    MaxSize          int64
    ExcludeDirs      []string
}

func (wf *WalkFilter) shouldInclude(path string, info os.FileInfo) bool {
    // Skip directories if they're in exclude list
    if info.IsDir() {
        for _, excludeDir := range wf.ExcludeDirs {
            if strings.Contains(path, excludeDir) {
                return false
            }
        }
        return true
    }
    
    // Check file extension
    ext := strings.ToLower(filepath.Ext(path))
    if len(wf.IncludeExtensions) > 0 {
        found := false
        for _, includeExt := range wf.IncludeExtensions {
            if ext == strings.ToLower(includeExt) {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    
    // Check excluded extensions
    for _, excludeExt := range wf.ExcludeExtensions {
        if ext == strings.ToLower(excludeExt) {
            return false
        }
    }
    
    // Check size constraints
    if wf.MinSize > 0 && info.Size() < wf.MinSize {
        return false
    }
    if wf.MaxSize > 0 && info.Size() > wf.MaxSize {
        return false
    }
    
    return true
}

func walkWithFilter(root string, filter *WalkFilter) error {
    return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if filter.shouldInclude(path, info) {
            if info.IsDir() {
                fmt.Printf("Directory: %s\n", path)
            } else {
                fmt.Printf("File: %s (%d bytes)\n", path, info.Size())
            }
        }
        
        return nil
    })
}

func main() {
    filter := &WalkFilter{
        IncludeExtensions: []string{".go", ".txt", ".md"},
        ExcludeExtensions: []string{".tmp", ".log"},
        MinSize:          100, // At least 100 bytes
        MaxSize:          1024 * 1024, // At most 1MB
        ExcludeDirs:      []string{"vendor", "node_modules", ".git"},
    }
    
    fmt.Println("Walking directory with filters:")
    err := walkWithFilter(".", filter)
    if err != nil {
        fmt.Printf("Error walking directory: %v\n", err)
    }
}