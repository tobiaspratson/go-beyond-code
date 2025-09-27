package main

import "fmt"

// Generic set implementation
type Set[T comparable] struct {
	items map[T]bool
}

// Create a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

// Add an item to the set
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Remove an item from the set
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Check if set contains an item
func (s *Set[T]) Contains(item T) bool {
	return s.items[item]
}

// Get set size
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Get all items in the set
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

// Union of two sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for item := range s.items {
		result.Add(item)
	}

	for item := range other.items {
		result.Add(item)
	}

	return result
}

// Intersection of two sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Difference of two sets
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

func main() {
	// Create sets
	set1 := NewSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2 := NewSet[int]()
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	fmt.Printf("Set 1: %v\n", set1.Items())
	fmt.Printf("Set 2: %v\n", set2.Items())

	// Test union
	union := set1.Union(set2)
	fmt.Printf("Union: %v\n", union.Items())

	// Test intersection
	intersection := set1.Intersection(set2)
	fmt.Printf("Intersection: %v\n", intersection.Items())

	// Test difference
	difference := set1.Difference(set2)
	fmt.Printf("Difference: %v\n", difference.Items())

	// Test with strings
	stringSet1 := NewSet[string]()
	stringSet1.Add("apple")
	stringSet1.Add("banana")
	stringSet1.Add("cherry")

	stringSet2 := NewSet[string]()
	stringSet2.Add("banana")
	stringSet2.Add("cherry")
	stringSet2.Add("date")

	fmt.Printf("String Set 1: %v\n", stringSet1.Items())
	fmt.Printf("String Set 2: %v\n", stringSet2.Items())

	stringUnion := stringSet1.Union(stringSet2)
	fmt.Printf("String Union: %v\n", stringUnion.Items())
}
