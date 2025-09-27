package main

import "fmt"

// Generic map implementation with multiple type parameters
// K is the key type (must be comparable), V is the value type (can be any)
type GenericMap[K comparable, V any] struct {
    data map[K]V
}

// Constructor for creating a new generic map
func NewGenericMap[K comparable, V any]() *GenericMap[K, V] {
    return &GenericMap[K, V]{
        data: make(map[K]V),
    }
}

// Set a key-value pair
func (m *GenericMap[K, V]) Set(key K, value V) {
    m.data[key] = value
}

// Get a value by key
func (m *GenericMap[K, V]) Get(key K) (V, bool) {
    value, exists := m.data[key]
    return value, exists
}

// Delete a key
func (m *GenericMap[K, V]) Delete(key K) {
    delete(m.data, key)
}

// Check if key exists
func (m *GenericMap[K, V]) Has(key K) bool {
    _, exists := m.data[key]
    return exists
}

// Get all keys
func (m *GenericMap[K, V]) Keys() []K {
    keys := make([]K, 0, len(m.data))
    for k := range m.data {
        keys = append(keys, k)
    }
    return keys
}

// Get all values
func (m *GenericMap[K, V]) Values() []V {
    values := make([]V, 0, len(m.data))
    for _, v := range m.data {
        values = append(values, v)
    }
    return values
}

// Get map size
func (m *GenericMap[K, V]) Size() int {
    return len(m.data)
}

// Clear all entries
func (m *GenericMap[K, V]) Clear() {
    m.data = make(map[K]V)
}

// Update a value if it exists, otherwise set it
func (m *GenericMap[K, V]) Update(key K, updater func(V) V) bool {
    if value, exists := m.data[key]; exists {
        m.data[key] = updater(value)
        return true
    }
    return false
}

// Get or set a default value
func (m *GenericMap[K, V]) GetOrSet(key K, defaultValue V) V {
    if value, exists := m.data[key]; exists {
        return value
    }
    m.data[key] = defaultValue
    return defaultValue
}

func main() {
    // Create a map with string keys and int values
    stringIntMap := NewGenericMap[string, int]()
    
    // Set some values
    stringIntMap.Set("apple", 5)
    stringIntMap.Set("banana", 3)
    stringIntMap.Set("cherry", 8)
    
    fmt.Printf("Map size: %d\n", stringIntMap.Size())
    fmt.Printf("Keys: %v\n", stringIntMap.Keys())
    fmt.Printf("Values: %v\n", stringIntMap.Values())
    
    // Get values
    if value, exists := stringIntMap.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }
    
    if value, exists := stringIntMap.Get("grape"); exists {
        fmt.Printf("grape: %d\n", value)
    } else {
        fmt.Println("grape not found")
    }
    
    // Test Update method
    updated := stringIntMap.Update("apple", func(v int) int {
        return v * 2
    })
    fmt.Printf("Updated apple: %t\n", updated)
    if value, exists := stringIntMap.Get("apple"); exists {
        fmt.Printf("apple after update: %d\n", value)
    }
    
    // Test GetOrSet method
    defaultValue := stringIntMap.GetOrSet("grape", 10)
    fmt.Printf("grape default value: %d\n", defaultValue)
    
    // Create a map with int keys and string values
    intStringMap := NewGenericMap[int, string]()
    
    // Set some values
    intStringMap.Set(1, "one")
    intStringMap.Set(2, "two")
    intStringMap.Set(3, "three")
    
    fmt.Printf("\nInt-String Map size: %d\n", intStringMap.Size())
    fmt.Printf("Keys: %v\n", intStringMap.Keys())
    fmt.Printf("Values: %v\n", intStringMap.Values())
    
    // Create a map with custom struct keys
    type Person struct {
        Name string
        Age  int
    }
    
    personMap := NewGenericMap[Person, string]()
    personMap.Set(Person{Name: "Alice", Age: 30}, "Engineer")
    personMap.Set(Person{Name: "Bob", Age: 25}, "Designer")
    
    fmt.Printf("\nPerson Map size: %d\n", personMap.Size())
    
    // Note: Person struct is comparable because all its fields are comparable
    if job, exists := personMap.Get(Person{Name: "Alice", Age: 30}); exists {
        fmt.Printf("Alice's job: %s\n", job)
    }
}