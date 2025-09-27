package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generic binary tree node
type TreeNode[T constraints.Ordered] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Generic binary tree
type BinaryTree[T constraints.Ordered] struct {
	Root *TreeNode[T]
}

// Create a new binary tree
func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// Insert a value into the tree
func (bt *BinaryTree[T]) Insert(value T) {
	bt.Root = bt.insert(bt.Root, value)
}

// Helper function for insertion
func (bt *BinaryTree[T]) insert(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{Value: value}
	}

	if value < node.Value {
		node.Left = bt.insert(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.insert(node.Right, value)
	}

	return node
}

// Search for a value in the tree
func (bt *BinaryTree[T]) Search(value T) bool {
	return bt.search(bt.Root, value)
}

// Helper function for search
func (bt *BinaryTree[T]) search(node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return bt.search(node.Left, value)
	} else {
		return bt.search(node.Right, value)
	}
}

// In-order traversal
func (bt *BinaryTree[T]) InOrder() []T {
	var result []T
	bt.inOrder(bt.Root, &result)
	return result
}

// Helper function for in-order traversal
func (bt *BinaryTree[T]) inOrder(node *TreeNode[T], result *[]T) {
	if node != nil {
		bt.inOrder(node.Left, result)
		*result = append(*result, node.Value)
		bt.inOrder(node.Right, result)
	}
}

// Pre-order traversal
func (bt *BinaryTree[T]) PreOrder() []T {
	var result []T
	bt.preOrder(bt.Root, &result)
	return result
}

// Helper function for pre-order traversal
func (bt *BinaryTree[T]) preOrder(node *TreeNode[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Value)
		bt.preOrder(node.Left, result)
		bt.preOrder(node.Right, result)
	}
}

// Post-order traversal
func (bt *BinaryTree[T]) PostOrder() []T {
	var result []T
	bt.postOrder(bt.Root, &result)
	return result
}

// Helper function for post-order traversal
func (bt *BinaryTree[T]) postOrder(node *TreeNode[T], result *[]T) {
	if node != nil {
		bt.postOrder(node.Left, result)
		bt.postOrder(node.Right, result)
		*result = append(*result, node.Value)
	}
}

// Get the height of the tree
func (bt *BinaryTree[T]) Height() int {
	return bt.height(bt.Root)
}

// Helper function for height calculation
func (bt *BinaryTree[T]) height(node *TreeNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := bt.height(node.Left)
	rightHeight := bt.height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Check if the tree is empty
func (bt *BinaryTree[T]) IsEmpty() bool {
	return bt.Root == nil
}

// Get the size of the tree
func (bt *BinaryTree[T]) Size() int {
	return bt.size(bt.Root)
}

// Helper function for size calculation
func (bt *BinaryTree[T]) size(node *TreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + bt.size(node.Left) + bt.size(node.Right)
}

func main() {
	// Create a binary tree of integers
	intTree := NewBinaryTree[int]()

	// Insert some integers
	intTree.Insert(5)
	intTree.Insert(3)
	intTree.Insert(7)
	intTree.Insert(1)
	intTree.Insert(9)
	intTree.Insert(4)
	intTree.Insert(6)

	fmt.Printf("Tree size: %d\n", intTree.Size())
	fmt.Printf("Tree height: %d\n", intTree.Height())

	// Search for values
	fmt.Printf("Search 5: %t\n", intTree.Search(5))
	fmt.Printf("Search 6: %t\n", intTree.Search(6))
	fmt.Printf("Search 10: %t\n", intTree.Search(10))

	// Different traversals
	fmt.Printf("In-order: %v\n", intTree.InOrder())
	fmt.Printf("Pre-order: %v\n", intTree.PreOrder())
	fmt.Printf("Post-order: %v\n", intTree.PostOrder())

	// Create a binary tree of strings
	stringTree := NewBinaryTree[string]()

	// Insert some strings
	stringTree.Insert("banana")
	stringTree.Insert("apple")
	stringTree.Insert("cherry")
	stringTree.Insert("date")
	stringTree.Insert("elderberry")
	stringTree.Insert("fig")

	fmt.Printf("\nString tree size: %d\n", stringTree.Size())
	fmt.Printf("String tree height: %d\n", stringTree.Height())

	// Search for values
	fmt.Printf("Search 'apple': %t\n", stringTree.Search("apple"))
	fmt.Printf("Search 'grape': %t\n", stringTree.Search("grape"))

	// Different traversals
	fmt.Printf("In-order: %v\n", stringTree.InOrder())
	fmt.Printf("Pre-order: %v\n", stringTree.PreOrder())
	fmt.Printf("Post-order: %v\n", stringTree.PostOrder())
}
