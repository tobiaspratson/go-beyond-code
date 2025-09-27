package main

import (
    "errors"
    "fmt"
)

// Node represents a single node in the linked list
type Node[T any] struct {
    Data T
    Next *Node[T]
}

// LinkedList represents a linked list
type LinkedList[T any] struct {
    Head *Node[T]
    Size int
}

// Custom error types
type LinkedListError struct {
    Operation string
    Message   string
    Index     int
}

func (e LinkedListError) Error() string {
    return fmt.Sprintf("linked list error in %s: %s (index: %d)", 
        e.Operation, e.Message, e.Index)
}

var (
    ErrIndexOutOfBounds = errors.New("index out of bounds")
    ErrEmptyList        = errors.New("list is empty")
    ErrInvalidIndex     = errors.New("invalid index")
)

// Create a new linked list
func NewLinkedList[T any]() *LinkedList[T] {
    return &LinkedList[T]{}
}

// Add element to the beginning
func (ll *LinkedList[T]) Prepend(data T) {
    newNode := &Node[T]{
        Data: data,
        Next: ll.Head,
    }
    ll.Head = newNode
    ll.Size++
}

// Add element to the end
func (ll *LinkedList[T]) Append(data T) {
    newNode := &Node[T]{Data: data}
    
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

// Insert element at specific index
func (ll *LinkedList[T]) InsertAt(index int, data T) error {
    if index < 0 || index > ll.Size {
        return LinkedListError{
            Operation: "INSERT_AT",
            Message:   "index out of bounds",
            Index:     index,
        }
    }
    
    if index == 0 {
        ll.Prepend(data)
        return nil
    }
    
    if index == ll.Size {
        ll.Append(data)
        return nil
    }
    
    newNode := &Node[T]{Data: data}
    current := ll.Head
    
    for i := 0; i < index-1; i++ {
        current = current.Next
    }
    
    newNode.Next = current.Next
    current.Next = newNode
    ll.Size++
    
    return nil
}

// Remove element at specific index
func (ll *LinkedList[T]) RemoveAt(index int) (T, error) {
    var zero T
    
    if ll.Size == 0 {
        return zero, LinkedListError{
            Operation: "REMOVE_AT",
            Message:   "list is empty",
            Index:     index,
        }
    }
    
    if index < 0 || index >= ll.Size {
        return zero, LinkedListError{
            Operation: "REMOVE_AT",
            Message:   "index out of bounds",
            Index:     index,
        }
    }
    
    if index == 0 {
        data := ll.Head.Data
        ll.Head = ll.Head.Next
        ll.Size--
        return data, nil
    }
    
    current := ll.Head
    for i := 0; i < index-1; i++ {
        current = current.Next
    }
    
    data := current.Next.Data
    current.Next = current.Next.Next
    ll.Size--
    
    return data, nil
}

// Get element at specific index
func (ll *LinkedList[T]) GetAt(index int) (T, error) {
    var zero T
    
    if ll.Size == 0 {
        return zero, LinkedListError{
            Operation: "GET_AT",
            Message:   "list is empty",
            Index:     index,
        }
    }
    
    if index < 0 || index >= ll.Size {
        return zero, LinkedListError{
            Operation: "GET_AT",
            Message:   "index out of bounds",
            Index:     index,
        }
    }
    
    current := ll.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    
    return current.Data, nil
}

// Check if list is empty
func (ll *LinkedList[T]) IsEmpty() bool {
    return ll.Size == 0
}

// Get list size
func (ll *LinkedList[T]) GetSize() int {
    return ll.Size
}

// Convert to slice
func (ll *LinkedList[T]) ToSlice() []T {
    result := make([]T, 0, ll.Size)
    current := ll.Head
    
    for current != nil {
        result = append(result, current.Data)
        current = current.Next
    }
    
    return result
}

// Clear the list
func (ll *LinkedList[T]) Clear() {
    ll.Head = nil
    ll.Size = 0
}

// Find element and return its index
func (ll *LinkedList[T]) Find(data T) (int, error) {
    current := ll.Head
    index := 0
    
    for current != nil {
        // Note: This is a simple comparison, for complex types you'd need a comparator
        if fmt.Sprintf("%v", current.Data) == fmt.Sprintf("%v", data) {
            return index, nil
        }
        current = current.Next
        index++
    }
    
    return -1, LinkedListError{
        Operation: "FIND",
        Message:   "element not found",
        Index:     -1,
    }
}

func main() {
    // Test integer linked list
    fmt.Println("=== Integer Linked List Tests ===")
    intList := NewLinkedList[int]()
    
    // Append elements
    for i := 1; i <= 5; i++ {
        intList.Append(i * 10)
        fmt.Printf("Appended: %d\n", i*10)
    }
    
    // Prepend element
    intList.Prepend(0)
    fmt.Println("Prepended: 0")
    
    // Insert at specific position
    err := intList.InsertAt(3, 25)
    if err != nil {
        fmt.Printf("InsertAt error: %v\n", err)
    } else {
        fmt.Println("Inserted: 25 at index 3")
    }
    
    // Print list
    fmt.Printf("List contents: %v\n", intList.ToSlice())
    fmt.Printf("List size: %d\n", intList.GetSize())
    
    // Get element at index
    for i := 0; i < intList.GetSize(); i++ {
        element, err := intList.GetAt(i)
        if err != nil {
            fmt.Printf("GetAt error at index %d: %v\n", i, err)
        } else {
            fmt.Printf("Element at index %d: %d\n", i, element)
        }
    }
    
    // Remove elements
    fmt.Println("\n=== Removing Elements ===")
    for i := 0; i < 3; i++ {
        element, err := intList.RemoveAt(0)
        if err != nil {
            fmt.Printf("RemoveAt error: %v\n", err)
            break
        }
        fmt.Printf("Removed: %d\n", element)
    }
    
    fmt.Printf("Remaining list: %v\n", intList.ToSlice())
    
    // Test string linked list
    fmt.Println("\n=== String Linked List Tests ===")
    strList := NewLinkedList[string]()
    
    words := []string{"hello", "world", "golang", "programming"}
    for _, word := range words {
        strList.Append(word)
    }
    
    fmt.Printf("String list: %v\n", strList.ToSlice())
    
    // Find element
    index, err := strList.Find("golang")
    if err != nil {
        fmt.Printf("Find error: %v\n", err)
    } else {
        fmt.Printf("Found 'golang' at index: %d\n", index)
    }
    
    // Test error conditions
    fmt.Println("\n=== Error Tests ===")
    
    // Try to get from empty list
    emptyList := NewLinkedList[int]()
    _, err = emptyList.GetAt(0)
    if err != nil {
        fmt.Printf("GetAt from empty list: %v\n", err)
    }
    
    // Try to remove from empty list
    _, err = emptyList.RemoveAt(0)
    if err != nil {
        fmt.Printf("RemoveAt from empty list: %v\n", err)
    }
    
    // Try to insert at invalid index
    err = strList.InsertAt(100, "invalid")
    if err != nil {
        fmt.Printf("InsertAt invalid index: %v\n", err)
    }
    
    // Try to get at invalid index
    _, err = strList.GetAt(-1)
    if err != nil {
        fmt.Printf("GetAt invalid index: %v\n", err)
    }
}