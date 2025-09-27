package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

type MutexCounter struct {
    value int64
    mutex sync.Mutex
}

func (mc *MutexCounter) Increment() {
    mc.mutex.Lock()
    defer mc.mutex.Unlock()
    mc.value++
}

func (mc *MutexCounter) Value() int64 {
    mc.mutex.Lock()
    defer mc.mutex.Unlock()
    return mc.value
}

type AtomicCounter struct {
    value int64
}

func (ac *AtomicCounter) Increment() {
    atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&ac.value)
}

func benchmarkMutex(counter *MutexCounter, numOps int) time.Duration {
    start := time.Now()
    var wg sync.WaitGroup
    
    for i := 0; i < numOps; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    return time.Since(start)
}

func benchmarkAtomic(counter *AtomicCounter, numOps int) time.Duration {
    start := time.Now()
    var wg sync.WaitGroup
    
    for i := 0; i < numOps; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    return time.Since(start)
}

func main() {
    numOps := 10000
    
    mutexCounter := &MutexCounter{}
    atomicCounter := &AtomicCounter{}
    
    mutexTime := benchmarkMutex(mutexCounter, numOps)
    atomicTime := benchmarkAtomic(atomicCounter, numOps)
    
    fmt.Printf("Mutex time: %v\n", mutexTime)
    fmt.Printf("Atomic time: %v\n", atomicTime)
    fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}