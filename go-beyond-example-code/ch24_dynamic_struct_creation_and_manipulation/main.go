package main

import (
    "fmt"
    "reflect"
)

// Dynamic struct builder with enhanced features
type StructBuilder struct {
    fields map[string]reflect.Type
    tags   map[string]string
}

func NewStructBuilder() *StructBuilder {
    return &StructBuilder{
        fields: make(map[string]reflect.Type),
        tags:   make(map[string]string),
    }
}

func (sb *StructBuilder) AddField(name string, fieldType reflect.Type) *StructBuilder {
    sb.fields[name] = fieldType
    return sb
}

func (sb *StructBuilder) AddFieldWithTag(name string, fieldType reflect.Type, tag string) *StructBuilder {
    sb.fields[name] = fieldType
    sb.tags[name] = tag
    return sb
}

func (sb *StructBuilder) Build() reflect.Type {
    fields := make([]reflect.StructField, 0, len(sb.fields))
    
    for name, fieldType := range sb.fields {
        field := reflect.StructField{
            Name: name,
            Type: fieldType,
        }
        
        // Add tag if specified
        if tag, exists := sb.tags[name]; exists {
            field.Tag = reflect.StructTag(tag)
        } else {
            field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, name))
        }
        
        fields = append(fields, field)
    }
    
    return reflect.StructOf(fields)
}

func (sb *StructBuilder) CreateInstance() reflect.Value {
    structType := sb.Build()
    return reflect.New(structType).Elem()
}

// Enhanced struct manipulator
type StructManipulator struct {
    instance reflect.Value
    structType reflect.Type
}

func NewStructManipulator(instance reflect.Value) *StructManipulator {
    return &StructManipulator{
        instance:   instance,
        structType: instance.Type(),
    }
}

func (sm *StructManipulator) SetField(name string, value interface{}) error {
    field := sm.instance.FieldByName(name)
    if !field.IsValid() {
        return fmt.Errorf("field %s not found", name)
    }
    
    if !field.CanSet() {
        return fmt.Errorf("field %s cannot be set", name)
    }
    
    valueVal := reflect.ValueOf(value)
    if !valueVal.Type().AssignableTo(field.Type()) {
        return fmt.Errorf("cannot assign %s to field %s of type %s", 
            valueVal.Type(), name, field.Type())
    }
    
    field.Set(valueVal)
    return nil
}

func (sm *StructManipulator) GetField(name string) (interface{}, error) {
    field := sm.instance.FieldByName(name)
    if !field.IsValid() {
        return nil, fmt.Errorf("field %s not found", name)
    }
    
    return field.Interface(), nil
}

func (sm *StructManipulator) GetFieldNames() []string {
    var names []string
    for i := 0; i < sm.structType.NumField(); i++ {
        names = append(names, sm.structType.Field(i).Name)
    }
    return names
}

func (sm *StructManipulator) GetFieldInfo() map[string]interface{} {
    info := make(map[string]interface{})
    for i := 0; i < sm.structType.NumField(); i++ {
        field := sm.structType.Field(i)
        value := sm.instance.Field(i)
        info[field.Name] = value.Interface()
    }
    return info
}

func (sm *StructManipulator) ToMap() map[string]interface{} {
    return sm.GetFieldInfo()
}

func main() {
    fmt.Println("=== Dynamic Struct Creation ===")
    
    // Create a dynamic struct with various field types
    builder := NewStructBuilder()
    builder.AddFieldWithTag("Name", reflect.TypeOf(""), `json:"name" validate:"required"`)
    builder.AddFieldWithTag("Age", reflect.TypeOf(0), `json:"age" validate:"min=18"`)
    builder.AddFieldWithTag("Email", reflect.TypeOf(""), `json:"email" validate:"email"`)
    builder.AddFieldWithTag("Active", reflect.TypeOf(false), `json:"active"`)
    builder.AddFieldWithTag("Tags", reflect.TypeOf([]string{}), `json:"tags"`)
    
    // Build the struct type
    structType := builder.Build()
    fmt.Printf("Dynamic struct type: %v\n", structType)
    
    // Create an instance
    instance := builder.CreateInstance()
    manipulator := NewStructManipulator(instance)
    
    fmt.Println("\n=== Setting Fields ===")
    // Set values using the manipulator
    fields := map[string]interface{}{
        "Name":   "Alice Johnson",
        "Age":    28,
        "Email":  "alice@example.com",
        "Active": true,
        "Tags":   []string{"admin", "user", "premium"},
    }
    
    for name, value := range fields {
        err := manipulator.SetField(name, value)
        if err != nil {
            fmt.Printf("Error setting %s: %v\n", name, err)
        } else {
            fmt.Printf("Set %s to %v\n", name, value)
        }
    }
    
    fmt.Println("\n=== Field Information ===")
    fieldInfo := manipulator.GetFieldInfo()
    for name, value := range fieldInfo {
        fmt.Printf("%s: %v (type: %T)\n", name, value, value)
    }
    
    fmt.Println("\n=== Field Names ===")
    fieldNames := manipulator.GetFieldNames()
    for _, name := range fieldNames {
        fmt.Printf("- %s\n", name)
    }
    
    fmt.Println("\n=== Individual Field Access ===")
    name, _ := manipulator.GetField("Name")
    age, _ := manipulator.GetField("Age")
    email, _ := manipulator.GetField("Email")
    
    fmt.Printf("Name: %s, Age: %d, Email: %s\n", name, age, email)
    
    fmt.Println("\n=== Convert to Map ===")
    dataMap := manipulator.ToMap()
    fmt.Printf("As map: %+v\n", dataMap)
}