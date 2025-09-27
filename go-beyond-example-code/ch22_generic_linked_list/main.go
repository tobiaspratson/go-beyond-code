package main

import "fmt"

// Generic linked list node
type ListNode[T comparable] struct {
	Value T
	Next  *ListNode[T]
}

// Generic linked list
type LinkedList[T comparable] struct {
	Head *ListNode[T]
	Size int
}

// Constructor for creating a new linked list
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add an item to the front of the list
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &ListNode[T]{
		Value: value,
		Next:  ll.Head,
	}
	ll.Head = newNode
	ll.Size++
}

// Add an item to the end of the list
func (ll *LinkedList[T]) Append(value T) {
	newNode := &ListNode[T]{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

// Remove the first occurrence of a value
func (ll *LinkedList[T]) Remove(value T) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}

	return false
}

// Check if the list contains a value
func (ll *LinkedList[T]) Contains(value T) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Convert the list to a slice
func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.Size)
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Clear all items from the list
func (ll *LinkedList[T]) Clear() {
	ll.Head = nil
	ll.Size = 0
}

func main() {
	// Create a linked list of integers
	intList := NewLinkedList[int]()

	// Add some integers
	intList.Append(1)
	intList.Append(2)
	intList.Append(3)
	intList.Prepend(0)

	fmt.Printf("List size: %d\n", intList.Size)
	fmt.Printf("List contents: %v\n", intList.ToSlice())

	// Check if list contains a value
	fmt.Printf("Contains 2: %t\n", intList.Contains(2))
	fmt.Printf("Contains 5: %t\n", intList.Contains(5))

	// Remove a value
	removed := intList.Remove(2)
	fmt.Printf("Removed 2: %t\n", removed)
	fmt.Printf("List after removal: %v\n", intList.ToSlice())

	// Create a linked list of strings
	stringList := NewLinkedList[string]()
	stringList.Append("first")
	stringList.Append("second")
	stringList.Prepend("zero")

	fmt.Printf("\nString list: %v\n", stringList.ToSlice())
}
