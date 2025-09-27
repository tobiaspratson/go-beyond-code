package main

import (
    "fmt"
    "time"
)

type ChannelHealth struct {
    ID       int
    LastSeen time.Time
    Healthy  bool
}

func channelHealthMonitor(channels []chan string, healthChan chan<- ChannelHealth, done <-chan bool) {
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()
    
    lastSeen := make(map[int]time.Time)
    
    for {
        select {
        case <-done:
            fmt.Println("Health monitor stopping")
            return
        case <-ticker.C:
            // Check each channel's health
            for i, ch := range channels {
                select {
                case msg := <-ch:
                    lastSeen[i] = time.Now()
                    healthChan <- ChannelHealth{
                        ID:       i,
                        LastSeen:  time.Now(),
                        Healthy:  true,
                    }
                    fmt.Printf("Channel %d: %s\n", i, msg)
                default:
                    // Channel is not ready, check if it's been too long
                    if lastSeenTime, exists := lastSeen[i]; exists {
                        if time.Since(lastSeenTime) > 2*time.Second {
                            healthChan <- ChannelHealth{
                                ID:       i,
                                LastSeen:  lastSeenTime,
                                Healthy:   false,
                            }
                        }
                    }
                }
            }
        }
    }
}

func main() {
    channels := make([]chan string, 3)
    for i := range channels {
        channels[i] = make(chan string, 2)
    }
    
    healthChan := make(chan ChannelHealth, 10)
    done := make(chan bool)
    
    // Start health monitor
    go channelHealthMonitor(channels, healthChan, done)
    
    // Start data producers
    for i, ch := range channels {
        go func(id int, ch chan<- string) {
            for j := 1; j <= 5; j++ {
                ch <- fmt.Sprintf("Message %d from channel %d", j, id)
                time.Sleep(time.Duration(id*200+300) * time.Millisecond)
            }
        }(i, ch)
    }
    
    // Monitor health
    go func() {
        for health := range healthChan {
            if health.Healthy {
                fmt.Printf("✅ Channel %d is healthy (last seen: %v)\n", 
                    health.ID, health.LastSeen.Format("15:04:05"))
            } else {
                fmt.Printf("❌ Channel %d is unhealthy (last seen: %v)\n", 
                    health.ID, health.LastSeen.Format("15:04:05"))
            }
        }
    }()
    
    time.Sleep(5 * time.Second)
    done <- true
    time.Sleep(100 * time.Millisecond)
}