package main

import "fmt"

// Generic queue implementation
type Queue[T any] struct {
    items []T
}

// Create a new queue
func NewQueue[T any]() *Queue[T] {
    return &Queue[T]{}
}

// Enqueue an item
func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

// Dequeue an item
func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

// Peek at the front item without removing it
func (q *Queue[T]) Peek() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    return q.items[0], true
}

// Check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
    return len(q.items) == 0
}

// Get queue size
func (q *Queue[T]) Size() int {
    return len(q.items)
}

// Get all items in the queue (for debugging)
func (q *Queue[T]) Items() []T {
    return q.items
}

// Clear the queue
func (q *Queue[T]) Clear() {
    q.items = q.items[:0]
}

func main() {
    // Create a queue of integers
    intQueue := NewQueue[int]()
    
    // Enqueue some integers
    intQueue.Enqueue(1)
    intQueue.Enqueue(2)
    intQueue.Enqueue(3)
    
    fmt.Printf("Queue size: %d\n", intQueue.Size())
    fmt.Printf("Queue items: %v\n", intQueue.Items())
    
    // Dequeue items
    for !intQueue.IsEmpty() {
        item, ok := intQueue.Dequeue()
        if ok {
            fmt.Printf("Dequeued: %d\n", item)
        }
    }
    
    // Create a queue of strings
    stringQueue := NewQueue[string]()
    
    // Enqueue some strings
    stringQueue.Enqueue("apple")
    stringQueue.Enqueue("banana")
    stringQueue.Enqueue("cherry")
    
    fmt.Printf("Queue size: %d\n", stringQueue.Size())
    
    // Peek at the front item
    if item, ok := stringQueue.Peek(); ok {
        fmt.Printf("Front item: %s\n", item)
    }
    
    // Dequeue items
    for !stringQueue.IsEmpty() {
        item, ok := stringQueue.Dequeue()
        if ok {
            fmt.Printf("Dequeued: %s\n", item)
        }
    }
    
    // Test with custom types
    type Task struct {
        ID          int
        Description string
        Priority    int
    }
    
    taskQueue := NewQueue[Task]()
    taskQueue.Enqueue(Task{1, "Write code", 1})
    taskQueue.Enqueue(Task{2, "Test code", 2})
    taskQueue.Enqueue(Task{3, "Deploy code", 3})
    
    fmt.Printf("Task queue size: %d\n", taskQueue.Size())
    for !taskQueue.IsEmpty() {
        task, ok := taskQueue.Dequeue()
        if ok {
            fmt.Printf("Dequeued task: %+v\n", task)
        }
    }
}