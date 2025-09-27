package main

import (
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()
    
    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()
    
    _, err = io.Copy(destFile, sourceFile)
    if err != nil {
        return err
    }
    
    return destFile.Sync()
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <source> <destination>")
        return
    }
    
    src := os.Args[1]
    dst := os.Args[2]
    
    err := copyFile(src, dst)
    if err != nil {
        fmt.Printf("Error copying file: %v\n", err)
        return
    }
    
    fmt.Printf("File copied from %s to %s\n", src, dst)
}