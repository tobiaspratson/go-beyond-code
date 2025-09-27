package main

import (
    "fmt"
    "math"
)

type Matrix struct {
    data [][]float64
    rows int
    cols int
}

// Create a new matrix
func NewMatrix(rows, cols int) *Matrix {
    if rows <= 0 || cols <= 0 {
        return nil
    }
    
    data := make([][]float64, rows)
    for i := range data {
        data[i] = make([]float64, cols)
    }
    return &Matrix{data: data, rows: rows, cols: cols}
}

// Create matrix from 2D slice
func NewMatrixFromData(data [][]float64) *Matrix {
    if len(data) == 0 {
        return nil
    }
    
    rows := len(data)
    cols := len(data[0])
    
    // Check if all rows have same length
    for i := 1; i < rows; i++ {
        if len(data[i]) != cols {
            return nil
        }
    }
    
    return &Matrix{data: data, rows: rows, cols: cols}
}

// Create identity matrix
func IdentityMatrix(n int) *Matrix {
    if n <= 0 {
        return nil
    }
    
    matrix := NewMatrix(n, n)
    for i := 0; i < n; i++ {
        matrix.Set(i, i, 1.0)
    }
    return matrix
}

// Create zero matrix
func ZeroMatrix(rows, cols int) *Matrix {
    return NewMatrix(rows, cols)
}

// Set value at position
func (m *Matrix) Set(row, col int, value float64) error {
    if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
        return fmt.Errorf("index out of bounds")
    }
    m.data[row][col] = value
    return nil
}

// Get value at position
func (m *Matrix) Get(row, col int) (float64, error) {
    if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
        return 0, fmt.Errorf("index out of bounds")
    }
    return m.data[row][col], nil
}

// Matrix addition
func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
    if m.rows != other.rows || m.cols != other.cols {
        return nil, fmt.Errorf("matrix dimensions don't match for addition")
    }
    
    result := NewMatrix(m.rows, m.cols)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            result.Set(i, j, m.data[i][j]+other.data[i][j])
        }
    }
    return result, nil
}

// Matrix subtraction
func (m *Matrix) Subtract(other *Matrix) (*Matrix, error) {
    if m.rows != other.rows || m.cols != other.cols {
        return nil, fmt.Errorf("matrix dimensions don't match for subtraction")
    }
    
    result := NewMatrix(m.rows, m.cols)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            result.Set(i, j, m.data[i][j]-other.data[i][j])
        }
    }
    return result, nil
}

// Scalar multiplication
func (m *Matrix) ScalarMultiply(scalar float64) *Matrix {
    result := NewMatrix(m.rows, m.cols)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            result.Set(i, j, m.data[i][j]*scalar)
        }
    }
    return result
}

// Matrix multiplication
func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
    if m.cols != other.rows {
        return nil, fmt.Errorf("matrix dimensions don't match for multiplication")
    }
    
    result := NewMatrix(m.rows, other.cols)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < other.cols; j++ {
            sum := 0.0
            for k := 0; k < m.cols; k++ {
                sum += m.data[i][k] * other.data[k][j]
            }
            result.Set(i, j, sum)
        }
    }
    return result, nil
}

// Matrix transpose
func (m *Matrix) Transpose() *Matrix {
    result := NewMatrix(m.cols, m.rows)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            result.Set(j, i, m.data[i][j])
        }
    }
    return result
}

// Calculate determinant (for square matrices)
func (m *Matrix) Determinant() (float64, error) {
    if m.rows != m.cols {
        return 0, fmt.Errorf("determinant only defined for square matrices")
    }
    
    if m.rows == 1 {
        return m.data[0][0], nil
    }
    
    if m.rows == 2 {
        return m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0], nil
    }
    
    // For larger matrices, use cofactor expansion
    det := 0.0
    for j := 0; j < m.cols; j++ {
        cofactor := m.cofactor(0, j)
        det += m.data[0][j] * cofactor
    }
    return det, nil
}

// Calculate cofactor
func (m *Matrix) cofactor(row, col int) float64 {
    minor := m.minor(row, col)
    det, _ := minor.Determinant()
    
    if (row+col)%2 == 1 {
        return -det
    }
    return det
}

// Get minor matrix
func (m *Matrix) minor(row, col int) *Matrix {
    result := NewMatrix(m.rows-1, m.cols-1)
    r := 0
    for i := 0; i < m.rows; i++ {
        if i == row {
            continue
        }
        c := 0
        for j := 0; j < m.cols; j++ {
            if j == col {
                continue
            }
            result.Set(r, c, m.data[i][j])
            c++
        }
        r++
    }
    return result
}

// Check if matrix is square
func (m *Matrix) IsSquare() bool {
    return m.rows == m.cols
}

// Check if matrix is symmetric
func (m *Matrix) IsSymmetric() bool {
    if !m.IsSquare() {
        return false
    }
    
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            if m.data[i][j] != m.data[j][i] {
                return false
            }
        }
    }
    return true
}

// Calculate trace (sum of diagonal elements)
func (m *Matrix) Trace() (float64, error) {
    if !m.IsSquare() {
        return 0, fmt.Errorf("trace only defined for square matrices")
    }
    
    sum := 0.0
    for i := 0; i < m.rows; i++ {
        sum += m.data[i][i]
    }
    return sum, nil
}

// Calculate matrix norm (Frobenius norm)
func (m *Matrix) Norm() float64 {
    sum := 0.0
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            sum += m.data[i][j] * m.data[i][j]
        }
    }
    return math.Sqrt(sum)
}

// Print matrix
func (m *Matrix) Print() {
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            fmt.Printf("%8.3f ", m.data[i][j])
        }
        fmt.Println()
    }
}

// Print matrix with custom precision
func (m *Matrix) PrintWithPrecision(precision int) {
    format := fmt.Sprintf("%%8.%df ", precision)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            fmt.Printf(format, m.data[i][j])
        }
        fmt.Println()
    }
}

func main() {
    fmt.Println("=== Advanced Matrix Operations ===")
    
    // Create matrices
    fmt.Println("\n--- Creating Matrices ---")
    a := NewMatrix(3, 3)
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            a.Set(i, j, float64(i*3+j+1))
        }
    }
    
    fmt.Println("Matrix A:")
    a.Print()
    
    // Identity matrix
    identity := IdentityMatrix(3)
    fmt.Println("\nIdentity Matrix:")
    identity.Print()
    
    // Matrix operations
    fmt.Println("\n--- Matrix Operations ---")
    
    // Addition
    sum, err := a.Add(identity)
    if err == nil {
        fmt.Println("A + I:")
        sum.Print()
    }
    
    // Scalar multiplication
    scaled := a.ScalarMultiply(2.0)
    fmt.Println("\n2 * A:")
    scaled.Print()
    
    // Transpose
    transposed := a.Transpose()
    fmt.Println("\nA^T:")
    transposed.Print()
    
    // Matrix properties
    fmt.Println("\n--- Matrix Properties ---")
    fmt.Printf("Is square: %t\n", a.IsSquare())
    fmt.Printf("Is symmetric: %t\n", a.IsSymmetric())
    
    trace, _ := a.Trace()
    fmt.Printf("Trace: %.3f\n", trace)
    
    norm := a.Norm()
    fmt.Printf("Frobenius norm: %.3f\n", norm)
    
    // Determinant
    det, err := a.Determinant()
    if err == nil {
        fmt.Printf("Determinant: %.3f\n", det)
    }
    
    // Matrix multiplication
    fmt.Println("\n--- Matrix Multiplication ---")
    b := NewMatrix(3, 2)
    for i := 0; i < 3; i++ {
        for j := 0; j < 2; j++ {
            b.Set(i, j, float64(i*2+j+1))
        }
    }
    
    fmt.Println("Matrix B:")
    b.Print()
    
    product, err := a.Multiply(b)
    if err == nil {
        fmt.Println("A * B:")
        product.Print()
    }
    
    // Symmetric matrix example
    fmt.Println("\n--- Symmetric Matrix ---")
    symmetric := NewMatrix(3, 3)
    symmetric.Set(0, 0, 1)
    symmetric.Set(0, 1, 2)
    symmetric.Set(0, 2, 3)
    symmetric.Set(1, 0, 2)
    symmetric.Set(1, 1, 4)
    symmetric.Set(1, 2, 5)
    symmetric.Set(2, 0, 3)
    symmetric.Set(2, 1, 5)
    symmetric.Set(2, 2, 6)
    
    fmt.Println("Symmetric Matrix:")
    symmetric.Print()
    fmt.Printf("Is symmetric: %t\n", symmetric.IsSymmetric())
}