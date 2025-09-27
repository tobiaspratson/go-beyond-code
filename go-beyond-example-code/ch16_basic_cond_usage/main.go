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
	q := &Queue{
		items: make([]int, 0),
	}
	q.cond = sync.NewCond(&q.mutex)
	return q
}

func (q *Queue) Enqueue(item int) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
	fmt.Printf("Enqueued: %d, queue length: %d\n", item, len(q.items))
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
	fmt.Printf("Dequeued: %d, queue length: %d\n", item, len(q.items))
	return item
}

func main() {
	queue := NewQueue()
	var wg sync.WaitGroup

	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			item := queue.Dequeue()
			fmt.Printf("Consumer: %d\n", item)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			queue.Enqueue(i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()
}
