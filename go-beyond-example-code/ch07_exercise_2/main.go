package main

import "fmt"
import "math"

// Shape interface with multiple methods
type Shape interface {
    Area() float64
    Perimeter() float64
    Name() string
    String() string
}

// Rectangle implementation
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

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Circle implementation
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (c Circle) Name() string {
    return "Circle"
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Triangle implementation
type Triangle struct {
    Base   float64
    Height float64
    Side1  float64
    Side2  float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
    return t.Base + t.Side1 + t.Side2
}

func (t Triangle) Name() string {
    return "Triangle"
}

func (t Triangle) String() string {
    return fmt.Sprintf("Triangle{Base: %.2f, Height: %.2f}", t.Base, t.Height)
}

// Shape calculator with advanced operations
type ShapeCalculator struct {
    Shapes []Shape
}

func (sc *ShapeCalculator) AddShape(shape Shape) {
    sc.Shapes = append(sc.Shapes, shape)
}

func (sc ShapeCalculator) TotalArea() float64 {
    total := 0.0
    for _, shape := range sc.Shapes {
        total += shape.Area()
    }
    return total
}

func (sc ShapeCalculator) TotalPerimeter() float64 {
    total := 0.0
    for _, shape := range sc.Shapes {
        total += shape.Perimeter()
    }
    return total
}

func (sc ShapeCalculator) ListShapes() {
    fmt.Println("=== Shape Collection ===")
    for i, shape := range sc.Shapes {
        fmt.Printf("%d. %s - Area: %.2f, Perimeter: %.2f\n", 
            i+1, shape.String(), shape.Area(), shape.Perimeter())
    }
}

func (sc ShapeCalculator) GetShapeStats() {
    fmt.Printf("\n=== Shape Statistics ===\n")
    fmt.Printf("Total shapes: %d\n", len(sc.Shapes))
    fmt.Printf("Total area: %.2f\n", sc.TotalArea())
    fmt.Printf("Total perimeter: %.2f\n", sc.TotalPerimeter())
    
    if len(sc.Shapes) > 0 {
        fmt.Printf("Average area: %.2f\n", sc.TotalArea()/float64(len(sc.Shapes)))
        fmt.Printf("Average perimeter: %.2f\n", sc.TotalPerimeter()/float64(len(sc.Shapes)))
    }
}

func (sc ShapeCalculator) FindLargestArea() Shape {
    if len(sc.Shapes) == 0 {
        return nil
    }
    
    largest := sc.Shapes[0]
    for _, shape := range sc.Shapes {
        if shape.Area() > largest.Area() {
            largest = shape
        }
    }
    return largest
}

func (sc ShapeCalculator) FindSmallestPerimeter() Shape {
    if len(sc.Shapes) == 0 {
        return nil
    }
    
    smallest := sc.Shapes[0]
    for _, shape := range sc.Shapes {
        if shape.Perimeter() < smallest.Perimeter() {
            smallest = shape
        }
    }
    return smallest
}

func main() {
    // Create shape calculator
    calculator := ShapeCalculator{}
    
    // Add various shapes
    calculator.AddShape(Rectangle{Width: 10, Height: 5})
    calculator.AddShape(Circle{Radius: 7})
    calculator.AddShape(Triangle{Base: 8, Height: 6, Side1: 5, Side2: 7})
    calculator.AddShape(Rectangle{Width: 3, Height: 4})
    calculator.AddShape(Circle{Radius: 3})
    
    // Display all shapes
    calculator.ListShapes()
    
    // Show statistics
    calculator.GetShapeStats()
    
    // Find extremes
    largest := calculator.FindLargestArea()
    if largest != nil {
        fmt.Printf("\nLargest area: %s (%.2f)\n", largest.String(), largest.Area())
    }
    
    smallest := calculator.FindSmallestPerimeter()
    if smallest != nil {
        fmt.Printf("Smallest perimeter: %s (%.2f)\n", smallest.String(), smallest.Perimeter())
    }
}