package main

import "fmt"

// Generic binary tree node
type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Generic binary tree
type BinaryTree[T comparable] struct {
	Root *TreeNode[T]
}

// Create a new binary tree
func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// Insert a value into the tree
func (bt *BinaryTree[T]) Insert(value T) {
	bt.Root = bt.insertRecursive(bt.Root, value)
}

// Helper function for recursive insertion
func (bt *BinaryTree[T]) insertRecursive(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{Value: value}
	}

	// Simple insertion strategy - left for smaller, right for larger
	// This assumes T supports comparison (you'd need a custom comparator in practice)
	// For demonstration, we'll use a simple approach
	if node.Left == nil {
		node.Left = &TreeNode[T]{Value: value}
	} else if node.Right == nil {
		node.Right = &TreeNode[T]{Value: value}
	} else {
		// Recursively insert into left subtree
		node.Left = bt.insertRecursive(node.Left, value)
	}

	return node
}

// In-order traversal
func (bt *BinaryTree[T]) InOrder() []T {
	var result []T
	bt.inOrderRecursive(bt.Root, &result)
	return result
}

// Helper function for in-order traversal
func (bt *BinaryTree[T]) inOrderRecursive(node *TreeNode[T], result *[]T) {
	if node != nil {
		bt.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		bt.inOrderRecursive(node.Right, result)
	}
}

// Pre-order traversal
func (bt *BinaryTree[T]) PreOrder() []T {
	var result []T
	bt.preOrderRecursive(bt.Root, &result)
	return result
}

// Helper function for pre-order traversal
func (bt *BinaryTree[T]) preOrderRecursive(node *TreeNode[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Value)
		bt.preOrderRecursive(node.Left, result)
		bt.preOrderRecursive(node.Right, result)
	}
}

// Post-order traversal
func (bt *BinaryTree[T]) PostOrder() []T {
	var result []T
	bt.postOrderRecursive(bt.Root, &result)
	return result
}

// Helper function for post-order traversal
func (bt *BinaryTree[T]) postOrderRecursive(node *TreeNode[T], result *[]T) {
	if node != nil {
		bt.postOrderRecursive(node.Left, result)
		bt.postOrderRecursive(node.Right, result)
		*result = append(*result, node.Value)
	}
}

// Search for a value in the tree
func (bt *BinaryTree[T]) Search(value T) bool {
	return bt.searchRecursive(bt.Root, value)
}

// Helper function for recursive search
func (bt *BinaryTree[T]) searchRecursive(node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	}

	return bt.searchRecursive(node.Left, value) || bt.searchRecursive(node.Right, value)
}

func main() {
	// Create a binary tree of integers
	tree := NewBinaryTree[int]()

	// Insert some values
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(9)

	fmt.Printf("In-order traversal: %v\n", tree.InOrder())
	fmt.Printf("Pre-order traversal: %v\n", tree.PreOrder())
	fmt.Printf("Post-order traversal: %v\n", tree.PostOrder())

	// Search for values
	fmt.Printf("Search for 5: %t\n", tree.Search(5))
	fmt.Printf("Search for 6: %t\n", tree.Search(6))

	// Create a binary tree of strings
	stringTree := NewBinaryTree[string]()
	stringTree.Insert("apple")
	stringTree.Insert("banana")
	stringTree.Insert("cherry")

	fmt.Printf("String tree in-order: %v\n", stringTree.InOrder())
}
