package main

import (
    "fmt"
    "os"
)

func main() {
    // Get file information
    fileInfo, err := os.Stat("example.txt")
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }
    
    fmt.Printf("File name: %s\n", fileInfo.Name())
    fmt.Printf("File size: %d bytes\n", fileInfo.Size())
    fmt.Printf("File mode: %s\n", fileInfo.Mode())
    fmt.Printf("File modified: %s\n", fileInfo.ModTime())
    fmt.Printf("Is directory: %t\n", fileInfo.IsDir())
}