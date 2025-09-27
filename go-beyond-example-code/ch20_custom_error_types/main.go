package main

import (
    "fmt"
    "os"
    "time"
)

type FileError struct {
    Path    string
    Op      string
    Err     error
    Time    time.Time
}

func (fe *FileError) Error() string {
    return fmt.Sprintf("file operation failed: %s %s: %v (at %v)", 
        fe.Op, fe.Path, fe.Err, fe.Time)
}

func (fe *FileError) Unwrap() error {
    return fe.Err
}

func openFileWithCustomError(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, &FileError{
            Path: filename,
            Op:   "open",
            Err:  err,
            Time: time.Now(),
        }
    }
    return file, nil
}

func main() {
    file, err := openFileWithCustomError("test.txt")
    if err != nil {
        fmt.Printf("Custom error: %v\n", err)
        return
    }
    defer file.Close()
    
    fmt.Println("File opened successfully")
}