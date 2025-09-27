package main

import "fmt"

// Define an interface - this is a contract
type Shape interface {
    Area() float64
    Perimeter() float64
    Name() string  // Added method for better demonstration
}

// Rectangle implements Shape (automatically!)
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) Name() string {
    return "Rectangle"
}

// Circle implements Shape (automatically!)
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func (c Circle) Name() string {
    return "Circle"
}

// Triangle implements Shape (automatically!)
type Triangle struct {
    Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
    // Simplified - assumes right triangle
    return t.Base + t.Height + (t.Base*t.Base + t.Height*t.Height)
}

func (t Triangle) Name() string {
    return "Triangle"
}

// Function that works with ANY Shape
func printShapeInfo(s Shape) {
    fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", 
        s.Name(), s.Area(), s.Perimeter())
}

// Function that calculates total area of multiple shapes
func totalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func main() {
    // Create different shapes
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}
    triangle := Triangle{Base: 8, Height: 6}
    
    // All shapes can be used as Shape interface
    shapes := []Shape{rect, circle, triangle}
    
    fmt.Println("Individual shapes:")
    for _, shape := range shapes {
        printShapeInfo(shape)
    }
    
    fmt.Printf("\nTotal area of all shapes: %.2f\n", totalArea(shapes))
}