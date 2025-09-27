package main

import (
    "fmt"
    "reflect"
)

func modifyValues() {
    fmt.Println("=== Basic Value Modification ===")
    
    // Modify an integer
    x := 42
    v := reflect.ValueOf(&x).Elem() // Get addressable value
    fmt.Printf("Before: %d\n", x)
    v.SetInt(100)
    fmt.Printf("After: %d\n", x)
    
    // Modify a string
    s := "Hello"
    vs := reflect.ValueOf(&s).Elem()
    fmt.Printf("Before: %s\n", s)
    vs.SetString("World")
    fmt.Printf("After: %s\n", s)
    
    // Modify a boolean
    b := false
    vb := reflect.ValueOf(&b).Elem()
    fmt.Printf("Before: %t\n", b)
    vb.SetBool(true)
    fmt.Printf("After: %t\n", b)
    
    // Modify a float
    f := 3.14
    vf := reflect.ValueOf(&f).Elem()
    fmt.Printf("Before: %f\n", f)
    vf.SetFloat(2.71)
    fmt.Printf("After: %f\n", f)
}

func modifySlices() {
    fmt.Println("\n=== Slice Modification ===")
    
    // Modify slice elements
    slice := []int{1, 2, 3}
    vslice := reflect.ValueOf(&slice).Elem()
    fmt.Printf("Before: %v\n", slice)
    
    // Modify existing element
    vslice.Index(0).SetInt(10)
    fmt.Printf("After modifying [0]: %v\n", slice)
    
    // Add element to slice
    newSlice := reflect.Append(vslice, reflect.ValueOf(4))
    vslice.Set(newSlice)
    fmt.Printf("After appending: %v\n", slice)
    
    // Create new slice with different elements
    newSlice2 := reflect.MakeSlice(vslice.Type(), 2, 2)
    newSlice2.Index(0).SetInt(100)
    newSlice2.Index(1).SetInt(200)
    vslice.Set(newSlice2)
    fmt.Printf("After replacing: %v\n", slice)
}

func modifyMaps() {
    fmt.Println("\n=== Map Modification ===")
    
    // Create and modify a map
    m := make(map[string]int)
    vm := reflect.ValueOf(&m).Elem()
    
    fmt.Printf("Before: %v\n", m)
    
    // Add key-value pairs
    key1 := reflect.ValueOf("a")
    value1 := reflect.ValueOf(1)
    vm.SetMapIndex(key1, value1)
    
    key2 := reflect.ValueOf("b")
    value2 := reflect.ValueOf(2)
    vm.SetMapIndex(key2, value2)
    
    fmt.Printf("After adding entries: %v\n", m)
    
    // Modify existing value
    key1 = reflect.ValueOf("a")
    value1 = reflect.ValueOf(10)
    vm.SetMapIndex(key1, value1)
    
    fmt.Printf("After modifying 'a': %v\n", m)
    
    // Delete a key
    key2 = reflect.ValueOf("b")
    vm.SetMapIndex(key2, reflect.Value{}) // Setting to zero value deletes the key
    fmt.Printf("After deleting 'b': %v\n", m)
}

func modifyStructs() {
    fmt.Println("\n=== Struct Modification ===")
    
    type Person struct {
        Name string
        Age  int
    }
    
    p := Person{Name: "Alice", Age: 25}
    vp := reflect.ValueOf(&p).Elem()
    
    fmt.Printf("Before: %+v\n", p)
    
    // Modify struct fields
    nameField := vp.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("Bob")
    }
    
    ageField := vp.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(30)
    }
    
    fmt.Printf("After: %+v\n", p)
}

func main() {
    modifyValues()
    modifySlices()
    modifyMaps()
    modifyStructs()
}