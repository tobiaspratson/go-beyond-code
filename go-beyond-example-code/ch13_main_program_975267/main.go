package main

import (
    "fmt"
    "strings"
    "time"
)

func stage1(input <-chan string, output chan<- string) {
    defer close(output)
    
    for text := range input {
        // Simulate processing
        time.Sleep(100 * time.Millisecond)
        output <- strings.ToUpper(text)
    }
}

func stage2(input <-chan string, output chan<- string) {
    defer close(output)
    
    for text := range input {
        // Simulate processing
        time.Sleep(100 * time.Millisecond)
        output <- strings.TrimSpace(text)
    }
}

func stage3(input <-chan string, output chan<- string) {
    defer close(output)
    
    for text := range input {
        // Simulate processing
        time.Sleep(100 * time.Millisecond)
        output <- fmt.Sprintf("Processed: %s", text)
    }
}

func main() {
    // Create pipeline stages
    input := make(chan string, 5)
    stage1Out := make(chan string, 5)
    stage2Out := make(chan string, 5)
    output := make(chan string, 5)
    
    // Start pipeline
    go stage1(input, stage1Out)
    go stage2(stage1Out, stage2Out)
    go stage3(stage2Out, output)
    
    // Send input data
    texts := []string{"hello", "world", "golang", "concurrency"}
    for _, text := range texts {
        input <- text
    }
    close(input)
    
    // Collect results
    for result := range output {
        fmt.Println(result)
    }
}