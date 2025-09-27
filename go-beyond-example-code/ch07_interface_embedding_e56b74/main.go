package main

import "fmt"

// Base interfaces
type Reader interface {
    Read() string
}

type Writer interface {
    Write() string
}

// Embedded interface
type ReadWriter interface {
    Reader
    Writer
}

// Implementation
type File struct {
    Name string
}

func (f File) Read() string {
    return fmt.Sprintf("Reading from %s", f.Name)
}

func (f File) Write() string {
    return fmt.Sprintf("Writing to %s", f.Name)
}

func main() {
    file := File{Name: "document.txt"}
    
    // Use as ReadWriter
    var rw ReadWriter = file
    fmt.Println(rw.Read())
    fmt.Println(rw.Write())
    
    // Use as Reader
    var r Reader = file
    fmt.Println(r.Read())
    
    // Use as Writer
    var w Writer = file
    fmt.Println(w.Write())
}