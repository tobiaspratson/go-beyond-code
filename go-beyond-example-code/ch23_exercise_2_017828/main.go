package main

import (
    "fmt"
    "time"
)

// Generic cache item
type CacheItem[T any] struct {
    Value     T
    ExpiresAt time.Time
}

// Generic cache
type Cache[K comparable, V any] struct {
    items map[K]CacheItem[V]
}

// Create a new cache
func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        items: make(map[K]CacheItem[V]),
    }
}

// Set a value with TTL
func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
    c.items[key] = CacheItem[V]{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

// Get a value
func (c *Cache[K, V]) Get(key K) (V, bool) {
    item, exists := c.items[key]
    if !exists {
        var zero V
        return zero, false
    }
    
    if time.Now().After(item.ExpiresAt) {
        delete(c.items, key)
        var zero V
        return zero, false
    }
    
    return item.Value, true
}

// Clear expired items
func (c *Cache[K, V]) ClearExpired() {
    now := time.Now()
    for key, item := range c.items {
        if now.After(item.ExpiresAt) {
            delete(c.items, key)
        }
    }
}

// Get cache size
func (c *Cache[K, V]) Size() int {
    return len(c.items)
}

func main() {
    // Create a cache with string keys and int values
    cache := NewCache[string, int]()
    
    // Set some values with TTL
    cache.Set("apple", 5, 1*time.Second)
    cache.Set("banana", 3, 2*time.Second)
    cache.Set("cherry", 8, 3*time.Second)
    
    fmt.Printf("Cache size: %d\n", cache.Size())
    
    // Get values
    if value, exists := cache.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }
    
    if value, exists := cache.Get("banana"); exists {
        fmt.Printf("banana: %d\n", value)
    }
    
    // Wait for some items to expire
    time.Sleep(1500 * time.Millisecond)
    
    // Clear expired items
    cache.ClearExpired()
    fmt.Printf("Cache size after cleanup: %d\n", cache.Size())
    
    // Try to get expired item
    if value, exists := cache.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    } else {
        fmt.Println("apple expired")
    }
}