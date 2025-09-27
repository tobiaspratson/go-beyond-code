package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV: %v\n", err)
        return
    }
    
    fmt.Println("CSV Data:")
    for i, record := range records {
        fmt.Printf("Row %d: %v\n", i+1, record)
    }
}