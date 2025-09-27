package main

import "fmt"

// Base interfaces
type Reader interface {
	Read() string
}

type Writer interface {
	Write() string
}

type Closer interface {
	Close() error
}

// Base structs with specific functionality
type FileReader struct {
	Name string
}

func (fr FileReader) Read() string {
	return fmt.Sprintf("Reading from %s", fr.Name)
}

type FileWriter struct {
	Name string
}

func (fw FileWriter) Write() string {
	return fmt.Sprintf("Writing to %s", fw.Name)
}

type FileCloser struct {
	Name string
}

func (fc FileCloser) Close() error {
	fmt.Printf("Closing %s\n", fc.Name)
	return nil
}

// Multiple embedding - File gets all methods from embedded types
type File struct {
	FileReader // Embedded - methods promoted to File
	FileWriter // Embedded - methods promoted to File
	FileCloser // Embedded - methods promoted to File
	Size       int64
}

// File's own methods
func (f File) GetSize() int64 {
	return f.Size
}

func (f File) SetSize(size int64) {
	f.Size = size
}

// Method that uses embedded functionality
func (f File) Copy() string {
	content := f.Read()      // Uses FileReader's Read method
	writeResult := f.Write() // Uses FileWriter's Write method
	return fmt.Sprintf("Copied: %s -> %s", content, writeResult)
}

// Method that demonstrates method resolution
func (f File) String() string {
	return fmt.Sprintf("File: %s, Size: %d bytes", f.FileReader.Name, f.Size)
}

func main() {
	file := File{
		FileReader: FileReader{Name: "source.txt"},
		FileWriter: FileWriter{Name: "dest.txt"},
		FileCloser: FileCloser{Name: "dest.txt"},
		Size:       1024,
	}

	// Use embedded methods directly on File
	fmt.Println("=== Using Embedded Methods ===")
	fmt.Println(file.Read())  // Calls FileReader.Read()
	fmt.Println(file.Write()) // Calls FileWriter.Write()
	file.Close()              // Calls FileCloser.Close()

	// Use combined functionality
	fmt.Println("\n=== Combined Functionality ===")
	fmt.Println(file.Copy())
	fmt.Printf("File size: %d bytes\n", file.GetSize())
	fmt.Println(file.String())

	// Demonstrate interface satisfaction
	fmt.Println("\n=== Interface Satisfaction ===")
	var reader Reader = file
	var writer Writer = file
	var closer Closer = file

	fmt.Println("As Reader:", reader.Read())
	fmt.Println("As Writer:", writer.Write())
	closer.Close()
}
