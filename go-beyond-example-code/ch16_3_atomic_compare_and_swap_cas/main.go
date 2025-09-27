package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

type LockFreeStack struct {
	head unsafe.Pointer
}

type Node struct {
	value int
	next  unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	newNode := &Node{value: value}

	for {
		head := atomic.LoadPointer(&s.head)
		newNode.next = head

		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}

		node := (*Node)(head)
		next := atomic.LoadPointer(&node.next)

		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return node.value, true
		}
	}
}

func main() {
	stack := NewLockFreeStack()
	var wg sync.WaitGroup

	// Push values
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			stack.Push(id)
			fmt.Printf("Pushed %d\n", id)
		}(i)
	}

	wg.Wait()

	// Pop values
	for {
		if value, ok := stack.Pop(); ok {
			fmt.Printf("Popped %d\n", value)
		} else {
			break
		}
	}
}
