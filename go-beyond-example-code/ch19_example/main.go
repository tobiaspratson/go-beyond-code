package main

import (
    "fmt"
    "math"
)

// Calculate distance between two points
func Distance(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Sqrt(dx*dx + dy*dy)
}

// Calculate angle between two points
func Angle(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Atan2(dy, dx)
}

// Convert angle from radians to degrees
func RadiansToDegrees(radians float64) float64 {
    return radians * 180 / math.Pi
}

func main() {
    // Two points
    x1, y1 := 0.0, 0.0
    x2, y2 := 3.0, 4.0
    
    // Calculate distance and angle
    distance := Distance(x1, y1, x2, y2)
    angle := Angle(x1, y1, x2, y2)
    angleDegrees := RadiansToDegrees(angle)
    
    fmt.Printf("Point 1: (%.1f, %.1f)\n", x1, y1)
    fmt.Printf("Point 2: (%.1f, %.1f)\n", x2, y2)
    fmt.Printf("Distance: %.2f\n", distance)
    fmt.Printf("Angle: %.2f radians (%.2f degrees)\n", angle, angleDegrees)
}