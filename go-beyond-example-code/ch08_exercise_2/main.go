package main

import (
	"errors"
	"fmt"
	"sync"
)

// Custom error types for stack operations
type StackError struct {
	Operation string
	Message   string
	Size      int
}

func (e StackError) Error() string {
	return fmt.Sprintf("stack error in %s: %s (stack size: %d)",
		e.Operation, e.Message, e.Size)
}

var (
	ErrStackEmpty  = errors.New("stack is empty")
	ErrStackFull   = errors.New("stack is full")
	ErrInvalidItem = errors.New("invalid item")
)

// Generic stack with error handling
type Stack[T any] struct {
	items   []T
	maxSize int
	mutex   sync.RWMutex
}

func NewStack[T any](maxSize int) *Stack[T] {
	return &Stack[T]{
		items:   make([]T, 0),
		maxSize: maxSize,
	}
}

func (s *Stack[T]) Push(item T) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.maxSize > 0 && len(s.items) >= s.maxSize {
		return StackError{
			Operation: "PUSH",
			Message:   "stack is full",
			Size:      len(s.items),
		}
	}

	s.items = append(s.items, item)
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var zero T
	if len(s.items) == 0 {
		return zero, StackError{
			Operation: "POP",
			Message:   "stack is empty",
			Size:      len(s.items),
		}
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var zero T
	if len(s.items) == 0 {
		return zero, StackError{
			Operation: "PEEK",
			Message:   "stack is empty",
			Size:      len(s.items),
		}
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.items) == 0
}

func (s *Stack[T]) IsFull() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.maxSize > 0 && len(s.items) >= s.maxSize
}

func (s *Stack[T]) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = s.items[:0]
}

func (s *Stack[T]) ToSlice() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// Stack operations with error handling
func (s *Stack[T]) PushMultiple(items []T) error {
	for _, item := range items {
		if err := s.Push(item); err != nil {
			return fmt.Errorf("failed to push item %v: %w", item, err)
		}
	}
	return nil
}

func (s *Stack[T]) PopMultiple(count int) ([]T, error) {
	if count <= 0 {
		return nil, StackError{
			Operation: "POP_MULTIPLE",
			Message:   "count must be positive",
			Size:      s.Size(),
		}
	}

	if count > s.Size() {
		return nil, StackError{
			Operation: "POP_MULTIPLE",
			Message:   "not enough items in stack",
			Size:      s.Size(),
		}
	}

	result := make([]T, count)
	for i := 0; i < count; i++ {
		item, err := s.Pop()
		if err != nil {
			return nil, fmt.Errorf("failed to pop item %d: %w", i, err)
		}
		result[count-1-i] = item // Reverse order
	}

	return result, nil
}

// Stack with panic recovery
func (s *Stack[T]) SafePop() (item T, recovered bool) {
	defer func() {
		if r := recover(); r != nil {
			recovered = true
			fmt.Printf("Recovered from panic in SafePop: %v\n", r)
		}
	}()

	item, err := s.Pop()
	if err != nil {
		panic(fmt.Sprintf("Pop failed: %v", err))
	}

	return item, false
}

func main() {
	// Test integer stack
	fmt.Println("=== Integer Stack Tests ===")
	intStack := NewStack[int](5)

	// Push items
	for i := 1; i <= 3; i++ {
		err := intStack.Push(i * 10)
		if err != nil {
			fmt.Printf("Push error: %v\n", err)
		} else {
			fmt.Printf("Pushed: %d\n", i*10)
		}
	}

	// Peek at top
	top, err := intStack.Peek()
	if err != nil {
		fmt.Printf("Peek error: %v\n", err)
	} else {
		fmt.Printf("Top item: %d\n", top)
	}

	// Pop items
	for !intStack.IsEmpty() {
		item, err := intStack.Pop()
		if err != nil {
			fmt.Printf("Pop error: %v\n", err)
			break
		}
		fmt.Printf("Popped: %d\n", item)
	}

	// Test string stack
	fmt.Println("\n=== String Stack Tests ===")
	strStack := NewStack[string](3)

	words := []string{"hello", "world", "golang"}
	err = strStack.PushMultiple(words)
	if err != nil {
		fmt.Printf("PushMultiple error: %v\n", err)
	} else {
		fmt.Printf("Pushed %d items\n", len(words))
	}

	// Pop multiple items
	items, err := strStack.PopMultiple(2)
	if err != nil {
		fmt.Printf("PopMultiple error: %v\n", err)
	} else {
		fmt.Printf("Popped items: %v\n", items)
	}

	// Test error conditions
	fmt.Println("\n=== Error Tests ===")

	// Try to pop from empty stack
	_, err = intStack.Pop()
	if err != nil {
		fmt.Printf("Pop from empty stack: %v\n", err)
	}

	// Try to peek at empty stack
	_, err = intStack.Peek()
	if err != nil {
		fmt.Printf("Peek at empty stack: %v\n", err)
	}

	// Fill stack to capacity
	smallStack := NewStack[int](2)
	for i := 1; i <= 3; i++ {
		err := smallStack.Push(i)
		if err != nil {
			fmt.Printf("Push to full stack: %v\n", err)
			break
		}
		fmt.Printf("Pushed to small stack: %d\n", i)
	}

	// Test safe pop with panic recovery
	fmt.Println("\n=== Safe Pop Tests ===")
	safeStack := NewStack[string](0)
	item, recovered := safeStack.SafePop()
	if recovered {
		fmt.Printf("SafePop recovered from panic\n")
	} else {
		fmt.Printf("SafePop result: %s\n", item)
	}
}
