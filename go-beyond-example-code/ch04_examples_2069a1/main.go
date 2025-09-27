package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("Simple Calculator")
    fmt.Println("Enter two numbers and an operation (+, -, *, /)")
    
    var input1, input2, operation string
    
    fmt.Print("First number: ")
    fmt.Scanln(&input1)
    
    fmt.Print("Operation (+, -, *, /): ")
    fmt.Scanln(&operation)
    
    fmt.Print("Second number: ")
    fmt.Scanln(&input2)
    
    // Convert strings to numbers
    num1, err1 := strconv.ParseFloat(input1, 64)
    num2, err2 := strconv.ParseFloat(input2, 64)
    
    if err1 != nil || err2 != nil {
        fmt.Println("Error: Invalid numbers")
        return
    }
    
    var result float64
    
    switch operation {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 == 0 {
            fmt.Println("Error: Division by zero")
            return
        }
        result = num1 / num2
    default:
        fmt.Println("Error: Invalid operation")
        return
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operation, num2, result)
}