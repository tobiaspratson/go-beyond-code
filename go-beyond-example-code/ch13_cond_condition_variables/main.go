package main

import (
    "fmt"
    "sync"
    "time"
)

type Queue struct {
    items []int
    mutex sync.Mutex
    cond  *sync.Cond
}

func NewQueue() *Queue {
    q := &Queue{}
    q.cond = sync.NewCond(&q.mutex)
    return q
}

func (q *Queue) Enqueue(item int) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    
    q.items = append(q.items, item)
    fmt.Printf("Enqueued: %d\n", item)
    q.cond.Signal() // Wake up one waiting goroutine
}

func (q *Queue) Dequeue() int {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    
    // Wait until queue has items
    for len(q.items) == 0 {
        fmt.Println("Queue empty, waiting...")
        q.cond.Wait()
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    fmt.Printf("Dequeued: %d\n", item)
    return item
}

func main() {
    queue := NewQueue()
    var wg sync.WaitGroup
    
    // Start consumers
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 3; j++ {
                item := queue.Dequeue()
                fmt.Printf("Consumer %d processed: %d\n", id, item)
            }
        }(i)
    }
    
    // Start producer
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 1; i <= 9; i++ {
            queue.Enqueue(i)
            time.Sleep(100 * time.Millisecond)
        }
    }()
    
    wg.Wait()
}