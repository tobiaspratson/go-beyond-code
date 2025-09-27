package main

import "fmt"

// Function that only sends
func sendOnly(ch chan<- string) {
    ch <- "Hello from send-only function"
}

// Function that only receives
func receiveOnly(ch <-chan string) {
    msg := <-ch
    fmt.Printf("Received: %s\n", msg)
}

func main() {
    ch := make(chan string)
    
    // Send in one goroutine
    go sendOnly(ch)
    
    // Receive in main goroutine
    receiveOnly(ch)
}