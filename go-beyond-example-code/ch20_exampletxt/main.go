package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // Open file for reading
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close() // Always close the file
    
    // Read file in chunks
    buffer := make([]byte, 1024) // 1KB buffer
    var content []byte
    
    for {
        n, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break // End of file reached
            }
            fmt.Printf("Error reading file: %v\n", err)
            return
        }
        
        // Append read bytes to content
        content = append(content, buffer[:n]...)
    }
    
    fmt.Printf("Read %d bytes: %s\n", len(content), string(content))
}