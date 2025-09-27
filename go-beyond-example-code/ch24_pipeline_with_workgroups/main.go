package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Pipeline stage
type PipelineStage[T, U any] struct {
    Name     string
    Process  func(context.Context, T) (U, error)
    Workers  int
}

// Pipeline with workgroups
type Pipeline[T, U any] struct {
    stages []PipelineStage[T, U]
    ctx    context.Context
    cancel context.CancelFunc
}

func NewPipeline[T, U any](ctx context.Context) *Pipeline[T, U] {
    ctx, cancel := context.WithCancel(ctx)
    return &Pipeline[T, U]{
        ctx:    ctx,
        cancel: cancel,
    }
}

func (p *Pipeline[T, U]) AddStage(stage PipelineStage[T, U]) {
    p.stages = append(p.stages, stage)
}

func (p *Pipeline[T, U]) Process(inputs []T) ([]U, error) {
    var results []U
    var mu sync.Mutex
    
    for _, stage := range p.stages {
        var wg sync.WaitGroup
        var stageResults []U
        
        for i := 0; i < stage.Workers; i++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                for _, input := range inputs {
                    select {
                    case <-p.ctx.Done():
                        return
                    default:
                        result, err := stage.Process(p.ctx, input)
                        if err != nil {
                            fmt.Printf("Stage %s error: %v\n", stage.Name, err)
                            continue
                        }
                        
                        mu.Lock()
                        stageResults = append(stageResults, result)
                        mu.Unlock()
                    }
                }
            }()
        }
        
        wg.Wait()
        results = stageResults
        inputs = make([]T, len(results))
        for i, result := range results {
            inputs[i] = any(result).(T)
        }
    }
    
    return results, nil
}

func (p *Pipeline[T, U]) Cancel() {
    p.cancel()
}

func main() {
    // Create context
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Create pipeline
    pipeline := NewPipeline[int, int](ctx)
    
    // Add stages
    pipeline.AddStage(PipelineStage[int, int]{
        Name:    "Double",
        Process: func(ctx context.Context, x int) (int, error) {
            time.Sleep(100 * time.Millisecond) // Simulate work
            return x * 2, nil
        },
        Workers: 2,
    })
    
    pipeline.AddStage(PipelineStage[int, int]{
        Name:    "Add 10",
        Process: func(ctx context.Context, x int) (int, error) {
            time.Sleep(100 * time.Millisecond) // Simulate work
            return x + 10, nil
        },
        Workers: 2,
    })
    
    // Process inputs
    inputs := []int{1, 2, 3, 4, 5}
    results, err := pipeline.Process(inputs)
    
    if err != nil {
        fmt.Printf("Pipeline error: %v\n", err)
        return
    }
    
    fmt.Printf("Inputs: %v\n", inputs)
    fmt.Printf("Results: %v\n", results)
}