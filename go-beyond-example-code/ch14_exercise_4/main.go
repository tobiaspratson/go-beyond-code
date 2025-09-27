package main

import (
    "fmt"
    "time"
)

type Request struct {
    ID      int
    Data    string
    Timeout time.Duration
}

type Response struct {
    ID     int
    Result string
    Error  string
}

func processRequest(req Request) (Response, error) {
    // Simulate variable processing time
    processingTime := time.Duration(req.ID*100) * time.Millisecond
    time.Sleep(processingTime)
    
    return Response{
        ID:     req.ID,
        Result: fmt.Sprintf("Processed: %s", req.Data),
    }, nil
}

func advancedServer(requests <-chan Request, responses chan<- Response) {
    for req := range requests {
        resultChan := make(chan Response, 1)
        errorChan := make(chan error, 1)
        
        // Start processing
        go func() {
            resp, err := processRequest(req)
            if err != nil {
                errorChan <- err
            } else {
                resultChan <- resp
            }
        }()
        
        // Wait for result or timeout
        select {
        case resp := <-resultChan:
            responses <- resp
        case err := <-errorChan:
            responses <- Response{
                ID:    req.ID,
                Error: err.Error(),
            }
        case <-time.After(req.Timeout):
            responses <- Response{
                ID:    req.ID,
                Error: "Request timeout",
            }
        }
    }
}

func main() {
    requests := make(chan Request, 10)
    responses := make(chan Response, 10)
    
    // Start server
    go advancedServer(requests, responses)
    
    // Send requests with different timeouts
    timeouts := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}
    for i := 1; i <= 6; i++ {
        req := Request{
            ID:      i,
            Data:    fmt.Sprintf("Request data %d", i),
            Timeout: timeouts[i%len(timeouts)],
        }
        requests <- req
    }
    close(requests)
    
    // Collect responses
    for i := 0; i < 6; i++ {
        resp := <-responses
        if resp.Error != "" {
            fmt.Printf("Response %d: Error - %s\n", resp.ID, resp.Error)
        } else {
            fmt.Printf("Response %d: %s\n", resp.ID, resp.Result)
        }
    }
}