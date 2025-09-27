package main

import (
    "fmt"
    "sync"
    "time"
)

type ResourceManager struct {
    resources map[string]interface{}
    once      sync.Once
    mu       sync.Mutex
}

func NewResourceManager() *ResourceManager {
    return &ResourceManager{
        resources: make(map[string]interface{}),
    }
}

func (rm *ResourceManager) Initialize() {
    rm.once.Do(func() {
        fmt.Println("Initializing resource manager...")
        time.Sleep(200 * time.Millisecond)
        
        rm.mu.Lock()
        rm.resources["database"] = "connected"
        rm.resources["cache"] = "ready"
        rm.resources["logger"] = "configured"
        rm.mu.Unlock()
        
        fmt.Println("Resource manager initialized")
    })
}

func (rm *ResourceManager) GetResource(name string) (interface{}, bool) {
    rm.Initialize() // Ensure initialization
    
    rm.mu.Lock()
    defer rm.mu.Unlock()
    
    resource, exists := rm.resources[name]
    return resource, exists
}

func main() {
    rm := NewResourceManager()
    var wg sync.WaitGroup
    
    // Multiple goroutines accessing resources
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            if resource, exists := rm.GetResource("database"); exists {
                fmt.Printf("Goroutine %d: Database status: %v\n", id, resource)
            }
        }(i)
    }
    
    wg.Wait()
}