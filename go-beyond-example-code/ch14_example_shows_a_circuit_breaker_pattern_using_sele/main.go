package main

import (
    "fmt"
    "time"
)

type RateLimiter struct {
    tokens    chan struct{}
    rate      time.Duration
    burst     int
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
    rl := &RateLimiter{
        tokens: make(chan struct{}, burst),
        rate:   rate,
        burst:  burst,
    }
    
    // Fill initial tokens
    for i := 0; i < burst; i++ {
        rl.tokens <- struct{}{}
    }
    
    // Start token refill goroutine
    go rl.refill()
    
    return rl
}

func (rl *RateLimiter) refill() {
    ticker := time.NewTicker(rl.rate)
    defer ticker.Stop()
    
    for range ticker.C {
        select {
        case rl.tokens <- struct{}{}:
            // Token added
        default:
            // Bucket is full
        }
    }
}

func (rl *RateLimiter) Allow() bool {
    select {
    case <-rl.tokens:
        return true
    default:
        return false
    }
}

func (rl *RateLimiter) Wait() {
    <-rl.tokens
}

func main() {
    // Create rate limiter: 2 requests per second, burst of 5
    limiter := NewRateLimiter(500*time.Millisecond, 5)
    
    // Simulate requests
    for i := 0; i < 10; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
        time.Sleep(100 * time.Millisecond)
    }
    
    // Wait and try again
    time.Sleep(2 * time.Second)
    fmt.Println("After waiting...")
    
    for i := 0; i < 5; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
        time.Sleep(100 * time.Millisecond)
    }
}