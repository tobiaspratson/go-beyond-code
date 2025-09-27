package main

import "fmt"

func main() {
	fmt.Println("=== Continue with Data Validation ===")
	// Skip invalid data
	data := []int{1, -2, 3, 0, 5, -6, 7, 8, 0, 10}
	sum := 0
	count := 0

	for _, value := range data {
		if value <= 0 {
			fmt.Printf("Skipping invalid value: %d\n", value)
			continue
		}
		sum += value
		count++
		fmt.Printf("Added %d, sum is now %d\n", value, sum)
	}
	fmt.Printf("Final sum: %d, valid numbers: %d\n", sum, count)

	fmt.Println("\n=== Continue with String Processing ===")
	// Skip certain characters
	text := "Hello, World! 123"
	consonants := ""

	for _, char := range text {
		if char >= '0' && char <= '9' {
			fmt.Printf("Skipping digit: %c\n", char)
			continue
		}
		if char == ' ' || char == ',' || char == '!' {
			fmt.Printf("Skipping punctuation: %c\n", char)
			continue
		}
		consonants += string(char)
	}
	fmt.Printf("Consonants only: %s\n", consonants)

	fmt.Println("\n=== Continue with Early Exit Logic ===")
	// Process until we find what we need
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 50
	found := false

	for i, num := range numbers {
		if num < target {
			fmt.Printf("Skipping %d (too small)\n", num)
			continue
		}
		if num == target {
			fmt.Printf("Found target %d at index %d!\n", target, i)
			found = true
			break
		}
		fmt.Printf("Processing %d (greater than target)\n", num)
	}

	if !found {
		fmt.Printf("Target %d not found\n", target)
	}
}
