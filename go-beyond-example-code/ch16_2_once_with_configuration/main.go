package main

import (
    "fmt"
    "sync"
    "time"
)

type Config struct {
    DatabaseURL string
    APIKey      string
    Debug       bool
}

var (
    config *Config
    once   sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        fmt.Println("Loading configuration...")
        time.Sleep(100 * time.Millisecond)
        
        config = &Config{
            DatabaseURL: "localhost:5432",
            APIKey:      "secret-key-123",
            Debug:       true,
        }
        fmt.Println("Configuration loaded")
    })
    
    return config
}

func main() {
    var wg sync.WaitGroup
    
    // Multiple goroutines accessing config
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            
            cfg := GetConfig()
            fmt.Printf("Goroutine %d: Database URL: %s\n", id, cfg.DatabaseURL)
        }(i)
    }
    
    wg.Wait()
}