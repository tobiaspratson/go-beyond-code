package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Advanced Calculator!")
	fmt.Println("This calculator demonstrates various loop patterns.")

	operationCount := 0

	for {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("Choose an operation:")
		fmt.Println("1. Basic Arithmetic (Add, Subtract, Multiply, Divide)")
		fmt.Println("2. Advanced Math (Power, Square Root, Factorial)")
		fmt.Println("3. Statistics (Sum, Average, Min, Max)")
		fmt.Println("4. Number Sequences (Fibonacci, Prime Check)")
		fmt.Println("5. Pattern Generator")
		fmt.Println("6. Exit")
		fmt.Println(strings.Repeat("=", 50))

		var choice int
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choice)

		if choice == 6 {
			fmt.Printf("Goodbye! You performed %d operations.\n", operationCount)
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice! Please try again.")
			continue
		}

		operationCount++

		switch choice {
		case 1:
			basicArithmetic()
		case 2:
			advancedMath()
		case 3:
			statistics()
		case 4:
			numberSequences()
		case 5:
			patternGenerator()
		}
	}
}

func basicArithmetic() {
	fmt.Println("\n=== Basic Arithmetic ===")
	var a, b float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	if b == 0 {
		fmt.Println("Error: Division by zero!")
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	}
}

func advancedMath() {
	fmt.Println("\n=== Advanced Math ===")
	var n float64
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	fmt.Printf("Power of 2: %.2f\n", math.Pow(n, 2))
	fmt.Printf("Power of 3: %.2f\n", math.Pow(n, 3))
	fmt.Printf("Square root: %.2f\n", math.Sqrt(n))

	if n >= 0 && n == float64(int(n)) {
		// Calculate factorial for integers
		fact := 1
		for i := 1; i <= int(n); i++ {
			fact *= i
		}
		fmt.Printf("Factorial: %d\n", fact)
	}
}

func statistics() {
	fmt.Println("\n=== Statistics Calculator ===")
	fmt.Print("How many numbers? ")
	var count int
	fmt.Scanln(&count)

	if count <= 0 {
		fmt.Println("Invalid count!")
		return
	}

	numbers := make([]float64, count)
	sum := 0.0

	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&numbers[i])
		sum += numbers[i]
	}

	// Find min and max
	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", sum/float64(count))
	fmt.Printf("Min: %.2f\n", min)
	fmt.Printf("Max: %.2f\n", max)
}

func numberSequences() {
	fmt.Println("\n=== Number Sequences ===")
	fmt.Print("Enter a number: ")
	var n int
	fmt.Scanln(&n)

	// Fibonacci sequence
	fmt.Printf("Fibonacci sequence up to %d terms:\n", n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()

	// Prime check
	fmt.Printf("Prime numbers up to %d:\n", n)
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func patternGenerator() {
	fmt.Println("\n=== Pattern Generator ===")
	fmt.Print("Enter pattern size: ")
	var size int
	fmt.Scanln(&size)

	// Triangle pattern
	fmt.Println("Triangle pattern:")
	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Number pyramid
	fmt.Println("\nNumber pyramid:")
	for i := 1; i <= size; i++ {
		for j := 1; j <= size-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
