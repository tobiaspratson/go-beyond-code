package main

import "fmt"

func main() {
    var a int32 = 100
    var b int64 = 200
    
    // Convert int32 to int64
    var c int64 = int64(a)
    
    // Convert int64 to int32 (may lose data!)
    var d int32 = int32(b)
    
    // Convert between signed and unsigned
    var e uint32 = uint32(a)
    
    fmt.Printf("a: %d (type: %T)\n", a, a)
    fmt.Printf("c: %d (type: %T)\n", c, c)
    fmt.Printf("d: %d (type: %T)\n", d, d)
    fmt.Printf("e: %d (type: %T)\n", e, e)
}