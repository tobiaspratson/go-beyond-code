package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Configuration parsing
    type Config struct {
        Port        int
        Timeout     float64
        Debug       bool
        MaxUsers    int64
    }
    
    // Simulate reading from environment or config file
    portStr := "8080"
    timeoutStr := "30.5"
    debugStr := "true"
    maxUsersStr := "1000000"
    
    config := Config{}
    
    // Parse configuration with error handling
    if port, err := strconv.Atoi(portStr); err == nil {
        config.Port = port
    } else {
        fmt.Printf("Invalid port: %v\n", err)
    }
    
    if timeout, err := strconv.ParseFloat(timeoutStr, 64); err == nil {
        config.Timeout = timeout
    } else {
        fmt.Printf("Invalid timeout: %v\n", err)
    }
    
    if debug, err := strconv.ParseBool(debugStr); err == nil {
        config.Debug = debug
    } else {
        fmt.Printf("Invalid debug flag: %v\n", err)
    }
    
    if maxUsers, err := strconv.ParseInt(maxUsersStr, 10, 64); err == nil {
        config.MaxUsers = maxUsers
    } else {
        fmt.Printf("Invalid max users: %v\n", err)
    }
    
    fmt.Printf("Config: %+v\n", config)
    
    // Data validation and conversion
    func validateAndConvert(input string) (int, error) {
        // Check if string is numeric
        if _, err := strconv.Atoi(input); err != nil {
            return 0, fmt.Errorf("input '%s' is not a valid integer", input)
        }
        
        // Convert and validate range
        val, _ := strconv.Atoi(input)
        if val < 0 || val > 1000 {
            return 0, fmt.Errorf("value %d is out of range (0-1000)", val)
        }
        
        return val, nil
    }
    
    // Test validation
    testInputs := []string{"42", "999", "1001", "abc", "-5"}
    
    for _, input := range testInputs {
        if val, err := validateAndConvert(input); err == nil {
            fmt.Printf("'%s' -> %d (valid)\n", input, val)
        } else {
            fmt.Printf("'%s' -> Error: %v\n", input, err)
        }
    }
}