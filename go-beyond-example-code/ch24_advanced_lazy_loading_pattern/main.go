package main

import (
    "fmt"
    "sync"
    "time"
)

// Generic lazy loader with caching and expiration
type LazyLoader[T any] struct {
    once     sync.Once
    value    T
    loader   func() T
    mu       sync.RWMutex
    loadedAt time.Time
    ttl      time.Duration
}

func NewLazyLoader[T any](loader func() T) *LazyLoader[T] {
    return &LazyLoader[T]{
        loader: loader,
        ttl:    -1, // No expiration by default
    }
}

func NewLazyLoaderWithTTL[T any](loader func() T, ttl time.Duration) *LazyLoader[T] {
    return &LazyLoader[T]{
        loader: loader,
        ttl:    ttl,
    }
}

func (ll *LazyLoader[T]) Get() T {
    ll.mu.RLock()
    if ll.isExpired() {
        ll.mu.RUnlock()
        ll.mu.Lock()
        defer ll.mu.Unlock()
        
        if ll.isExpired() {
            ll.once = sync.Once{} // Reset once to allow reloading
        }
    } else {
        defer ll.mu.RUnlock()
    }
    
    ll.once.Do(func() {
        ll.value = ll.loader()
        ll.loadedAt = time.Now()
    })
    
    return ll.value
}

func (ll *LazyLoader[T]) isExpired() bool {
    if ll.ttl < 0 {
        return false
    }
    return time.Since(ll.loadedAt) > ll.ttl
}

func (ll *LazyLoader[T]) IsLoaded() bool {
    ll.mu.RLock()
    defer ll.mu.RUnlock()
    return !ll.loadedAt.IsZero()
}

func (ll *LazyLoader[T]) ForceReload() T {
    ll.mu.Lock()
    defer ll.mu.Unlock()
    
    ll.once = sync.Once{}
    ll.value = ll.loader()
    ll.loadedAt = time.Now()
    
    return ll.value
}

// Example: Expensive configuration loader
type Config struct {
    DatabaseURL string
    APIKey      string
    Timeout     time.Duration
    CacheSize   int
}

func LoadConfig() Config {
    // Simulate expensive operation
    time.Sleep(2 * time.Second)
    fmt.Println("Loading configuration...")
    
    return Config{
        DatabaseURL: "postgres://localhost:5432/mydb",
        APIKey:      "secret-key-12345",
        Timeout:     30 * time.Second,
        CacheSize:   1000,
    }
}

func main() {
    fmt.Println("=== Lazy Loading Example ===")
    
    // Create lazy loader with TTL
    configLoader := NewLazyLoaderWithTTL(LoadConfig, 5*time.Second)
    
    fmt.Println("Starting application...")
    fmt.Printf("Config loaded: %t\n", configLoader.IsLoaded())
    
    // First call will load the config
    start := time.Now()
    config1 := configLoader.Get()
    fmt.Printf("First call took: %v\n", time.Since(start))
    fmt.Printf("Config: %+v\n", config1)
    fmt.Printf("Config loaded: %t\n", configLoader.IsLoaded())
    
    // Second call will use cached value
    start = time.Now()
    config2 := configLoader.Get()
    fmt.Printf("Second call took: %v\n", time.Since(start))
    fmt.Printf("Config: %+v\n", config2)
    
    // Wait for expiration
    fmt.Println("Waiting for config to expire...")
    time.Sleep(6 * time.Second)
    
    // This call will reload the config
    start = time.Now()
    config3 := configLoader.Get()
    fmt.Printf("Expired call took: %v\n", time.Since(start))
    fmt.Printf("Config: %+v\n", config3)
}