package main

import "fmt"

// Scenario 2: Nil pointer dereference
func nilPointerExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from nil pointer panic: %v\n", r)
		}
	}()

	var p *int
	fmt.Printf("Value: %d\n", *p) // This will panic
}

// Scenario 3: Type assertion panic
func typeAssertionExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from type assertion panic: %v\n", r)
		}
	}()

	var i interface{} = "hello"
	num := i.(int) // This will panic - string is not int
	fmt.Printf("Number: %d\n", num)
}

func main() {
	fmt.Println("Testing panic scenarios...")

	// Test nil pointer
	nilPointerExample()

	// Test type assertion
	typeAssertionExample()

	fmt.Println("Program continues...")

	// Scenario 1: Index out of bounds
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Printf("Accessing index 10: %d\n", arr[10]) // This will panic

}
