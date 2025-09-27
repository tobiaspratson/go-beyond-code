package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeMap struct {
    data map[string]int
    mutex sync.RWMutex
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        data: make(map[string]int),
    }
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mutex.Lock()
    defer sm.mutex.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mutex.RLock()
    defer sm.mutex.RUnlock()
    value, exists := sm.data[key]
    return value, exists
}

func (sm *SafeMap) Size() int {
    sm.mutex.RLock()
    defer sm.mutex.RUnlock()
    return len(sm.data)
}

func (sm *SafeMap) Keys() []string {
    sm.mutex.RLock()
    defer sm.mutex.RUnlock()
    
    keys := make([]string, 0, len(sm.data))
    for key := range sm.data {
        keys = append(keys, key)
    }
    return keys
}

func main() {
    sm := NewSafeMap()
    var wg sync.WaitGroup
    
    // Writers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 5; j++ {
                key := fmt.Sprintf("key-%d-%d", id, j)
                sm.Set(key, id*10+j)
                time.Sleep(10 * time.Millisecond)
            }
        }(i)
    }
    
    // Readers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 3; j++ {
                key := fmt.Sprintf("key-%d-%d", 0, j)
                if value, exists := sm.Get(key); exists {
                    fmt.Printf("Reader %d: %s = %d\n", id, key, value)
                }
                time.Sleep(5 * time.Millisecond)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final map size: %d\n", sm.Size())
}