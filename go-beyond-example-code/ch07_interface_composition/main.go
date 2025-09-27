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

// Composed interfaces - combine multiple behaviors
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// File implements all interfaces
type File struct {
    content string
    closed  bool
}

func (f *File) Read() string {
    if f.closed {
        return "File is closed"
    }
    return f.content
}

func (f *File) Write(data string) {
    if f.closed {
        fmt.Println("Cannot write to closed file")
        return
    }
    f.content = data
}

func (f *File) Close() error {
    f.closed = true
    return nil
}

// Function that works with any Reader
func readFromSource(r Reader) {
    fmt.Printf("Reading: %s\n", r.Read())
}

// Function that works with any Writer
func writeToDestination(w Writer, data string) {
    w.Write(data)
}

// Function that works with any ReadWriter
func copyData(rw ReadWriter) {
    data := rw.Read()
    rw.Write("Modified: " + data)
}

// Function that works with any ReadWriteCloser
func processFile(rwc ReadWriteCloser) {
    fmt.Printf("Initial content: %s\n", rwc.Read())
    rwc.Write("New content")
    fmt.Printf("After write: %s\n", rwc.Read())
    rwc.Close()
    fmt.Printf("After close: %s\n", rwc.Read())
}

func main() {
    file := &File{content: "Hello, World!"}
    
    // Use as individual interfaces
    fmt.Println("=== Using as Reader ===")
    readFromSource(file)
    
    fmt.Println("\n=== Using as Writer ===")
    writeToDestination(file, "Updated content")
    readFromSource(file)
    
    fmt.Println("\n=== Using as ReadWriter ===")
    copyData(file)
    readFromSource(file)
    
    fmt.Println("\n=== Using as ReadWriteCloser ===")
    processFile(file)
}