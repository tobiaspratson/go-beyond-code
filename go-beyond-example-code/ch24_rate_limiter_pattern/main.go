package main

import (
    "fmt"
    "sync"
    "time"
)

// Rate limiter
type RateLimiter struct {
    limit    int
    interval time.Duration
    tokens   int
    lastRefill time.Time
    mu       sync.Mutex
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
    return &RateLimiter{
        limit:    limit,
        interval: interval,
        tokens:   limit,
        lastRefill: time.Now(),
    }
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    // Refill tokens if needed
    now := time.Now()
    if now.Sub(rl.lastRefill) >= rl.interval {
        rl.tokens = rl.limit
        rl.lastRefill = now
    }
    
    // Check if we have tokens
    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    
    return false
}

func (rl *RateLimiter) Wait() {
    for !rl.Allow() {
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {
    // Create rate limiter: 5 requests per second
    limiter := NewRateLimiter(5, time.Second)
    
    // Simulate requests
    for i := 0; i < 20; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
        
        time.Sleep(100 * time.Millisecond)
    }
    
    // Wait for rate limit to reset
    fmt.Println("Waiting for rate limit to reset...")
    time.Sleep(2 * time.Second)
    
    // Try again
    for i := 0; i < 5; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed\n", i+1)
        } else {
            fmt.Printf("Request %d: Rate limited\n", i+1)
        }
    }
}