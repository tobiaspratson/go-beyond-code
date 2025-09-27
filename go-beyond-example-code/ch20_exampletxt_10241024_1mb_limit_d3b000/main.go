package main

import (
    "fmt"
    "os"
    "path/filepath"
    "sync"
)

type FileProcessor struct {
    mu sync.Mutex
    processed int
    errors    []error
}

func (fp *FileProcessor) processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        fp.mu.Lock()
        fp.errors = append(fp.errors, fmt.Errorf("failed to open %s: %w", filename, err))
        fp.mu.Unlock()
        return err
    }
    defer file.Close()
    
    // Process file...
    fmt.Printf("Processing %s\n", filename)
    
    fp.mu.Lock()
    fp.processed++
    fp.mu.Unlock()
    
    return nil
}

func (fp *FileProcessor) processFilesConcurrently(files []string) {
    var wg sync.WaitGroup
    
    for _, file := range files {
        wg.Add(1)
        go func(filename string) {
            defer wg.Done()
            fp.processFile(filename)
        }(file)
    }
    
    wg.Wait()
}

func main() {
    processor := &FileProcessor{}
    
    // Get list of files to process
    files := []string{"file1.txt", "file2.txt", "file3.txt"}
    
    processor.processFilesConcurrently(files)
    
    fmt.Printf("Processed %d files\n", processor.processed)
    if len(processor.errors) > 0 {
        fmt.Printf("Errors: %v\n", processor.errors)
    }
}