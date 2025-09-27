package main

import (
    "fmt"
    "sync"
    "time"
)

type Resource struct {
    ID   int
    Busy bool
}

type ResourcePool struct {
    resources []*Resource
    mutex     sync.Mutex
    cond      *sync.Cond
}

func NewResourcePool(size int) *ResourcePool {
    rp := &ResourcePool{
        resources: make([]*Resource, size),
    }
    rp.cond = sync.NewCond(&rp.mutex)
    
    // Initialize resources
    for i := 0; i < size; i++ {
        rp.resources[i] = &Resource{ID: i, Busy: false}
    }
    
    return rp
}

func (rp *ResourcePool) Acquire() *Resource {
    rp.mutex.Lock()
    defer rp.mutex.Unlock()
    
    // Wait until a resource is available
    for {
        for _, resource := range rp.resources {
            if !resource.Busy {
                resource.Busy = true
                fmt.Printf("Resource %d acquired\n", resource.ID)
                return resource
            }
        }
        fmt.Println("No resources available, waiting...")
        rp.cond.Wait()
    }
}

func (rp *ResourcePool) Release(resource *Resource) {
    rp.mutex.Lock()
    defer rp.mutex.Unlock()
    
    resource.Busy = false
    fmt.Printf("Resource %d released\n", resource.ID)
    rp.cond.Signal() // Wake up one waiting goroutine
}

func main() {
    pool := NewResourcePool(3)
    var wg sync.WaitGroup
    
    // Multiple workers trying to acquire resources
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            resource := pool.Acquire()
            fmt.Printf("Worker %d using resource %d\n", workerID, resource.ID)
            time.Sleep(200 * time.Millisecond)
            pool.Release(resource)
        }(i)
    }
    
    wg.Wait()
}