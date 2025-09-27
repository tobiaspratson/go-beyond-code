package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    // Read entire file into memory
    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }
    
    // Convert byte slice to string
    text := string(content)
    fmt.Printf("File content (%d bytes):\n%s\n", len(content), text)
    
    // Note: This loads the entire file into memory
    // Not suitable for very large files
}