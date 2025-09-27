package main

import "fmt"

// Generic stack implementation
// T is the type parameter, any is the constraint (most permissive)
type Stack[T any] struct {
    items []T
}

// Constructor function for creating a new stack
func NewStack[T any]() *Stack[T] {
    return &Stack[T]{
        items: make([]T, 0),
    }
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

// Check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Get stack size
func (s *Stack[T]) Size() int {
    return len(s.items)
}

// Peek at the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

// Clear all items from the stack
func (s *Stack[T]) Clear() {
    s.items = s.items[:0]
}

// Get all items as a slice (for debugging/inspection)
func (s *Stack[T]) ToSlice() []T {
    result := make([]T, len(s.items))
    copy(result, s.items)
    return result
}

func main() {
    // Create a stack of integers
    intStack := NewStack[int]()
    
    // Push some integers
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
    
    fmt.Printf("Stack size: %d\n", intStack.Size())
    fmt.Printf("Stack contents: %v\n", intStack.ToSlice())
    
    // Peek at the top item
    if item, ok := intStack.Peek(); ok {
        fmt.Printf("Top item: %d\n", item)
    }
    
    // Pop items
    fmt.Println("Popping items:")
    for !intStack.IsEmpty() {
        item, ok := intStack.Pop()
        if ok {
            fmt.Printf("Popped: %d\n", item)
        }
    }
    
    fmt.Printf("Stack size after popping: %d\n", intStack.Size())
    
    // Create a stack of strings
    stringStack := NewStack[string]()
    
    // Push some strings
    stringStack.Push("apple")
    stringStack.Push("banana")
    stringStack.Push("cherry")
    
    fmt.Printf("\nString stack size: %d\n", stringStack.Size())
    fmt.Printf("String stack contents: %v\n", stringStack.ToSlice())
    
    // Peek at the top item
    if item, ok := stringStack.Peek(); ok {
        fmt.Printf("Top item: %s\n", item)
    }
    
    // Pop items
    fmt.Println("Popping string items:")
    for !stringStack.IsEmpty() {
        item, ok := stringStack.Pop()
        if ok {
            fmt.Printf("Popped: %s\n", item)
        }
    }
    
    // Create a stack of custom structs
    type Person struct {
        Name string
        Age  int
    }
    
    personStack := NewStack[Person]()
    personStack.Push(Person{Name: "Alice", Age: 30})
    personStack.Push(Person{Name: "Bob", Age: 25})
    
    fmt.Printf("\nPerson stack size: %d\n", personStack.Size())
    
    if person, ok := personStack.Peek(); ok {
        fmt.Printf("Top person: %+v\n", person)
    }
}