package main

import (
    "fmt"
    "sync"
    "time"
)

type FileProcessor struct {
    files []string
    results chan string
    wg sync.WaitGroup
}

func NewFileProcessor(files []string) *FileProcessor {
    return &FileProcessor{
        files: files,
        results: make(chan string, len(files)),
    }
}

func (fp *FileProcessor) processFile(filename string) {
    defer fp.wg.Done()
    
    // Simulate file processing
    time.Sleep(200 * time.Millisecond)
    
    result := fmt.Sprintf("Processed: %s", filename)
    fp.results <- result
}

func (fp *FileProcessor) ProcessAll() []string {
    for _, file := range fp.files {
        fp.wg.Add(1)
        go fp.processFile(file)
    }
    
    fp.wg.Wait()
    close(fp.results)
    
    var results []string
    for result := range fp.results {
        results = append(results, result)
    }
    
    return results
}

func main() {
    files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}
    processor := NewFileProcessor(files)
    
    start := time.Now()
    results := processor.ProcessAll()
    elapsed := time.Since(start)
    
    fmt.Printf("Processed %d files in %v\n", len(results), elapsed)
    for _, result := range results {
        fmt.Println(result)
    }
}