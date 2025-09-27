package main

import (
    "fmt"
    "os"
    "path/filepath"
)

type FileSystem interface {
    Open(name string) (*os.File, error)
    Stat(name string) (os.FileInfo, error)
    ReadDir(name string) ([]os.DirEntry, error)
}

type RealFileSystem struct{}

func (fs RealFileSystem) Open(name string) (*os.File, error) {
    return os.Open(name)
}

func (fs RealFileSystem) Stat(name string) (os.FileInfo, error) {
    return os.Stat(name)
}

func (fs RealFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
    return os.ReadDir(name)
}

func processFilesWithFS(fs FileSystem, dir string) error {
    entries, err := fs.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    
    for _, entry := range entries {
        if !entry.IsDir() {
            file, err := fs.Open(filepath.Join(dir, entry.Name()))
            if err != nil {
                return fmt.Errorf("failed to open file: %w", err)
            }
            defer file.Close()
            
            // Process file...
            fmt.Printf("Processing %s\n", entry.Name())
        }
    }
    
    return nil
}

func main() {
    fs := RealFileSystem{}
    err := processFilesWithFS(fs, ".")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}