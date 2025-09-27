package main

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type Config struct {
    DatabaseHost string `config:"db_host"`
    DatabasePort int    `config:"db_port"`
    DebugMode    bool   `config:"debug"`
    MaxConnections int  `config:"max_conn"`
}

func loadConfig(config interface{}, envVars map[string]string) error {
    val := reflect.ValueOf(config)
    if val.Kind() != reflect.Ptr {
        return fmt.Errorf("config must be a pointer")
    }
    
    elem := val.Elem()
    t := elem.Type()
    
    for i := 0; i < elem.NumField(); i++ {
        field := elem.Field(i)
        fieldType := t.Field(i)
        
        // Get config tag
        configTag := fieldType.Tag.Get("config")
        if configTag == "" {
            continue
        }
        
        // Get value from environment
        value, exists := envVars[configTag]
        if !exists {
            continue
        }
        
        // Set the field value
        if !field.CanSet() {
            continue
        }
        
        switch field.Kind() {
        case reflect.String:
            field.SetString(value)
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
                field.SetInt(intVal)
            }
        case reflect.Bool:
            if boolVal, err := strconv.ParseBool(value); err == nil {
                field.SetBool(boolVal)
            }
        }
    }
    
    return nil
}

func main() {
    // Simulate environment variables
    envVars := map[string]string{
        "db_host":   "localhost",
        "db_port":   "5432",
        "debug":     "true",
        "max_conn":  "100",
    }
    
    config := &Config{}
    err := loadConfig(config, envVars)
    if err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        return
    }
    
    fmt.Printf("Config loaded: %+v\n", config)
}