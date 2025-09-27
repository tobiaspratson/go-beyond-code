package main

import (
    "fmt"
    "sync"
    "time"
)

// Rate limiter with multiple algorithms
type RateLimiter struct {
    limit      int
    interval   time.Duration
    tokens     int
    lastRefill time.Time
    mu         sync.Mutex
    algorithm  string
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
    return &RateLimiter{
        limit:      limit,
        interval:   interval,
        tokens:     limit,
        lastRefill: time.Now(),
        algorithm:  "token_bucket",
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

func (rl *RateLimiter) GetTokens() int {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    return rl.tokens
}

func (rl *RateLimiter) GetLimit() int {
    return rl.limit
}

func (rl *RateLimiter) GetInterval() time.Duration {
    return rl.interval
}

// Sliding window rate limiter
type SlidingWindowRateLimiter struct {
    limit     int
    window    time.Duration
    requests  []time.Time
    mu        sync.Mutex
}

func NewSlidingWindowRateLimiter(limit int, window time.Duration) *SlidingWindowRateLimiter {
    return &SlidingWindowRateLimiter{
        limit:    limit,
        window:   window,
        requests: make([]time.Time, 0),
    }
}

func (swrl *SlidingWindowRateLimiter) Allow() bool {
    swrl.mu.Lock()
    defer swrl.mu.Unlock()
    
    now := time.Now()
    
    // Remove old requests outside the window
    cutoff := now.Add(-swrl.window)
    for len(swrl.requests) > 0 && swrl.requests[0].Before(cutoff) {
        swrl.requests = swrl.requests[1:]
    }
    
    // Check if we're under the limit
    if len(swrl.requests) < swrl.limit {
        swrl.requests = append(swrl.requests, now)
        return true
    }
    
    return false
}

func (swrl *SlidingWindowRateLimiter) Wait() {
    for !swrl.Allow() {
        time.Sleep(10 * time.Millisecond)
    }
}

func (swrl *SlidingWindowRateLimiter) GetRequestCount() int {
    swrl.mu.Lock()
    defer swrl.mu.Unlock()
    return len(swrl.requests)
}

func main() {
    fmt.Println("=== Rate Limiter Examples ===")
    
    // Token bucket rate limiter
    fmt.Println("\n--- Token Bucket Rate Limiter ---")
    limiter := NewRateLimiter(5, time.Second)
    
    for i := 0; i < 20; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed (tokens: %d)\n", i+1, limiter.GetTokens())
        } else {
            fmt.Printf("Request %d: Rate limited (tokens: %d)\n", i+1, limiter.GetTokens())
        }
        
        time.Sleep(100 * time.Millisecond)
    }
    
    // Wait for rate limit to reset
    fmt.Println("Waiting for rate limit to reset...")
    time.Sleep(2 * time.Second)
    
    // Try again
    for i := 0; i < 5; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d: Allowed (tokens: %d)\n", i+1, limiter.GetTokens())
        } else {
            fmt.Printf("Request %d: Rate limited (tokens: %d)\n", i+1, limiter.GetTokens())
        }
    }
    
    // Sliding window rate limiter
    fmt.Println("\n--- Sliding Window Rate Limiter ---")
    swLimiter := NewSlidingWindowRateLimiter(3, 2*time.Second)
    
    for i := 0; i < 10; i++ {
        if swLimiter.Allow() {
            fmt.Printf("Request %d: Allowed (count: %d)\n", i+1, swLimiter.GetRequestCount())
        } else {
            fmt.Printf("Request %d: Rate limited (count: %d)\n", i+1, swLimiter.GetRequestCount())
        }
        
        time.Sleep(500 * time.Millisecond)
    }
}