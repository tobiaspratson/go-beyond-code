package main

import (
    "fmt"
    "time"
)

// Generic cache item that stores a value and its expiration time
type CacheItem[T any] struct {
    Value     T
    ExpiresAt time.Time
}

// Generic cache with key type K and value type V
type Cache[K comparable, V any] struct {
    items map[K]CacheItem[V]
}

// Create a new cache
func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        items: make(map[K]CacheItem[V]),
    }
}

// Set a value with TTL (time-to-live)
func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
    c.items[key] = CacheItem[V]{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

// Get a value (returns false if expired or not found)
func (c *Cache[K, V]) Get(key K) (V, bool) {
    item, exists := c.items[key]
    if !exists {
        var zero V
        return zero, false
    }
    
    // Check if the item has expired
    if time.Now().After(item.ExpiresAt) {
        delete(c.items, key)
        var zero V
        return zero, false
    }
    
    return item.Value, true
}

// Delete a key
func (c *Cache[K, V]) Delete(key K) {
    delete(c.items, key)
}

// Clear all expired items
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

// Check if a key exists (and is not expired)
func (c *Cache[K, V]) Has(key K) bool {
    _, exists := c.Get(key)
    return exists
}

// Get all keys (including expired ones)
func (c *Cache[K, V]) Keys() []K {
    keys := make([]K, 0, len(c.items))
    for key := range c.items {
        keys = append(keys, key)
    }
    return keys
}

// Get all non-expired values
func (c *Cache[K, V]) Values() []V {
    var values []V
    for key := range c.items {
        if value, exists := c.Get(key); exists {
            values = append(values, value)
        }
    }
    return values
}

// Clear all items
func (c *Cache[K, V]) Clear() {
    c.items = make(map[K]CacheItem[V])
}

// Update a value if it exists and is not expired
func (c *Cache[K, V]) Update(key K, updater func(V) V) bool {
    if value, exists := c.Get(key); exists {
        c.Set(key, updater(value), time.Hour) // Reset TTL to 1 hour
        return true
    }
    return false
}

// Get or set a default value
func (c *Cache[K, V]) GetOrSet(key K, defaultValue V, ttl time.Duration) V {
    if value, exists := c.Get(key); exists {
        return value
    }
    c.Set(key, defaultValue, ttl)
    return defaultValue
}

func main() {
    // Create a cache with string keys and int values
    cache := NewCache[string, int]()
    
    // Set some values with different TTLs
    cache.Set("apple", 5, 1*time.Second)
    cache.Set("banana", 3, 2*time.Second)
    cache.Set("cherry", 8, 3*time.Second)
    cache.Set("date", 12, 5*time.Second)
    
    fmt.Printf("Cache size: %d\n", cache.Size())
    fmt.Printf("Cache keys: %v\n", cache.Keys())
    
    // Get values
    if value, exists := cache.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }
    
    if value, exists := cache.Get("banana"); exists {
        fmt.Printf("banana: %d\n", value)
    }
    
    // Check if keys exist
    fmt.Printf("Has 'apple': %t\n", cache.Has("apple"))
    fmt.Printf("Has 'grape': %t\n", cache.Has("grape"))
    
    // Wait for some items to expire
    fmt.Println("\nWaiting for some items to expire...")
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
    
    // Try to get non-expired item
    if value, exists := cache.Get("banana"); exists {
        fmt.Printf("banana: %d\n", value)
    } else {
        fmt.Println("banana expired")
    }
    
    // Test Update method
    updated := cache.Update("cherry", func(v int) int {
        return v * 2
    })
    fmt.Printf("Updated cherry: %t\n", updated)
    if value, exists := cache.Get("cherry"); exists {
        fmt.Printf("cherry after update: %d\n", value)
    }
    
    // Test GetOrSet method
    defaultValue := cache.GetOrSet("grape", 10, 1*time.Second)
    fmt.Printf("grape default value: %d\n", defaultValue)
    
    // Create a cache with different types
    stringCache := NewCache[int, string]()
    stringCache.Set(1, "one", 2*time.Second)
    stringCache.Set(2, "two", 3*time.Second)
    
    fmt.Printf("\nString cache size: %d\n", stringCache.Size())
    if value, exists := stringCache.Get(1); exists {
        fmt.Printf("Key 1: %s\n", value)
    }
    
    // Create a cache with custom struct values
    type Person struct {
        Name string
        Age  int
    }
    
    personCache := NewCache[string, Person]()
    personCache.Set("alice", Person{Name: "Alice", Age: 30}, 5*time.Second)
    personCache.Set("bob", Person{Name: "Bob", Age: 25}, 3*time.Second)
    
    fmt.Printf("\nPerson cache size: %d\n", personCache.Size())
    if person, exists := personCache.Get("alice"); exists {
        fmt.Printf("Alice: %+v\n", person)
    }
}