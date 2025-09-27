package main

import "fmt"

// Numeric constraint for types that support addition
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Generic matrix type
type Matrix[T Numeric] struct {
	data [][]T
	rows int
	cols int
}

// Create a new matrix
func NewMatrix[T Numeric](rows, cols int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	return &Matrix[T]{
		data: data,
		rows: rows,
		cols: cols,
	}
}

// Set a value at position (row, col)
func (m *Matrix[T]) Set(row, col int, value T) {
	m.data[row][col] = value
}

// Get a value at position (row, col)
func (m *Matrix[T]) Get(row, col int) T {
	return m.data[row][col]
}

// Add two matrices
func (m *Matrix[T]) Add(other *Matrix[T]) *Matrix[T] {
	if m.rows != other.rows || m.cols != other.cols {
		return nil
	}

	result := NewMatrix[T](m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.Set(i, j, m.Get(i, j)+other.Get(i, j))
		}
	}
	return result
}

// Print the matrix
func (m *Matrix[T]) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%v ", m.Get(i, j))
		}
		fmt.Println()
	}
}

func main() {
	// Create matrices
	matrix1 := NewMatrix[int](2, 3)
	matrix1.Set(0, 0, 1)
	matrix1.Set(0, 1, 2)
	matrix1.Set(0, 2, 3)
	matrix1.Set(1, 0, 4)
	matrix1.Set(1, 1, 5)
	matrix1.Set(1, 2, 6)

	matrix2 := NewMatrix[int](2, 3)
	matrix2.Set(0, 0, 7)
	matrix2.Set(0, 1, 8)
	matrix2.Set(0, 2, 9)
	matrix2.Set(1, 0, 10)
	matrix2.Set(1, 1, 11)
	matrix2.Set(1, 2, 12)

	fmt.Println("Matrix 1:")
	matrix1.Print()

	fmt.Println("Matrix 2:")
	matrix2.Print()

	// Add matrices
	result := matrix1.Add(matrix2)
	if result != nil {
		fmt.Println("Result:")
		result.Print()
	}
}
