package main

import "fmt"

type Counter struct {
    Value int
}

// Value receiver - cannot modify the original
func (c Counter) GetValue() int {
    return c.Value
}

// Value receiver - creates a copy, doesn't modify original
func (c Counter) Increment() Counter {
    c.Value++  // This modifies the COPY, not the original
    return c   // Returns the modified copy
}

// Value receiver - creates a copy for display
func (c Counter) String() string {
    return fmt.Sprintf("Counter{Value: %d}", c.Value)
}

// Pointer receiver - CAN modify the original
func (c *Counter) IncrementByRef() {
    c.Value++  // This modifies the ORIGINAL
}

// Pointer receiver - sets a new value
func (c *Counter) SetValue(newValue int) {
    c.Value = newValue
}

// Pointer receiver - resets to zero
func (c *Counter) Reset() {
    c.Value = 0
}

func main() {
    counter := Counter{Value: 0}
    
    fmt.Printf("Initial: %s\n", counter.String())
    
    // Value receiver doesn't modify original
    newCounter := counter.Increment()
    fmt.Printf("After value increment - original: %s\n", counter.String())
    fmt.Printf("After value increment - returned: %s\n", newCounter.String())
    
    // Pointer receiver modifies original
    counter.IncrementByRef()
    fmt.Printf("After pointer increment: %s\n", counter.String())
    
    // More pointer receiver operations
    counter.SetValue(10)
    fmt.Printf("After SetValue(10): %s\n", counter.String())
    
    counter.Reset()
    fmt.Printf("After Reset: %s\n", counter.String())
}