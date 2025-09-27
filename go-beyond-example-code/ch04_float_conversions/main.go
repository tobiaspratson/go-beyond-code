package main

import "fmt"

func main() {
    var f32 float32 = 3.14
    var f64 float64 = 2.718
    
    // Convert float32 to float64
    var converted float64 = float64(f32)
    
    // Convert float64 to float32 (may lose precision!)
    var converted32 float32 = float32(f64)
    
    // Convert int to float
    var integer int = 42
    var float float64 = float64(integer)
    
    fmt.Printf("f32: %f (type: %T)\n", f32, f32)
    fmt.Printf("converted: %f (type: %T)\n", converted, converted)
    fmt.Printf("converted32: %f (type: %T)\n", converted32, converted32)
    fmt.Printf("float: %f (type: %T)\n", float, float)
    
    // Demonstrate precision loss
    var precise float64 = 3.141592653589793
    var lessPrecise float32 = float32(precise)
    
    fmt.Printf("Precise: %.15f\n", precise)
    fmt.Printf("Less precise: %.15f\n", lessPrecise)
    fmt.Printf("Precision lost: %.15f\n", float64(lessPrecise)-precise)
}