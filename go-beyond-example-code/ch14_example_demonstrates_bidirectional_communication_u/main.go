package main

import (
    "fmt"
    "time"
)

func main() {
    requestChan := make(chan string)
    responseChan := make(chan string)
    
    // Start server
    go func() {
        for {
            select {
            case req := <-requestChan:
                fmt.Printf("Server received: %s\n", req)
                // Process request and send response
                time.Sleep(100 * time.Millisecond)
                responseChan <- fmt.Sprintf("Response to: %s", req)
            }
        }
    }()
    
    // Start client
    go func() {
        for i := 1; i <= 3; i++ {
            request := fmt.Sprintf("Request %d", i)
            requestChan <- request
            time.Sleep(200 * time.Millisecond)
        }
    }()
    
    // Collect responses
    for i := 0; i < 3; i++ {
        select {
        case resp := <-responseChan:
            fmt.Printf("Client received: %s\n", resp)
        }
    }
    
    time.Sleep(100 * time.Millisecond)
}