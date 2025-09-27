package main

import (
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    requests chan time.Time
    rate     time.Duration
    burst    int
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
    rl := &RateLimiter{
        requests: make(chan time.Time, burst),
        rate:     rate,
        burst:    burst,
    }
    
    // Start the rate limiter
    go rl.run()
    
    return rl
}

func (rl *RateLimiter) run() {
    ticker := time.NewTicker(rl.rate)
    defer ticker.Stop()
    
    for range ticker.C {
        select {
        case <-rl.requests:
            // Remove one request from the queue
        default:
            // No requests to process
        }
    }
}

func (rl *RateLimiter) Allow() bool {
    select {
    case rl.requests <- time.Now():
        return true
    default:
        return false
    }
}

func worker(id int, rateLimiter *RateLimiter, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 10; i++ {
        if rateLimiter.Allow() {
            fmt.Printf("Worker %d: Request %d allowed\n", id, i+1)
        } else {
            fmt.Printf("Worker %d: Request %d rate limited\n", id, i+1)
        }
        time.Sleep(50 * time.Millisecond)
    }
}

func main() {
    // Create rate limiter: 2 requests per second, burst of 5
    rateLimiter := NewRateLimiter(500*time.Millisecond, 5)
    
    var wg sync.WaitGroup
    
    // Start workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, rateLimiter, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}