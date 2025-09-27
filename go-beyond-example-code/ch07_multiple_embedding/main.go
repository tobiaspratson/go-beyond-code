package main

import "fmt"

// Base structs for different capabilities
type Logger struct {
    Level string
}

func (l Logger) Log(message string) {
    fmt.Printf("[%s] %s\n", l.Level, message)
}

type Metrics struct {
    Count int
}

func (m *Metrics) Increment() {
    m.Count++
}

func (m Metrics) GetCount() int {
    return m.Count
}

type Config struct {
    Name string
}

func (c Config) GetName() string {
    return c.Name
}

// Service that embeds multiple capabilities
type Service struct {
    Logger   // Embed Logger
    Metrics  // Embed Metrics
    Config   // Embed Config
    Status   string
}

func (s *Service) Start() {
    s.Log("Service starting...")
    s.Increment()
    s.Log(fmt.Sprintf("Service %s started", s.GetName()))
}

func (s *Service) Stop() {
    s.Log("Service stopping...")
    s.Increment()
    s.Log(fmt.Sprintf("Service %s stopped", s.GetName()))
}

func (s Service) GetStatus() string {
    return fmt.Sprintf("Service: %s, Status: %s, Count: %d", 
        s.GetName(), s.Status, s.GetCount())
}

func main() {
    service := Service{
        Logger: Logger{Level: "INFO"},
        Metrics: Metrics{Count: 0},
        Config: Config{Name: "MyService"},
        Status: "Running",
    }
    
    fmt.Println("=== SERVICE OPERATIONS ===")
    service.Start()
    service.Stop()
    
    fmt.Printf("Final status: %s\n", service.GetStatus())
    
    // Access embedded fields directly
    fmt.Printf("Log level: %s\n", service.Level)
    fmt.Printf("Service name: %s\n", service.Name)
    fmt.Printf("Metrics count: %d\n", service.Count)
}