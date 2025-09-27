package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type Resource struct {
    ID       int
    refCount int32
    data     string
}

func (r *Resource) AddRef() int32 {
    return atomic.AddInt32(&r.refCount, 1)
}

func (r *Resource) Release() int32 {
    newCount := atomic.AddInt32(&r.refCount, -1)
    if newCount == 0 {
        fmt.Printf("Resource %d is no longer referenced, cleaning up...\n", r.ID)
        // Cleanup logic here
    }
    return newCount
}

func (r *Resource) RefCount() int32 {
    return atomic.LoadInt32(&r.refCount)
}

func main() {
    resource := &Resource{
        ID:   1,
        data: "Important data",
    }
    
    var wg sync.WaitGroup
    
    // Multiple goroutines using the resource
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            // Add reference
            count := resource.AddRef()
            fmt.Printf("Goroutine %d: ref count = %d\n", id, count)
            
            // Simulate work
            time.Sleep(100 * time.Millisecond)
            
            // Release reference
            count = resource.Release()
            fmt.Printf("Goroutine %d: released, ref count = %d\n", id, count)
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Final ref count: %d\n", resource.RefCount())
}