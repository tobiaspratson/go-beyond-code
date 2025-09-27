package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver - cannot modify the original struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with value receiver - calculates perimeter
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Method with value receiver - returns a string representation
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Method with pointer receiver - CAN modify the original struct
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Method with pointer receiver - sets new dimensions
func (r *Rectangle) SetDimensions(width, height float64) {
    r.Width = width
    r.Height = height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Printf("Original: %s\n", rect.String())
    fmt.Printf("Area: %.2f\n", rect.Area())
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
    
    // Scale the rectangle (modifies the original)
    rect.Scale(2)
    fmt.Printf("After scaling: %s\n", rect.String())
    fmt.Printf("New area: %.2f\n", rect.Area())
    
    // Set new dimensions
    rect.SetDimensions(15, 8)
    fmt.Printf("After setting dimensions: %s\n", rect.String())
    fmt.Printf("Final area: %.2f\n", rect.Area())
}