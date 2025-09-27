package main

import "fmt"

func main() {
    // File permissions (Unix-style)
    const (
        ReadPermission    = 1 << 2  // 100 (binary)
        WritePermission   = 1 << 1  // 010 (binary)
        ExecutePermission = 1 << 0  // 001 (binary)
    )
    
    // Set permissions
    permissions := ReadPermission | WritePermission  // 110 (binary) = 6 (decimal)
    fmt.Printf("Permissions: %o (octal)\n", permissions)
    
    // Check permissions
    canRead := (permissions & ReadPermission) != 0
    canWrite := (permissions & WritePermission) != 0
    canExecute := (permissions & ExecutePermission) != 0
    
    fmt.Printf("Can read: %t\n", canRead)
    fmt.Printf("Can write: %t\n", canWrite)
    fmt.Printf("Can execute: %t\n", canExecute)
    
    // Add execute permission
    permissions |= ExecutePermission
    fmt.Printf("After adding execute: %o\n", permissions)
    
    // Remove write permission
    permissions &^= WritePermission
    fmt.Printf("After removing write: %o\n", permissions)
}