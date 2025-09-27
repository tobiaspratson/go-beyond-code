package main

import "fmt"

// Constraint interfaces
type Comparable interface {
    Compare(other Comparable) int
}

type Serializable interface {
    Serialize() []byte
    Deserialize(data []byte) error
}

// Combined interface
type ComparableSerializable interface {
    Comparable
    Serializable
}

// Implementation
type Version struct {
    Major, Minor, Patch int
}

func (v Version) Compare(other Comparable) int {
    otherVersion := other.(Version)
    
    if v.Major != otherVersion.Major {
        return v.Major - otherVersion.Major
    }
    if v.Minor != otherVersion.Minor {
        return v.Minor - otherVersion.Minor
    }
    return v.Patch - otherVersion.Patch
}

func (v Version) Serialize() []byte {
    return []byte(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
}

func (v *Version) Deserialize(data []byte) error {
    _, err := fmt.Sscanf(string(data), "%d.%d.%d", &v.Major, &v.Minor, &v.Patch)
    return err
}

func (v Version) String() string {
    return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func main() {
    v1 := Version{Major: 1, Minor: 2, Patch: 3}
    v2 := Version{Major: 1, Minor: 2, Patch: 4}
    
    // Use as Comparable
    var comp Comparable = v1
    result := comp.Compare(v2)
    fmt.Printf("Comparison result: %d\n", result)
    
    // Use as Serializable
    var ser Serializable = &v1
    data := ser.Serialize()
    fmt.Printf("Serialized: %s\n", string(data))
    
    // Use as combined interface
    var cs ComparableSerializable = &v1
    fmt.Printf("Version: %s\n", v1.String())
    fmt.Printf("Serialized: %s\n", string(cs.Serialize()))
}