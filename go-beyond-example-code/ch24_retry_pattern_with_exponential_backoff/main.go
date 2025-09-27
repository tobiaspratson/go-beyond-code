package main

import (
    "fmt"
    "math"
    "time"
)

// Retry configuration
type RetryConfig struct {
    MaxAttempts    int
    InitialDelay   time.Duration
    MaxDelay       time.Duration
    BackoffFactor  float64
    Jitter         bool
}

func DefaultRetryConfig() *RetryConfig {
    return &RetryConfig{
        MaxAttempts:   3,
        InitialDelay:  100 * time.Millisecond,
        MaxDelay:      5 * time.Second,
        BackoffFactor: 2.0,
        Jitter:        true,
    }
}

// Retry function with exponential backoff
func RetryWithBackoff(config *RetryConfig, fn func() error) error {
    var lastErr error
    
    for attempt := 0; attempt < config.MaxAttempts; attempt++ {
        err := fn()
        if err == nil {
            return nil
        }
        
        lastErr = err
        
        // Don't sleep after the last attempt
        if attempt == config.MaxAttempts-1 {
            break
        }
        
        // Calculate delay with exponential backoff
        delay := time.Duration(float64(config.InitialDelay) * math.Pow(config.BackoffFactor, float64(attempt)))
        
        // Apply jitter to prevent thundering herd
        if config.Jitter {
            jitter := time.Duration(float64(delay) * 0.1 * (0.5 - math.Mod(float64(time.Now().UnixNano()), 1.0)))
            delay += jitter
        }
        
        // Cap the delay at MaxDelay
        if delay > config.MaxDelay {
            delay = config.MaxDelay
        }
        
        fmt.Printf("Attempt %d failed: %v, retrying in %v\n", attempt+1, err, delay)
        time.Sleep(delay)
    }
    
    return fmt.Errorf("all %d attempts failed, last error: %v", config.MaxAttempts, lastErr)
}

// Example service that might fail
type Service struct {
    name     string
    failRate float64
    attempt  int
}

func NewService(name string, failRate float64) *Service {
    return &Service{
        name:     name,
        failRate: failRate,
    }
}

func (s *Service) DoWork() error {
    s.attempt++
    
    // Simulate failure based on fail rate
    if float64(s.attempt) < 1.0/s.failRate {
        return fmt.Errorf("service %s failed on attempt %d", s.name, s.attempt)
    }
    
    fmt.Printf("Service %s succeeded on attempt %d\n", s.name, s.attempt)
    return nil
}

func main() {
    fmt.Println("=== Retry Pattern with Exponential Backoff ===")
    
    // Create a service that fails initially
    service := NewService("DatabaseService", 0.3) // 30% success rate
    
    // Retry configuration
    config := DefaultRetryConfig()
    config.MaxAttempts = 5
    config.InitialDelay = 200 * time.Millisecond
    
    // Execute with retry
    err := RetryWithBackoff(config, service.DoWork)
    if err != nil {
        fmt.Printf("Service failed after all retries: %v\n", err)
    } else {
        fmt.Println("Service succeeded!")
    }
    
    // Test with different configurations
    fmt.Println("\n--- Testing Different Configurations ---")
    
    // Fast retry
    fastConfig := &RetryConfig{
        MaxAttempts:   2,
        InitialDelay:  50 * time.Millisecond,
        MaxDelay:      1 * time.Second,
        BackoffFactor: 1.5,
        Jitter:        false,
    }
    
    service2 := NewService("FastService", 0.5)
    err = RetryWithBackoff(fastConfig, service2.DoWork)
    if err != nil {
        fmt.Printf("Fast service failed: %v\n", err)
    }
    
    // Slow retry
    slowConfig := &RetryConfig{
        MaxAttempts:   4,
        InitialDelay:  1 * time.Second,
        MaxDelay:      10 * time.Second,
        BackoffFactor: 3.0,
        Jitter:        true,
    }
    
    service3 := NewService("SlowService", 0.2)
    err = RetryWithBackoff(slowConfig, service3.DoWork)
    if err != nil {
        fmt.Printf("Slow service failed: %v\n", err)
    }
}