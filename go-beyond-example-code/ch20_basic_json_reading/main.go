package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    file, err := os.Open("data.json")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    content, err := os.ReadFile("data.json")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    var person Person
    err = json.Unmarshal(content, &person)
    if err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        return
    }
    
    fmt.Printf("Person: %+v\n", person)
}