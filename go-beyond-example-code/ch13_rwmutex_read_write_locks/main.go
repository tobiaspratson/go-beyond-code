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
    sm.mutex.Lock()         // Exclusive lock for writing
    defer sm.mutex.Unlock()
    sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mutex.RLock()        // Shared lock for reading
    defer sm.mutex.RUnlock()
    value, exists := sm.data[key]
    return value, exists
}

func (sm *SafeMap) Size() int {
    sm.mutex.RLock()
    defer sm.mutex.RUnlock()
    return len(sm.data)
}

func writer(sm *SafeMap, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 100; i++ {
        key := fmt.Sprintf("key%d", i)
        sm.Set(key, i)
    }
}

func reader(sm *SafeMap, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 100; i++ {
        key := fmt.Sprintf("key%d", i)
        value, exists := sm.Get(key)
        if exists {
            fmt.Printf("Read %s: %d\n", key, value)
        }
    }
}

func main() {
    sm := NewSafeMap()
    var wg sync.WaitGroup
    
    // Start writers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go writer(sm, &wg)
    }
    
    // Start readers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go reader(sm, &wg)
    }
    
    wg.Wait()
    fmt.Printf("Final map size: %d\n", sm.Size())
}