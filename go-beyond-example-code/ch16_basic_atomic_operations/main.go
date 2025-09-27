package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&ac.value)
}

func (ac *AtomicCounter) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&ac.value, old, new)
}

func main() {
	counter := &AtomicCounter{}
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
}
