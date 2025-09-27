package main

import (
    "fmt"
    "os"
)

func processMultipleFiles(inputFile, outputFile string) error {
    // Open input file
    inFile, err := os.Open(inputFile)
    if err != nil {
        return fmt.Errorf("failed to open input file: %w", err)
    }
    defer func() {
        if closeErr := inFile.Close(); closeErr != nil {
            fmt.Printf("Warning: failed to close input file: %v\n", closeErr)
        }
    }()
    
    // Create output file
    outFile, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer func() {
        if closeErr := outFile.Close(); closeErr != nil {
            fmt.Printf("Warning: failed to close output file: %v\n", closeErr)
        }
    }()
    
    // Process files...
    fmt.Printf("Processing %s -> %s\n", inputFile, outputFile)
    
    return nil
}

func main() {
    err := processMultipleFiles("input.txt", "output.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}