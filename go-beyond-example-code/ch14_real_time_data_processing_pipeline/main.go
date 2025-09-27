package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type DataPoint struct {
    Timestamp time.Time
    Value     float64
    Source    string
}

type Processor struct {
    InputChan  <-chan DataPoint
    OutputChan chan<- DataPoint
    QuitChan   <-chan bool
}

func NewProcessor(input <-chan DataPoint, output chan<- DataPoint, quit <-chan bool) *Processor {
    return &Processor{
        InputChan:  input,
        OutputChan: output,
        QuitChan:   quit,
    }
}

func (p *Processor) Start() {
    go func() {
        for {
            select {
            case data := <-p.InputChan:
                // Process data (e.g., filter, transform, aggregate)
                processed := DataPoint{
                    Timestamp: data.Timestamp,
                    Value:     data.Value * 1.1, // Simple transformation
                    Source:    data.Source,
                }
                p.OutputChan <- processed
            case <-p.QuitChan:
                fmt.Println("Processor shutting down")
                return
            }
        }
    }()
}

func dataGenerator(output chan<- DataPoint, quit <-chan bool) {
    sources := []string{"sensor1", "sensor2", "sensor3"}
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            data := DataPoint{
                Timestamp: time.Now(),
                Value:     rand.Float64() * 100,
                Source:    sources[rand.Intn(len(sources))],
            }
            output <- data
        case <-quit:
            fmt.Println("Data generator stopping")
            return
        }
    }
}

func dataConsumer(input <-chan DataPoint, quit <-chan bool) {
    for {
        select {
        case data := <-input:
            fmt.Printf("Processed: %s at %v: %.2f\n", 
                data.Source, data.Timestamp.Format("15:04:05"), data.Value)
        case <-quit:
            fmt.Println("Consumer stopping")
            return
        }
    }
}

func main() {
    rawData := make(chan DataPoint, 10)
    processedData := make(chan DataPoint, 10)
    quit := make(chan bool)
    
    // Start data generator
    go dataGenerator(rawData, quit)
    
    // Start processor
    processor := NewProcessor(rawData, processedData, quit)
    processor.Start()
    
    // Start consumer
    go dataConsumer(processedData, quit)
    
    // Run for 5 seconds
    time.Sleep(5 * time.Second)
    
    // Shutdown
    close(quit)
    time.Sleep(100 * time.Millisecond)
}