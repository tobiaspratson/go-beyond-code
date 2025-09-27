package main

import "fmt"

func main() {
    map1 := map[string]int{
        "a": 1,
        "b": 2,
    }
    
    map2 := map[string]int{
        "b": 3,  // This will overwrite map1["b"]
        "c": 4,
    }
    
    // Method 1: Merge map2 into map1 (destructive)
    fmt.Printf("Before merge - map1: %v\n", map1)
    for k, v := range map2 {
        map1[k] = v
    }
    fmt.Printf("After merge - map1: %v\n", map1)
    
    // Method 2: Create new map with merge (non-destructive)
    map3 := map[string]int{
        "d": 5,
        "e": 6,
    }
    
    merged := mergeMaps(map1, map3)
    fmt.Printf("Merged new map: %v\n", merged)
    fmt.Printf("Original map1: %v (unchanged)\n", map1)
    
    // Method 3: Merge with conflict resolution
    map4 := map[string]int{
        "a": 10,  // Conflict with map1["a"]
        "f": 7,
    }
    
    resolved := mergeWithConflictResolution(map1, map4)
    fmt.Printf("Merged with conflict resolution: %v\n", resolved)
}

// Non-destructive merge
func mergeMaps(map1, map2 map[string]int) map[string]int {
    result := make(map[string]int)
    
    // Copy map1
    for k, v := range map1 {
        result[k] = v
    }
    
    // Add/overwrite with map2
    for k, v := range map2 {
        result[k] = v
    }
    
    return result
}

// Merge with custom conflict resolution (prefer map2 values)
func mergeWithConflictResolution(map1, map2 map[string]int) map[string]int {
    result := make(map[string]int)
    
    // Copy map1
    for k, v := range map1 {
        result[k] = v
    }
    
    // Overwrite with map2 (map2 takes precedence)
    for k, v := range map2 {
        result[k] = v
    }
    
    return result
}