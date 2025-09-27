package main

import "fmt"

// Generic queue implementation using a slice
type Queue[T any] struct {
    items []T
}

// Constructor for creating a new queue
func NewQueue[T any]() *Queue[T] {
    return &Queue[T]{
        items: make([]T, 0),
    }
}

// Enqueue adds an item to the back of the queue
func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

// Dequeue removes and returns an item from the front of the queue
func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

// Peek returns the front item without removing it
func (q *Queue[T]) Peek() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    return q.items[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
    return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
    return len(q.items)
}

// Clear removes all items from the queue
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
    
    // Peek at the front item
    if item, ok := intQueue.Peek(); ok {
        fmt.Printf("Front item: %d\n", item)
    }
    
    // Dequeue items
    fmt.Println("Dequeuing items:")
    for !intQueue.IsEmpty() {
        item, ok := intQueue.Dequeue()
        if ok {
            fmt.Printf("Dequeued: %d\n", item)
        }
    }
    
    // Create a queue of strings
    stringQueue := NewQueue[string]()
    stringQueue.Enqueue("first")
    stringQueue.Enqueue("second")
    stringQueue.Enqueue("third")
    
    fmt.Printf("\nString queue size: %d\n", stringQueue.Size())
    
    // Dequeue string items
    fmt.Println("Dequeuing string items:")
    for !stringQueue.IsEmpty() {
        item, ok := stringQueue.Dequeue()
        if ok {
            fmt.Printf("Dequeued: %s\n", item)
        }
    }
}