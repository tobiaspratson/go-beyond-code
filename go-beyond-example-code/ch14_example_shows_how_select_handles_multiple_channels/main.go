package main

import (
    "fmt"
    "time"
)

func main() {
    highPriority := make(chan string)
    normalPriority := make(chan string)
    lowPriority := make(chan string)
    
    // Start goroutines with different priorities
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(200 * time.Millisecond)
            highPriority <- fmt.Sprintf("HIGH: Message %d", i)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(300 * time.Millisecond)
            normalPriority <- fmt.Sprintf("NORMAL: Message %d", i)
        }
    }()
    
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(500 * time.Millisecond)
            lowPriority <- fmt.Sprintf("LOW: Message %d", i)
        }
    }()
    
    // Process messages with priority handling
    for i := 0; i < 9; i++ {
        select {
        case msg := <-highPriority:
            fmt.Printf("🔥 %s\n", msg)
        case msg := <-normalPriority:
            fmt.Printf("📝 %s\n", msg)
        case msg := <-lowPriority:
            fmt.Printf("📄 %s\n", msg)
        }
    }
}