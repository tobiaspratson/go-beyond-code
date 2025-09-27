package main

import (
	"fmt"
	"math"
	"sort"
)

// Calculate mean
func Mean(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("cannot calculate mean of empty dataset")
	}

	sum := 0.0
	for _, x := range data {
		sum += x
	}
	return sum / float64(len(data)), nil
}

// Calculate variance (population variance)
func Variance(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("cannot calculate variance of empty dataset")
	}
	if len(data) == 1 {
		return 0, nil
	}

	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for _, x := range data {
		sum += math.Pow(x-mean, 2)
	}
	return sum / float64(len(data)), nil
}

// Calculate sample variance (Bessel's correction)
func SampleVariance(data []float64) (float64, error) {
	if len(data) <= 1 {
		return 0, fmt.Errorf("sample variance requires at least 2 data points")
	}

	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for _, x := range data {
		sum += math.Pow(x-mean, 2)
	}
	return sum / float64(len(data)-1), nil
}

// Calculate standard deviation
func StandardDeviation(data []float64) (float64, error) {
	variance, err := Variance(data)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}

// Calculate median
func Median(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("cannot calculate median of empty dataset")
	}

	// Create a copy to avoid modifying original data
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	n := len(sortedData)
	if n%2 == 0 {
		return (sortedData[n/2-1] + sortedData[n/2]) / 2, nil
	}
	return sortedData[n/2], nil
}

// Calculate mode (most frequent value)
func Mode(data []float64) ([]float64, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("cannot calculate mode of empty dataset")
	}

	frequency := make(map[float64]int)
	for _, x := range data {
		frequency[x]++
	}

	maxFreq := 0
	for _, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	var modes []float64
	for value, freq := range frequency {
		if freq == maxFreq {
			modes = append(modes, value)
		}
	}

	return modes, nil
}

// Calculate range
func Range(data []float64) (float64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("cannot calculate range of empty dataset")
	}

	min := data[0]
	max := data[0]

	for _, x := range data {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return max - min, nil
}

// Calculate quartiles
func Quartiles(data []float64) (float64, float64, float64, error) {
	if len(data) == 0 {
		return 0, 0, 0, fmt.Errorf("cannot calculate quartiles of empty dataset")
	}

	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	q1 := percentile(sortedData, 25)
	q2 := percentile(sortedData, 50) // Same as median
	q3 := percentile(sortedData, 75)

	return q1, q2, q3, nil
}

// Calculate percentile
func percentile(data []float64, p float64) float64 {
	if p <= 0 {
		return data[0]
	}
	if p >= 100 {
		return data[len(data)-1]
	}

	n := len(data)
	index := p / 100 * float64(n-1)

	if index == float64(int(index)) {
		return data[int(index)]
	}

	lower := int(index)
	upper := lower + 1
	weight := index - float64(lower)

	return data[lower]*(1-weight) + data[upper]*weight
}

// Calculate skewness
func Skewness(data []float64) (float64, error) {
	if len(data) < 3 {
		return 0, fmt.Errorf("skewness requires at least 3 data points")
	}

	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	stdDev, err := StandardDeviation(data)
	if err != nil {
		return 0, err
	}

	if stdDev == 0 {
		return 0, nil // All values are the same
	}

	sum := 0.0
	for _, x := range data {
		sum += math.Pow((x-mean)/stdDev, 3)
	}

	return sum / float64(len(data)), nil
}

// Calculate kurtosis
func Kurtosis(data []float64) (float64, error) {
	if len(data) < 4 {
		return 0, fmt.Errorf("kurtosis requires at least 4 data points")
	}

	mean, err := Mean(data)
	if err != nil {
		return 0, err
	}

	stdDev, err := StandardDeviation(data)
	if err != nil {
		return 0, err
	}

	if stdDev == 0 {
		return 0, nil // All values are the same
	}

	sum := 0.0
	for _, x := range data {
		sum += math.Pow((x-mean)/stdDev, 4)
	}

	return sum/float64(len(data)) - 3, nil // Excess kurtosis
}

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("=== Statistical Analysis ===")
	fmt.Printf("Data: %v\n", data)

	// Basic statistics
	mean, _ := Mean(data)
	variance, _ := Variance(data)
	stdDev, _ := StandardDeviation(data)
	median, _ := Median(data)
	rangeVal, _ := Range(data)

	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Range: %.2f\n", rangeVal)

	// Advanced statistics
	modes, _ := Mode(data)
	fmt.Printf("Mode(s): %v\n", modes)

	q1, q2, q3, _ := Quartiles(data)
	fmt.Printf("Quartiles - Q1: %.2f, Q2: %.2f, Q3: %.2f\n", q1, q2, q3)

	skewness, _ := Skewness(data)
	kurtosis, _ := Kurtosis(data)
	fmt.Printf("Skewness: %.4f\n", skewness)
	fmt.Printf("Kurtosis: %.4f\n", kurtosis)

	// Test with different dataset
	fmt.Println("\n=== Different Dataset ===")
	data2 := []float64{1, 1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 5}
	fmt.Printf("Data: %v\n", data2)

	mean2, _ := Mean(data2)
	modes2, _ := Mode(data2)
	skewness2, _ := Skewness(data2)

	fmt.Printf("Mean: %.2f\n", mean2)
	fmt.Printf("Mode(s): %v\n", modes2)
	fmt.Printf("Skewness: %.4f\n", skewness2)
}
