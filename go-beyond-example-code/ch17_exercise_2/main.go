package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    data map[string]interface{}
    mutex sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
    select {
    case <-ctx.Done():
        return nil, false
    default:
        c.mutex.RLock()
        defer c.mutex.RUnlock()
        value, exists := c.data[key]
        return value, exists
    }
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        c.mutex.Lock()
        defer c.mutex.Unlock()
        c.data[key] = value
        return nil
    }
}

func main() {
    cache := NewCache()
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Set some values
    cache.Set(ctx, "key1", "value1")
    cache.Set(ctx, "key2", "value2")
    
    // Get values
    if value, exists := cache.Get(ctx, "key1"); exists {
        fmt.Printf("key1: %v\n", value)
    }
    
    if value, exists := cache.Get(ctx, "key2"); exists {
        fmt.Printf("key2: %v\n", value)
    }
}