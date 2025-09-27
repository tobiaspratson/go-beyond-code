package main

import "fmt"

func main() {
    var celsius float64
    fmt.Print("Enter temperature in Celsius: ")
    fmt.Scanln(&celsius)
    
    fahrenheit := (celsius * 9/5) + 32
    fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
}