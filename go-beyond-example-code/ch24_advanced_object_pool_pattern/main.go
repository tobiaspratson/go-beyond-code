package main

import (
    "fmt"
    "sync"
    "time"
)

// Enhanced object pool with metrics
type ObjectPool struct {
    pool     sync.Pool
    new      func() interface{}
    reset    func(interface{})
    mu       sync.RWMutex
    created  int64
    reused   int64
    returned int64
}

func NewObjectPool(newFunc func() interface{}, resetFunc func(interface{})) *ObjectPool {
    return &ObjectPool{
        new:   newFunc,
        reset: resetFunc,
        pool: sync.Pool{
            New: func() interface{} {
                obj := newFunc()
                return obj
            },
        },
    }
}

func (p *ObjectPool) Get() interface{} {
    obj := p.pool.Get()
    
    p.mu.Lock()
    if obj == nil {
        p.created++
    } else {
        p.reused++
    }
    p.mu.Unlock()
    
    return obj
}

func (p *ObjectPool) Put(obj interface{}) {
    if obj == nil {
        return
    }
    
    if p.reset != nil {
        p.reset(obj)
    }
    
    p.mu.Lock()
    p.returned++
    p.mu.Unlock()
    
    p.pool.Put(obj)
}

func (p *ObjectPool) Stats() map[string]int64 {
    p.mu.RLock()
    defer p.mu.RUnlock()
    
    return map[string]int64{
        "created":  p.created,
        "reused":   p.reused,
        "returned": p.returned,
    }
}

// Example: Enhanced buffer pool
type Buffer struct {
    data []byte
    size int
    id   int
}

func NewBuffer(size int) *Buffer {
    return &Buffer{
        data: make([]byte, 0, size),
        size: size,
    }
}

func (b *Buffer) Write(data []byte) {
    b.data = append(b.data, data...)
}

func (b *Buffer) String() string {
    return string(b.data)
}

func (b *Buffer) Reset() {
    b.data = b.data[:0]
}

func (b *Buffer) Len() int {
    return len(b.data)
}

func (b *Buffer) Cap() int {
    return cap(b.data)
}

func main() {
    // Create buffer pool
    pool := NewObjectPool(
        func() interface{} {
            return NewBuffer(1024)
        },
        func(obj interface{}) {
            if buf, ok := obj.(*Buffer); ok {
                buf.Reset()
            }
        },
    )
    
    fmt.Println("=== Object Pool Usage ===")
    
    // Use buffers from pool
    for i := 0; i < 5; i++ {
        buf := pool.Get().(*Buffer)
        buf.Write([]byte(fmt.Sprintf("Buffer %d: Hello, World! ", i)))
        fmt.Printf("Buffer content: %s (len: %d, cap: %d)\n", 
            buf.String(), buf.Len(), buf.Cap())
        
        // Return to pool
        pool.Put(buf)
    }
    
    // Show pool statistics
    stats := pool.Stats()
    fmt.Printf("\nPool Statistics: %+v\n", stats)
}