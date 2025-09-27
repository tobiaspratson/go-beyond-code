package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Print("What's your name? ")
    
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    
    fmt.Printf("Hello, %s! Welcome to Go programming!\n", name)
}