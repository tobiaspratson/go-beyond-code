package main

import (
    "fmt"
    "reflect"
)

// Type inspector for detailed analysis
type TypeInspector struct {
    typ reflect.Type
}

func NewTypeInspector(obj interface{}) *TypeInspector {
    typ := reflect.TypeOf(obj)
    if typ.Kind() == reflect.Ptr {
        typ = typ.Elem()
    }
    return &TypeInspector{typ: typ}
}

func (ti *TypeInspector) GetTypeInfo() map[string]interface{} {
    info := make(map[string]interface{})
    
    info["name"] = ti.typ.Name()
    info["kind"] = ti.typ.Kind().String()
    info["package"] = ti.typ.PkgPath()
    info["num_methods"] = ti.typ.NumMethod()
    info["num_fields"] = ti.typ.NumField()
    
    // Get method information
    methods := make([]map[string]interface{}, ti.typ.NumMethod())
    for i := 0; i < ti.typ.NumMethod(); i++ {
        method := ti.typ.Method(i)
        methods[i] = map[string]interface{}{
            "name":     method.Name,
            "type":     method.Type.String(),
            "pkg_path": method.PkgPath,
        }
    }
    info["methods"] = methods
    
    // Get field information
    fields := make([]map[string]interface{}, ti.typ.NumField())
    for i := 0; i < ti.typ.NumField(); i++ {
        field := ti.typ.Field(i)
        fields[i] = map[string]interface{}{
            "name":     field.Name,
            "type":     field.Type.String(),
            "tag":      string(field.Tag),
            "pkg_path": field.PkgPath,
            "anonymous": field.Anonymous,
        }
    }
    info["fields"] = fields
    
    return info
}

func (ti *TypeInspector) GetMethodSignatures() []string {
    var signatures []string
    for i := 0; i < ti.typ.NumMethod(); i++ {
        method := ti.typ.Method(i)
        signatures = append(signatures, fmt.Sprintf("%s%s", method.Name, method.Type.String()))
    }
    return signatures
}

func (ti *TypeInspector) GetFieldTypes() map[string]string {
    fieldTypes := make(map[string]string)
    for i := 0; i < ti.typ.NumField(); i++ {
        field := ti.typ.Field(i)
        fieldTypes[field.Name] = field.Type.String()
    }
    return fieldTypes
}

func (ti *TypeInspector) ImplementsInterface(iface reflect.Type) bool {
    return ti.typ.Implements(iface)
}

// Example interface for testing
type Stringer interface {
    String() string
}

// Example struct for testing
type Person struct {
    Name    string `json:"name" validate:"required"`
    Age     int    `json:"age" validate:"min=18"`
    Email   string `json:"email" validate:"email"`
    private string // unexported field
}

func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", p.Name, p.Age, p.Email)
}

func (p Person) GetInfo() map[string]interface{} {
    return map[string]interface{}{
        "name":  p.Name,
        "age":   p.Age,
        "email": p.Email,
    }
}

func main() {
    person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john@example.com",
    }
    
    inspector := NewTypeInspector(person)
    
    fmt.Println("=== Type Information ===")
    typeInfo := inspector.GetTypeInfo()
    for key, value := range typeInfo {
        fmt.Printf("%s: %v\n", key, value)
    }
    
    fmt.Println("\n=== Method Signatures ===")
    signatures := inspector.GetMethodSignatures()
    for _, sig := range signatures {
        fmt.Printf("- %s\n", sig)
    }
    
    fmt.Println("\n=== Field Types ===")
    fieldTypes := inspector.GetFieldTypes()
    for name, typ := range fieldTypes {
        fmt.Printf("%s: %s\n", name, typ)
    }
    
    fmt.Println("\n=== Interface Implementation ===")
    stringerType := reflect.TypeOf((*Stringer)(nil)).Elem()
    implementsStringer := inspector.ImplementsInterface(stringerType)
    fmt.Printf("Implements Stringer interface: %t\n", implementsStringer)
    
    fmt.Println("\n=== Method Call via Reflection ===")
    val := reflect.ValueOf(person)
    stringMethod := val.MethodByName("String")
    if stringMethod.IsValid() {
        results := stringMethod.Call(nil)
        fmt.Printf("String() result: %s\n", results[0].String())
    }
}