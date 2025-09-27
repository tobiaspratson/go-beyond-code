package main

import "fmt"

// Generic stack implementation
type Stack[T any] struct {
    items []T
}

// Create a new stack
func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}

// Push an item onto the stack
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Pop an item from the stack
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Peek at the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

// Check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Get stack size
func (s *Stack[T]) Size() int {
    return len(s.items)
}

// Get all items in the stack (for debugging)
func (s *Stack[T]) Items() []T {
    return s.items
}

// Clear the stack
func (s *Stack[T]) Clear() {
    s.items = s.items[:0]
}

func main() {
    // Create a stack of integers
    intStack := NewStack[int]()
    
    // Push some integers
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    
    fmt.Printf("Stack size: %d\n", intStack.Size())
    fmt.Printf("Stack items: %v\n", intStack.Items())
    
    // Pop items
    for !intStack.IsEmpty() {
        item, ok := intStack.Pop()
        if ok {
            fmt.Printf("Popped: %d\n", item)
        }
    }
    
    // Create a stack of strings
    stringStack := NewStack[string]()
    
    // Push some strings
    stringStack.Push("apple")
    stringStack.Push("banana")
    stringStack.Push("cherry")
    
    fmt.Printf("Stack size: %d\n", stringStack.Size())
    
    // Peek at the top item
    if item, ok := stringStack.Peek(); ok {
        fmt.Printf("Top item: %s\n", item)
    }
    
    // Pop items
    for !stringStack.IsEmpty() {
        item, ok := stringStack.Pop()
        if ok {
            fmt.Printf("Popped: %s\n", item)
        }
    }
    
    // Test with custom types
    type Person struct {
        Name string
        Age  int
    }
    
    personStack := NewStack[Person]()
    personStack.Push(Person{"Alice", 25})
    personStack.Push(Person{"Bob", 30})
    
    fmt.Printf("Person stack size: %d\n", personStack.Size())
    for !personStack.IsEmpty() {
        person, ok := personStack.Pop()
        if ok {
            fmt.Printf("Popped person: %+v\n", person)
        }
    }
}