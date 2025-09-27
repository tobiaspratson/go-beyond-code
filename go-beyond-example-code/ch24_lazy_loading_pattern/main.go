package main

import (
    "fmt"
    "sync"
    "time"
)

// Lazy loader for expensive operations
type LazyLoader[T any] struct {
    once   sync.Once
    value  T
    loader func() T
}

func NewLazyLoader[T any](loader func() T) *LazyLoader[T] {
    return &LazyLoader[T]{
        loader: loader,
    }
}

func (ll *LazyLoader[T]) Get() T {
    ll.once.Do(func() {
        ll.value = ll.loader()
    })
    return ll.value
}

// Example: Expensive configuration loader
type Config struct {
    DatabaseURL string
    APIKey      string
    Timeout     time.Duration
}

func LoadConfig() Config {
    // Simulate expensive operation
    time.Sleep(2 * time.Second)
    return Config{
        DatabaseURL: "postgres://localhost:5432/mydb",
        APIKey:      "secret-key",
        Timeout:     30 * time.Second,
    }
}

func main() {
    // Create lazy loader
    configLoader := NewLazyLoader(LoadConfig)
    
    fmt.Println("Starting application...")
    
    // First call will load the config
    start := time.Now()
    config1 := configLoader.Get()
    fmt.Printf("First call took: %v\n", time.Since(start))
    fmt.Printf("Config: %+v\n", config1)
    
    // Second call will use cached value
    start = time.Now()
    config2 := configLoader.Get()
    fmt.Printf("Second call took: %v\n", time.Since(start))
    fmt.Printf("Config: %+v\n", config2)
}