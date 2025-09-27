package main

import (
    "fmt"
    "sync"
    "time"
)

type Data struct {
    value int
    mutex sync.Mutex
    rwMutex sync.RWMutex
}

func (d *Data) ReadWithMutex() int {
    d.mutex.Lock()
    defer d.mutex.Unlock()
    return d.value
}

func (d *Data) ReadWithRWMutex() int {
    d.rwMutex.RLock()
    defer d.rwMutex.RUnlock()
    return d.value
}

func (d *Data) WriteWithMutex(value int) {
    d.mutex.Lock()
    defer d.mutex.Unlock()
    d.value = value
}

func (d *Data) WriteWithRWMutex(value int) {
    d.rwMutex.Lock()
    defer d.rwMutex.Unlock()
    d.value = value
}

func benchmarkReads(data *Data, useRWMutex bool, numReads int) time.Duration {
    start := time.Now()
    var wg sync.WaitGroup
    
    for i := 0; i < numReads; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            if useRWMutex {
                data.ReadWithRWMutex()
            } else {
                data.ReadWithMutex()
            }
        }()
    }
    
    wg.Wait()
    return time.Since(start)
}

func main() {
    data := &Data{value: 42}
    numReads := 1000
    
    // Benchmark with regular mutex
    mutexTime := benchmarkReads(data, false, numReads)
    fmt.Printf("Mutex reads: %v\n", mutexTime)
    
    // Benchmark with RWMutex
    rwMutexTime := benchmarkReads(data, true, numReads)
    fmt.Printf("RWMutex reads: %v\n", rwMutexTime)
    
    if rwMutexTime < mutexTime {
        fmt.Printf("RWMutex is %.2fx faster for reads\n", 
            float64(mutexTime)/float64(rwMutexTime))
    }
}