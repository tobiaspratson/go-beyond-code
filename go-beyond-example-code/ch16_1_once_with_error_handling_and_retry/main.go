package main

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

type Service struct {
    name string
    once sync.Once
    err  error
}

func NewService(name string) *Service {
    return &Service{name: name}
}

func (s *Service) Initialize() error {
    s.once.Do(func() {
        fmt.Printf("Initializing service %s...\n", s.name)
        time.Sleep(100 * time.Millisecond)
        
        // Simulate potential failure
        if s.name == "failing-service" {
            s.err = errors.New("service initialization failed")
            return
        }
        
        fmt.Printf("Service %s initialized successfully\n", s.name)
    })
    
    return s.err
}

func main() {
    services := []*Service{
        NewService("database"),
        NewService("cache"),
        NewService("failing-service"),
        NewService("api"),
    }
    
    var wg sync.WaitGroup
    
    for _, service := range services {
        wg.Add(1)
        go func(s *Service) {
            defer wg.Done()
            
            if err := s.Initialize(); err != nil {
                fmt.Printf("Service %s failed: %v\n", s.name, err)
            } else {
                fmt.Printf("Service %s is ready\n", s.name)
            }
        }(service)
    }
    
    wg.Wait()
}