package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Custom split function for comma-separated values
func commaSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    
    if i := strings.Index(string(data), ","); i >= 0 {
        return i + 1, data[0:i], nil
    }
    
    if atEOF {
        return len(data), data, nil
    }
    
    return 0, nil, nil
}

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Split(commaSplit)
    
    fieldCount := 0
    for scanner.Scan() {
        field := strings.TrimSpace(scanner.Text())
        fieldCount++
        fmt.Printf("Field %d: %s\n", fieldCount, field)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}