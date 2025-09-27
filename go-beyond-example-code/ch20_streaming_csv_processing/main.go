package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func processCSVStream(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1 // Allow variable number of fields
    
    lineNumber := 0
    for {
        record, err := reader.Read()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            return fmt.Errorf("error reading CSV at line %d: %w", lineNumber+1, err)
        }
        
        lineNumber++
        fmt.Printf("Line %d: %v\n", lineNumber, record)
        
        // Process each record as it's read
        // This is memory-efficient for large files
    }
    
    return nil
}

func main() {
    err := processCSVStream("large_data.csv")
    if err != nil {
        fmt.Printf("Error processing CSV: %v\n", err)
    }
}