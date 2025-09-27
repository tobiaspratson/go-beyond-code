package main

import (
    "fmt"
    "sort"
)

// Generic priority queue item
type PriorityItem[T any] struct {
    Value    T
    Priority int
}

// Generic priority queue implementation
type PriorityQueue[T any] struct {
    items []PriorityItem[T]
}

// Create a new priority queue
func NewPriorityQueue[T any]() *PriorityQueue[T] {
    return &PriorityQueue[T]{}
}

// Enqueue an item with priority
func (pq *PriorityQueue[T]) Enqueue(item T, priority int) {
    pq.items = append(pq.items, PriorityItem[T]{item, priority})
    // Sort by priority (higher priority first)
    sort.Slice(pq.items, func(i, j int) bool {
        return pq.items[i].Priority > pq.items[j].Priority
    })
}

// Dequeue the highest priority item
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
    if len(pq.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := pq.items[0]
    pq.items = pq.items[1:]
    return item.Value, true
}

// Peek at the highest priority item
func (pq *PriorityQueue[T]) Peek() (T, bool) {
    if len(pq.items) == 0 {
        var zero T
        return zero, false
    }
    return pq.items[0].Value, true
}

// Check if queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
    return len(pq.items) == 0
}

// Get queue size
func (pq *PriorityQueue[T]) Size() int {
    return len(pq.items)
}

// Get all items in the queue (for debugging)
func (pq *PriorityQueue[T]) Items() []PriorityItem[T] {
    return pq.items
}

func main() {
    // Create a priority queue of strings
    pq := NewPriorityQueue[string]()
    
    // Enqueue items with different priorities
    pq.Enqueue("Low priority task", 1)
    pq.Enqueue("High priority task", 10)
    pq.Enqueue("Medium priority task", 5)
    pq.Enqueue("Critical task", 15)
    
    fmt.Printf("Priority queue size: %d\n", pq.Size())
    fmt.Printf("Items: %v\n", pq.Items())
    
    // Dequeue items (should come out in priority order)
    for !pq.IsEmpty() {
        item, ok := pq.Dequeue()
        if ok {
            fmt.Printf("Dequeued: %s\n", item)
        }
    }
    
    // Test with custom types
    type Task struct {
        ID          int
        Description string
    }
    
    taskPQ := NewPriorityQueue[Task]()
    taskPQ.Enqueue(Task{1, "Write documentation"}, 2)
    taskPQ.Enqueue(Task{2, "Fix critical bug"}, 10)
    taskPQ.Enqueue(Task{3, "Code review"}, 5)
    
    fmt.Printf("Task priority queue size: %d\n", taskPQ.Size())
    for !taskPQ.IsEmpty() {
        task, ok := taskPQ.Dequeue()
        if ok {
            fmt.Printf("Dequeued task: %+v\n", task)
        }
    }
}