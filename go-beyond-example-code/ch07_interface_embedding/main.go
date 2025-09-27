package main

import "fmt"

// Basic interfaces - small and focused
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type Closer interface {
    Close() error
}

// Embedded interfaces - combine multiple behaviors
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Implementation that satisfies all interfaces
type File struct {
    Name    string
    content string
    closed  bool
}

func (f *File) Read() string {
    if f.closed {
        return "File is closed"
    }
    return fmt.Sprintf("Reading from %s: %s", f.Name, f.content)
}

func (f *File) Write(data string) {
    if f.closed {
        fmt.Println("Cannot write to closed file")
        return
    }
    f.content = data
    fmt.Printf("Writing to %s: %s\n", f.Name, data)
}

func (f *File) Close() error {
    f.closed = true
    fmt.Printf("Closing %s\n", f.Name)
    return nil
}

// Function that works with any Reader
func readData(r Reader) {
    fmt.Printf("Read result: %s\n", r.Read())
}

// Function that works with any Writer
func writeData(w Writer, data string) {
    w.Write(data)
}

// Function that works with any ReadWriter
func copyData(rw ReadWriter, data string) {
    fmt.Printf("Original: %s\n", rw.Read())
    rw.Write(data)
    fmt.Printf("After write: %s\n", rw.Read())
}

// Function that works with any ReadWriteCloser
func processResource(rwc ReadWriteCloser, data string) {
    fmt.Printf("Initial: %s\n", rwc.Read())
    rwc.Write(data)
    fmt.Printf("After write: %s\n", rwc.Read())
    rwc.Close()
    fmt.Printf("After close: %s\n", rwc.Read())
}

func main() {
    file := &File{Name: "document.txt", content: "Hello, World!"}
    
    fmt.Println("=== FILE OPERATIONS ===")
    // Use as individual interfaces
    readData(file)
    writeData(file, "New file content")
    
    fmt.Println("\n=== COPY OPERATION ===")
    copyData(file, "Copied content")
    
    fmt.Println("\n=== PROCESS FILE ===")
    processResource(file, "Final content")
}