package main

import (
    "context"
    "fmt"
    "time"
)

type RateLimiter struct {
    tokens   chan struct{}
    interval time.Duration
}

func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
    rl := &RateLimiter{
        tokens:   make(chan struct{}, rate),
        interval: interval,
    }
    
    // Fill initial tokens
    for i := 0; i < rate; i++ {
        rl.tokens <- struct{}{}
    }
    
    // Start token refill
    go rl.refill()
    
    return rl
}

func (rl *RateLimiter) refill() {
    ticker := time.NewTicker(rl.interval)
    defer ticker.Stop()
    
    for range ticker.C {
        select {
        case rl.tokens <- struct{}{}:
        default:
            // Token bucket is full
        }
    }
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
    select {
    case <-rl.tokens:
        return nil
    case <-ctx.Done():
        return fmt.Errorf("rate limiter wait cancelled: %v", ctx.Err())
    }
}

func main() {
    // Create rate limiter: 5 requests per second
    limiter := NewRateLimiter(5, time.Second)
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // Make requests with rate limiting
    for i := 1; i <= 10; i++ {
        if err := limiter.Wait(ctx); err != nil {
            fmt.Printf("Rate limiter error: %v\n", err)
            break
        }
        
        fmt.Printf("Request %d processed\n", i)
        time.Sleep(100 * time.Millisecond)
    }
}