package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string
    Age     int
    Tags    []string
    Address *Address
    Scores  map[string]int
}

type Address struct {
    Street string
    City   string
    Zip    string
}

func deepCopy(src interface{}) interface{} {
    val := reflect.ValueOf(src)
    
    // Handle nil values
    if !val.IsValid() {
        return nil
    }
    
    // Handle pointers
    if val.Kind() == reflect.Ptr {
        if val.IsNil() {
            return reflect.New(val.Type().Elem()).Interface()
        }
        return deepCopy(val.Elem().Interface())
    }
    
    // Handle different types
    switch val.Kind() {
    case reflect.Struct:
        return deepCopyStruct(val)
    case reflect.Slice:
        return deepCopySlice(val)
    case reflect.Map:
        return deepCopyMap(val)
    case reflect.Array:
        return deepCopyArray(val)
    case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
         reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
         reflect.Float32, reflect.Float64, reflect.Bool:
        // Primitive types - return a copy
        newVal := reflect.New(val.Type()).Elem()
        newVal.Set(val)
        return newVal.Interface()
    default:
        // For other types, return the original
        return src
    }
}

func deepCopyStruct(val reflect.Value) interface{} {
    newVal := reflect.New(val.Type()).Elem()
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        newField := newVal.Field(i)
        
        if newField.CanSet() {
            newField.Set(reflect.ValueOf(deepCopy(field.Interface())))
        }
    }
    
    return newVal.Interface()
}

func deepCopySlice(val reflect.Value) interface{} {
    newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
    
    for i := 0; i < val.Len(); i++ {
        newSlice.Index(i).Set(reflect.ValueOf(deepCopy(val.Index(i).Interface())))
    }
    
    return newSlice.Interface()
}

func deepCopyMap(val reflect.Value) interface{} {
    newMap := reflect.MakeMap(val.Type())
    
    for _, key := range val.MapKeys() {
        value := val.MapIndex(key)
        newKey := reflect.ValueOf(deepCopy(key.Interface()))
        newValue := reflect.ValueOf(deepCopy(value.Interface()))
        newMap.SetMapIndex(newKey, newValue)
    }
    
    return newMap.Interface()
}

func deepCopyArray(val reflect.Value) interface{} {
    newArray := reflect.New(val.Type()).Elem()
    
    for i := 0; i < val.Len(); i++ {
        newArray.Index(i).Set(reflect.ValueOf(deepCopy(val.Index(i).Interface())))
    }
    
    return newArray.Interface()
}

func main() {
    original := Person{
        Name: "Alice",
        Age:  25,
        Tags: []string{"developer", "golang"},
        Address: &Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        },
        Scores: map[string]int{
            "math":    95,
            "science": 87,
        },
    }
    
    copy := deepCopy(original).(Person)
    copy.Name = "Bob"
    copy.Tags = append(copy.Tags, "python")
    copy.Address.City = "Los Angeles"
    copy.Scores["math"] = 100
    
    fmt.Printf("Original: %+v\n", original)
    fmt.Printf("Copy: %+v\n", copy)
    fmt.Printf("Original Address: %+v\n", original.Address)
    fmt.Printf("Copy Address: %+v\n", copy.Address)
    fmt.Printf("Original Scores: %+v\n", original.Scores)
    fmt.Printf("Copy Scores: %+v\n", copy.Scores)
}