package main

import (
    "errors"
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "time"
)

// Custom error types for file operations
type FileSystemError struct {
    Operation string
    Path      string
    Err       error
    Time      time.Time
}

func (e FileSystemError) Error() string {
    return fmt.Sprintf("filesystem error in %s for %s at %v: %v", 
        e.Operation, e.Path, e.Time, e.Err)
}

func (e FileSystemError) Unwrap() error {
    return e.Err
}

// File metadata structure
type FileMetadata struct {
    Name    string
    Size    int64
    Mode    fs.FileMode
    ModTime time.Time
    IsDir   bool
}

// File system manager with comprehensive error handling
type FileSystemManager struct {
    basePath string
}

func NewFileSystemManager(basePath string) (*FileSystemManager, error) {
    if basePath == "" {
        return nil, FileSystemError{
            Operation: "INIT",
            Path:      basePath,
            Err:       errors.New("base path cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    // Create base directory if it doesn't exist
    if err := os.MkdirAll(basePath, 0755); err != nil {
        return nil, FileSystemError{
            Operation: "INIT",
            Path:      basePath,
            Err:       fmt.Errorf("failed to create base directory: %w", err),
            Time:      time.Now(),
        }
    }
    
    return &FileSystemManager{basePath: basePath}, nil
}

func (fsm *FileSystemManager) CreateFile(filename, content string) error {
    if filename == "" {
        return FileSystemError{
            Operation: "CREATE_FILE",
            Path:      filename,
            Err:       errors.New("filename cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    fullPath := filepath.Join(fsm.basePath, filename)
    
    // Check if file already exists
    if _, err := os.Stat(fullPath); err == nil {
        return FileSystemError{
            Operation: "CREATE_FILE",
            Path:      fullPath,
            Err:       errors.New("file already exists"),
            Time:      time.Now(),
        }
    }
    
    // Create directory if needed
    dir := filepath.Dir(fullPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return FileSystemError{
            Operation: "CREATE_FILE",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to create directory: %w", err),
            Time:      time.Now(),
        }
    }
    
    // Write file
    if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
        return FileSystemError{
            Operation: "CREATE_FILE",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to write file: %w", err),
            Time:      time.Now(),
        }
    }
    
    return nil
}

func (fsm *FileSystemManager) ReadFile(filename string) (string, error) {
    if filename == "" {
        return "", FileSystemError{
            Operation: "READ_FILE",
            Path:      filename,
            Err:       errors.New("filename cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    fullPath := filepath.Join(fsm.basePath, filename)
    
    data, err := os.ReadFile(fullPath)
    if err != nil {
        return "", FileSystemError{
            Operation: "READ_FILE",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to read file: %w", err),
            Time:      time.Now(),
        }
    }
    
    return string(data), nil
}

func (fsm *FileSystemManager) UpdateFile(filename, content string) error {
    if filename == "" {
        return FileSystemError{
            Operation: "UPDATE_FILE",
            Path:      filename,
            Err:       errors.New("filename cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    fullPath := filepath.Join(fsm.basePath, filename)
    
    // Check if file exists
    if _, err := os.Stat(fullPath); os.IsNotExist(err) {
        return FileSystemError{
            Operation: "UPDATE_FILE",
            Path:      fullPath,
            Err:       errors.New("file does not exist"),
            Time:      time.Now(),
        }
    }
    
    // Write updated content
    if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
        return FileSystemError{
            Operation: "UPDATE_FILE",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to update file: %w", err),
            Time:      time.Now(),
        }
    }
    
    return nil
}

func (fsm *FileSystemManager) DeleteFile(filename string) error {
    if filename == "" {
        return FileSystemError{
            Operation: "DELETE_FILE",
            Path:      filename,
            Err:       errors.New("filename cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    fullPath := filepath.Join(fsm.basePath, filename)
    
    // Check if file exists
    if _, err := os.Stat(fullPath); os.IsNotExist(err) {
        return FileSystemError{
            Operation: "DELETE_FILE",
            Path:      fullPath,
            Err:       errors.New("file does not exist"),
            Time:      time.Now(),
        }
    }
    
    // Delete file
    if err := os.Remove(fullPath); err != nil {
        return FileSystemError{
            Operation: "DELETE_FILE",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to delete file: %w", err),
            Time:      time.Now(),
        }
    }
    
    return nil
}

func (fsm *FileSystemManager) GetFileMetadata(filename string) (*FileMetadata, error) {
    if filename == "" {
        return nil, FileSystemError{
            Operation: "GET_METADATA",
            Path:      filename,
            Err:       errors.New("filename cannot be empty"),
            Time:      time.Now(),
        }
    }
    
    fullPath := filepath.Join(fsm.basePath, filename)
    
    info, err := os.Stat(fullPath)
    if err != nil {
        return nil, FileSystemError{
            Operation: "GET_METADATA",
            Path:      fullPath,
            Err:       fmt.Errorf("failed to get file info: %w", err),
            Time:      time.Now(),
        }
    }
    
    return &FileMetadata{
        Name:    info.Name(),
        Size:    info.Size(),
        Mode:    info.Mode(),
        ModTime: info.ModTime(),
        IsDir:   info.IsDir(),
    }, nil
}

func (fsm *FileSystemManager) ListFiles() ([]*FileMetadata, error) {
    var files []*FileMetadata
    
    err := filepath.Walk(fsm.basePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        // Skip the base directory itself
        if path == fsm.basePath {
            return nil
        }
        
        files = append(files, &FileMetadata{
            Name:    info.Name(),
            Size:    info.Size(),
            Mode:    info.Mode(),
            ModTime: info.ModTime(),
            IsDir:   info.IsDir(),
        })
        
        return nil
    })
    
    if err != nil {
        return nil, FileSystemError{
            Operation: "LIST_FILES",
            Path:      fsm.basePath,
            Err:       fmt.Errorf("failed to list files: %w", err),
            Time:      time.Now(),
        }
    }
    
    return files, nil
}

func main() {
    // Create file system manager
    fsm, err := NewFileSystemManager("./test_files")
    if err != nil {
        fmt.Printf("Failed to create file system manager: %v\n", err)
        return
    }
    fmt.Println("File system manager created successfully")
    
    // Test file operations
    fmt.Println("\n=== File Operations ===")
    
    // Create files
    files := map[string]string{
        "documents/readme.txt": "This is a readme file",
        "config/settings.json": `{"debug": true, "port": 8080}`,
        "logs/app.log":         "Application started\nUser logged in\n",
    }
    
    for filename, content := range files {
        err := fsm.CreateFile(filename, content)
        if err != nil {
            fmt.Printf("Failed to create %s: %v\n", filename, err)
        } else {
            fmt.Printf("Created file: %s\n", filename)
        }
    }
    
    // Read files
    fmt.Println("\n=== Reading Files ===")
    for filename := range files {
        content, err := fsm.ReadFile(filename)
        if err != nil {
            fmt.Printf("Failed to read %s: %v\n", filename, err)
        } else {
            fmt.Printf("Content of %s: %s\n", filename, content)
        }
    }
    
    // Update file
    fmt.Println("\n=== Updating File ===")
    err = fsm.UpdateFile("documents/readme.txt", "This is an updated readme file with more information")
    if err != nil {
        fmt.Printf("Failed to update file: %v\n", err)
    } else {
        fmt.Println("File updated successfully")
    }
    
    // Get file metadata
    fmt.Println("\n=== File Metadata ===")
    for filename := range files {
        metadata, err := fsm.GetFileMetadata(filename)
        if err != nil {
            fmt.Printf("Failed to get metadata for %s: %v\n", filename, err)
        } else {
            fmt.Printf("File: %s, Size: %d bytes, Modified: %v\n", 
                metadata.Name, metadata.Size, metadata.ModTime)
        }
    }
    
    // List all files
    fmt.Println("\n=== All Files ===")
    allFiles, err := fsm.ListFiles()
    if err != nil {
        fmt.Printf("Failed to list files: %v\n", err)
    } else {
        for _, file := range allFiles {
            fmt.Printf("- %s (%d bytes, %v)\n", 
                file.Name, file.Size, file.ModTime)
        }
    }
    
    // Test error conditions
    fmt.Println("\n=== Error Conditions ===")
    
    // Try to read non-existent file
    _, err = fsm.ReadFile("non_existent.txt")
    if err != nil {
        fmt.Printf("Expected error for non-existent file: %v\n", err)
    }
    
    // Try to create file with empty name
    err = fsm.CreateFile("", "content")
    if err != nil {
        fmt.Printf("Expected error for empty filename: %v\n", err)
    }
    
    // Try to update non-existent file
    err = fsm.UpdateFile("non_existent.txt", "content")
    if err != nil {
        fmt.Printf("Expected error for updating non-existent file: %v\n", err)
    }
}