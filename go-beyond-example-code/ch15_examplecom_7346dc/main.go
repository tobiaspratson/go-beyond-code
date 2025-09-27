package main

import (
    "fmt"
    "reflect"
)

func dynamicTypeCreation() {
    fmt.Println("=== Dynamic Type Creation ===")
    
    // Create a new slice type
    intSliceType := reflect.SliceOf(reflect.TypeOf(0))
    fmt.Printf("Created slice type: %v\n", intSliceType)
    
    // Create a new slice value
    intSlice := reflect.MakeSlice(intSliceType, 3, 3)
    fmt.Printf("Created slice value: %v\n", intSlice.Interface())
    
    // Set values in the slice
    for i := 0; i < intSlice.Len(); i++ {
        intSlice.Index(i).SetInt(int64(i * 10))
    }
    fmt.Printf("After setting values: %v\n", intSlice.Interface())
    
    // Create a new map type
    stringIntMapType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
    fmt.Printf("Created map type: %v\n", stringIntMapType)
    
    // Create a new map value
    stringIntMap := reflect.MakeMap(stringIntMapType)
    fmt.Printf("Created map value: %v\n", stringIntMap.Interface())
    
    // Set values in the map
    key1 := reflect.ValueOf("a")
    value1 := reflect.ValueOf(1)
    stringIntMap.SetMapIndex(key1, value1)
    
    key2 := reflect.ValueOf("b")
    value2 := reflect.ValueOf(2)
    stringIntMap.SetMapIndex(key2, value2)
    
    fmt.Printf("After setting values: %v\n", stringIntMap.Interface())
}

func dynamicStructCreation() {
    fmt.Println("\n=== Dynamic Struct Creation ===")
    
    // Create struct fields
    fields := []reflect.StructField{
        {
            Name: "Name",
            Type: reflect.TypeOf(""),
            Tag:  `json:"name"`,
        },
        {
            Name: "Age",
            Type: reflect.TypeOf(0),
            Tag:  `json:"age"`,
        },
        {
            Name: "Email",
            Type: reflect.TypeOf(""),
            Tag:  `json:"email"`,
        },
    }
    
    // Create a new struct type
    structType := reflect.StructOf(fields)
    fmt.Printf("Created struct type: %v\n", structType)
    
    // Create a new struct value
    structValue := reflect.New(structType).Elem()
    fmt.Printf("Created struct value: %v\n", structValue.Interface())
    
    // Set field values
    nameField := structValue.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("John Doe")
    }
    
    ageField := structValue.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(30)
    }
    
    emailField := structValue.FieldByName("Email")
    if emailField.IsValid() && emailField.CanSet() {
        emailField.SetString("john@example.com")
    }
    
    fmt.Printf("After setting values: %v\n", structValue.Interface())
}

func main() {
    dynamicTypeCreation()
    dynamicStructCreation()
}