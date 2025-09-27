package main

import "fmt"

// Generic graph node
type GraphNode[T any] struct {
    Value T
    Edges []*GraphNode[T]
}

// Generic graph
type Graph[T any] struct {
    nodes []*GraphNode[T]
}

// Create a new graph
func NewGraph[T any]() *Graph[T] {
    return &Graph[T]{}
}

// Add a node to the graph
func (g *Graph[T]) AddNode(value T) *GraphNode[T] {
    node := &GraphNode[T]{Value: value}
    g.nodes = append(g.nodes, node)
    return node
}

// Add an edge between two nodes
func (g *Graph[T]) AddEdge(from, to *GraphNode[T]) {
    from.Edges = append(from.Edges, to)
}

// Get all nodes in the graph
func (g *Graph[T]) GetNodes() []*GraphNode[T] {
    return g.nodes
}

// Depth-first search
func (g *Graph[T]) DFS(start *GraphNode[T], visited map[*GraphNode[T]]bool) []T {
    var result []T
    g.dfsRecursive(start, visited, &result)
    return result
}

// Helper function for recursive DFS
func (g *Graph[T]) dfsRecursive(node *GraphNode[T], visited map[*GraphNode[T]]bool, result *[]T) {
    if visited[node] {
        return
    }
    
    visited[node] = true
    *result = append(*result, node.Value)
    
    for _, neighbor := range node.Edges {
        g.dfsRecursive(neighbor, visited, result)
    }
}

// Breadth-first search
func (g *Graph[T]) BFS(start *GraphNode[T]) []T {
    var result []T
    visited := make(map[*GraphNode[T]]bool)
    queue := []*GraphNode[T]{start}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if visited[node] {
            continue
        }
        
        visited[node] = true
        result = append(result, node.Value)
        
        for _, neighbor := range node.Edges {
            if !visited[neighbor] {
                queue = append(queue, neighbor)
            }
        }
    }
    
    return result
}

func main() {
    // Create a graph of integers
    graph := NewGraph[int]()
    
    // Add nodes
    node1 := graph.AddNode(1)
    node2 := graph.AddNode(2)
    node3 := graph.AddNode(3)
    node4 := graph.AddNode(4)
    
    // Add edges
    graph.AddEdge(node1, node2)
    graph.AddEdge(node1, node3)
    graph.AddEdge(node2, node4)
    graph.AddEdge(node3, node4)
    
    // Perform DFS
    visited := make(map[*GraphNode[int]]bool)
    dfsResult := graph.DFS(node1, visited)
    fmt.Printf("DFS from node 1: %v\n", dfsResult)
    
    // Perform BFS
    bfsResult := graph.BFS(node1)
    fmt.Printf("BFS from node 1: %v\n", bfsResult)
    
    // Create a graph of strings
    stringGraph := NewGraph[string]()
    
    // Add nodes
    strNode1 := stringGraph.AddNode("A")
    strNode2 := stringGraph.AddNode("B")
    strNode3 := stringGraph.AddNode("C")
    
    // Add edges
    stringGraph.AddEdge(strNode1, strNode2)
    stringGraph.AddEdge(strNode2, strNode3)
    
    // Perform DFS
    strVisited := make(map[*GraphNode[string]]bool)
    strDfsResult := stringGraph.DFS(strNode1, strVisited)
    fmt.Printf("String graph DFS: %v\n", strDfsResult)
}