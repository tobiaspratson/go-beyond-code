package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func findFiles(root, pattern string) ([]string, error) {
    var matches []string
    
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if !info.IsDir() {
            // Simple pattern matching (you could use regexp for more complex patterns)
            if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
                matches = append(matches, path)
            }
        }
        
        return nil
    })
    
    return matches, err
}

func findFilesByExtension(root string, extensions []string) ([]string, error) {
    var matches []string
    
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        if !info.IsDir() {
            ext := strings.ToLower(filepath.Ext(path))
            for _, targetExt := range extensions {
                if ext == strings.ToLower(targetExt) {
                    matches = append(matches, path)
                    break
                }
            }
        }
        
        return nil
    })
    
    return matches, err
}

func main() {
    root := "."
    
    // Find files by name pattern
    fmt.Println("Finding files with 'test' in name:")
    testFiles, err := findFiles(root, "test")
    if err != nil {
        fmt.Printf("Error finding files: %v\n", err)
        return
    }
    
    for _, file := range testFiles {
        fmt.Printf("  %s\n", file)
    }
    
    // Find files by extension
    fmt.Println("\nFinding Go files:")
    goFiles, err := findFilesByExtension(root, []string{".go"})
    if err != nil {
        fmt.Printf("Error finding Go files: %v\n", err)
        return
    }
    
    for _, file := range goFiles {
        fmt.Printf("  %s\n", file)
    }
}