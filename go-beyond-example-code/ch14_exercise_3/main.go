package main

import (
    "fmt"
    "time"
)

type Request struct {
    ID   int
    Data string
}

type Response struct {
    ID     int
    Result string
}

func handleRequest(req Request, timeout time.Duration) (Response, error) {
    // Simulate processing time
    time.Sleep(200 * time.Millisecond)
    
    return Response{
        ID:     req.ID,
        Result: fmt.Sprintf("Processed: %s", req.Data),
    }, nil
}

func server(requests <-chan Request, responses chan<- Response) {
    for req := range requests {
        // Handle request with timeout
        select {
        case <-time.After(100 * time.Millisecond):
            // Timeout
            responses <- Response{
                ID:     req.ID,
                Result: "Request timeout",
            }
        default:
            // Process request
            resp, _ := handleRequest(req, 100*time.Millisecond)
            responses <- resp
        }
    }
}

func main() {
    requests := make(chan Request, 5)
    responses := make(chan Response, 5)
    
    // Start server
    go server(requests, responses)
    
    // Send requests
    for i := 1; i <= 5; i++ {
        req := Request{
            ID:   i,
            Data: fmt.Sprintf("Request data %d", i),
        }
        requests <- req
    }
    close(requests)
    
    // Collect responses
    for i := 0; i < 5; i++ {
        resp := <-responses
        fmt.Printf("Response: %+v\n", resp)
    }
}