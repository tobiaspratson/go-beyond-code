package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address
    Skills  []string
}

type Address struct {
    Street string
    City   string
    Zip    string
}

type ObjectBuilder struct {
    objType reflect.Type
    fields  map[string]interface{}
}

func NewObjectBuilder(obj interface{}) *ObjectBuilder {
    t := reflect.TypeOf(obj)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    
    return &ObjectBuilder{
        objType: t,
        fields:  make(map[string]interface{}),
    }
}

func (b *ObjectBuilder) Set(fieldName string, value interface{}) *ObjectBuilder {
    b.fields[fieldName] = value
    return b
}

func (b *ObjectBuilder) Build() interface{} {
    // Create new instance
    newObj := reflect.New(b.objType).Elem()
    
    // Set field values
    for fieldName, value := range b.fields {
        field := newObj.FieldByName(fieldName)
        if field.IsValid() && field.CanSet() {
            val := reflect.ValueOf(value)
            if val.Type().AssignableTo(field.Type()) {
                field.Set(val)
            } else if val.Type().ConvertibleTo(field.Type()) {
                field.Set(val.Convert(field.Type()))
            }
        }
    }
    
    return newObj.Interface()
}

func (b *ObjectBuilder) BuildPtr() interface{} {
    // Create new instance as pointer
    newObj := reflect.New(b.objType)
    elem := newObj.Elem()
    
    // Set field values
    for fieldName, value := range b.fields {
        field := elem.FieldByName(fieldName)
        if field.IsValid() && field.CanSet() {
            val := reflect.ValueOf(value)
            if val.Type().AssignableTo(field.Type()) {
                field.Set(val)
            } else if val.Type().ConvertibleTo(field.Type()) {
                field.Set(val.Convert(field.Type()))
            }
        }
    }
    
    return newObj.Interface()
}

func main() {
    // Build a Person object
    person := NewObjectBuilder(Person{}).
        Set("Name", "Alice").
        Set("Age", 25).
        Set("Email", "alice@example.com").
        Set("Address", Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        }).
        Set("Skills", []string{"Go", "Python", "JavaScript"}).
        Build()
    
    fmt.Printf("Built person: %+v\n", person)
    
    // Build a Person pointer
    personPtr := NewObjectBuilder(Person{}).
        Set("Name", "Bob").
        Set("Age", 30).
        Set("Email", "bob@example.com").
        Set("Address", Address{
            Street: "456 Oak Ave",
            City:   "San Francisco",
            Zip:    "94102",
        }).
        Set("Skills", []string{"Java", "Go", "Docker"}).
        BuildPtr()
    
    fmt.Printf("Built person pointer: %+v\n", personPtr)
}