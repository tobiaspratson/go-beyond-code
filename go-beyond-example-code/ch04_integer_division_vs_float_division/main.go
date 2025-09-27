package main

import "fmt"

// Safe division function
func safeDivide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false // Division by zero
	}
	return a / b, true
}

func main() {
	// Integer division truncates (no rounding)
	fmt.Printf("7 / 3 = %d (integer division)\n", 7/3)
	fmt.Printf("7 / 3 = %.2f (float division)\n", 7.0/3.0)

	// Negative integer division
	fmt.Printf("-7 / 3 = %d (integer division)\n", -7/3)
	fmt.Printf("-7 / 3 = %.2f (float division)\n", -7.0/3.0)

	// Modulo with negative numbers
	fmt.Printf("7 %% 3 = %d\n", 7%3)
	fmt.Printf("-7 %% 3 = %d\n", -7%3)
	fmt.Printf("7 %% -3 = %d\n", 7%-3)
	fmt.Printf("-7 %% -3 = %d\n", -7%-3)

	// Division by zero (runtime panic!)
	// fmt.Printf("7 / 0 = %d\n", 7/0)  // This would panic!

	result, ok := safeDivide(7, 3)
	if ok {
		fmt.Printf("Safe division: 7 / 3 = %d\n", result)
	}

	result, ok = safeDivide(7, 0)
	if !ok {
		fmt.Println("Safe division: Cannot divide by zero")
	}
}
