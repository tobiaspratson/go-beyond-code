package main

import (
    "context"
    "fmt"
    "time"
)

type FileProcessor struct {
    files []string
}

func NewFileProcessor(files []string) *FileProcessor {
    return &FileProcessor{files: files}
}

func (fp *FileProcessor) ProcessFiles(ctx context.Context) error {
    for i, file := range fp.files {
        select {
        case <-ctx.Done():
            fmt.Printf("Processing cancelled: %v\n", ctx.Err())
            return ctx.Err()
        default:
            fmt.Printf("Processing file %d: %s\n", i+1, file)
            time.Sleep(200 * time.Millisecond) // Simulate processing
            fmt.Printf("Completed file: %s\n", file)
        }
    }
    return nil
}

func main() {
    files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
    processor := NewFileProcessor(files)
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Process files
    err := processor.ProcessFiles(ctx)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}