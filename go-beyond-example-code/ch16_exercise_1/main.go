package main

import (
    "fmt"
    "sync"
    "time"
)

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

type SafeCache struct {
    items map[string]CacheItem
    mutex sync.RWMutex
}

func NewSafeCache() *SafeCache {
    return &SafeCache{
        items: make(map[string]CacheItem),
    }
}

func (c *SafeCache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.items[key] = CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}

func (c *SafeCache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    item, exists := c.items[key]
    if !exists {
        return nil, false
    }
    
    if time.Now().After(item.ExpiresAt) {
        return nil, false
    }
    
    return item.Value, true
}

func (c *SafeCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    delete(c.items, key)
}

func main() {
    cache := NewSafeCache()
    var wg sync.WaitGroup
    
    // Set some values
    cache.Set("key1", "value1", 1*time.Second)
    cache.Set("key2", "value2", 2*time.Second)
    
    // Multiple readers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 3; j++ {
                if value, exists := cache.Get("key1"); exists {
                    fmt.Printf("Reader %d: key1 = %v\n", id, value)
                } else {
                    fmt.Printf("Reader %d: key1 not found\n", id)
                }
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    wg.Wait()
}