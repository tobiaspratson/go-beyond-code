package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func calculateSize(path string) (int64, error) {
    info, err := os.Stat(path)
    if err != nil {
        return 0, err
    }
    
    if info.IsDir() {
        var totalSize int64
        err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() {
                totalSize += info.Size()
            }
            return nil
        })
        return totalSize, err
    }
    
    return info.Size(), nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <path>")
        return
    }
    
    path := os.Args[1]
    size, err := calculateSize(path)
    if err != nil {
        fmt.Printf("Error calculating size: %v\n", err)
        return
    }
    
    fmt.Printf("Size of %s: %d bytes\n", path, size)
}