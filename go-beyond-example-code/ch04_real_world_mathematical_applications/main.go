package main

import (
	"fmt"
	"math"
)

// Financial calculations
func calculateCompoundInterest(principal, rate float64, years int) float64 {
	return principal * math.Pow(1+rate, float64(years))
}

// Distance calculation (Pythagorean theorem)
func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

// Temperature conversion
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Statistical functions
func mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func standardDeviation(numbers []float64) float64 {
	m := mean(numbers)
	sum := 0.0
	for _, num := range numbers {
		diff := num - m
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(numbers)))
}

func main() {

	principal := 1000.0
	rate := 0.05 // 5% annual interest
	years := 10

	finalAmount := calculateCompoundInterest(principal, rate, years)
	fmt.Printf("Principal: $%.2f\n", principal)
	fmt.Printf("Rate: %.1f%%\n", rate*100)
	fmt.Printf("Years: %d\n", years)
	fmt.Printf("Final amount: $%.2f\n", finalAmount)

	tempC := 25.0
	tempF := celsiusToFahrenheit(tempC)
	tempCBack := fahrenheitToCelsius(tempF)

	fmt.Printf("%.1f°C = %.1f°F\n", tempC, tempF)
	fmt.Printf("%.1f°F = %.1f°C\n", tempF, tempCBack)
	dist := distance(0, 0, 3, 4)
	fmt.Printf("Distance from (0,0) to (3,4): %.2f\n", dist)
	data := []float64{1, 2, 3, 4, 5}
	avg := mean(data)
	stdDev := standardDeviation(data)

	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Mean: %.2f\n", avg)
	fmt.Printf("Standard deviation: %.2f\n", stdDev)
}
