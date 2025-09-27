package main

import (
    "fmt"
    "math"
)

// Convert degrees to radians
func DegreesToRadians(degrees float64) float64 {
    return degrees * math.Pi / 180
}

// Convert radians to degrees
func RadiansToDegrees(radians float64) float64 {
    return radians * 180 / math.Pi
}

// Calculate distance between two points
func Distance(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Sqrt(dx*dx + dy*dy)
}

// Calculate angle between two points (in radians)
func AngleBetweenPoints(x1, y1, x2, y2 float64) float64 {
    dx := x2 - x1
    dy := y2 - y1
    return math.Atan2(dy, dx)
}

// Calculate angle between two vectors
func AngleBetweenVectors(x1, y1, x2, y2 float64) float64 {
    dot := x1*x2 + y1*y2
    mag1 := math.Sqrt(x1*x1 + y1*y1)
    mag2 := math.Sqrt(x2*x2 + y2*y2)
    
    if mag1 == 0 || mag2 == 0 {
        return 0
    }
    
    cosAngle := dot / (mag1 * mag2)
    // Clamp to avoid numerical errors
    cosAngle = math.Max(-1, math.Min(1, cosAngle))
    return math.Acos(cosAngle)
}

// Rotate a point around origin
func RotatePoint(x, y, angleRadians float64) (float64, float64) {
    cos := math.Cos(angleRadians)
    sin := math.Sin(angleRadians)
    
    newX := x*cos - y*sin
    newY := x*sin + y*cos
    
    return newX, newY
}

// Generate sine wave data
func GenerateSineWave(frequency, amplitude, phase, duration, sampleRate float64) []float64 {
    samples := int(duration * sampleRate)
    wave := make([]float64, samples)
    
    for i := 0; i < samples; i++ {
        time := float64(i) / sampleRate
        wave[i] = amplitude * math.Sin(2*math.Pi*frequency*time + phase)
    }
    
    return wave
}

func main() {
    fmt.Println("=== Advanced Trigonometric Applications ===")
    
    // Distance and angle calculations
    fmt.Println("\n--- Distance and Angle Calculations ---")
    x1, y1 := 0.0, 0.0
    x2, y2 := 3.0, 4.0
    
    distance := Distance(x1, y1, x2, y2)
    angle := AngleBetweenPoints(x1, y1, x2, y2)
    angleDegrees := RadiansToDegrees(angle)
    
    fmt.Printf("Point 1: (%.1f, %.1f)\n", x1, y1)
    fmt.Printf("Point 2: (%.1f, %.1f)\n", x2, y2)
    fmt.Printf("Distance: %.2f\n", distance)
    fmt.Printf("Angle: %.4f rad (%.2f°)\n", angle, angleDegrees)
    
    // Vector angle calculation
    fmt.Println("\n--- Vector Angle Calculation ---")
    v1x, v1y := 1.0, 0.0  // Unit vector along x-axis
    v2x, v2y := 1.0, 1.0  // Vector at 45 degrees
    
    vectorAngle := AngleBetweenVectors(v1x, v1y, v2x, v2y)
    vectorAngleDegrees := RadiansToDegrees(vectorAngle)
    
    fmt.Printf("Vector 1: (%.1f, %.1f)\n", v1x, v1y)
    fmt.Printf("Vector 2: (%.1f, %.1f)\n", v2x, v2y)
    fmt.Printf("Angle between vectors: %.4f rad (%.2f°)\n", vectorAngle, vectorAngleDegrees)
    
    // Point rotation
    fmt.Println("\n--- Point Rotation ---")
    px, py := 1.0, 0.0
    rotationAngle := DegreesToRadians(90)  // 90 degrees
    
    newX, newY := RotatePoint(px, py, rotationAngle)
    fmt.Printf("Original point: (%.1f, %.1f)\n", px, py)
    fmt.Printf("Rotated 90°: (%.2f, %.2f)\n", newX, newY)
    
    // Sine wave generation
    fmt.Println("\n--- Sine Wave Generation ---")
    frequency := 1.0    // 1 Hz
    amplitude := 1.0    // Amplitude of 1
    phase := 0.0        // No phase shift
    duration := 1.0     // 1 second
    sampleRate := 10.0  // 10 samples per second
    
    wave := GenerateSineWave(frequency, amplitude, phase, duration, sampleRate)
    fmt.Printf("Generated %d samples of sine wave:\n", len(wave))
    for i, sample := range wave {
        if i < 10 {  // Show first 10 samples
            fmt.Printf("Sample %d: %.4f\n", i, sample)
        }
    }
}