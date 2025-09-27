package main

import (
    "fmt"
    "sync"
    "time"
)

type Barrier struct {
    count    int
    total    int
    mutex    sync.Mutex
    cond     *sync.Cond
    phase    int
}

func NewBarrier(total int) *Barrier {
    b := &Barrier{total: total}
    b.cond = sync.NewCond(&b.mutex)
    return b
}

func (b *Barrier) Wait() {
    b.mutex.Lock()
    defer b.mutex.Unlock()
    
    b.count++
    currentPhase := b.phase
    
    if b.count == b.total {
        // Last goroutine to arrive
        b.count = 0
        b.phase++
        b.cond.Broadcast() // Wake up all waiting goroutines
    } else {
        // Wait for all goroutines to arrive
        for currentPhase == b.phase {
            b.cond.Wait()
        }
    }
}

func main() {
    barrier := NewBarrier(3)
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            fmt.Printf("Goroutine %d starting phase 1\n", id)
            time.Sleep(time.Duration(id) * 100 * time.Millisecond)
            barrier.Wait()
            
            fmt.Printf("Goroutine %d starting phase 2\n", id)
            time.Sleep(time.Duration(id) * 100 * time.Millisecond)
            barrier.Wait()
            
            fmt.Printf("Goroutine %d finished\n", id)
        }(i)
    }
    
    wg.Wait()
}