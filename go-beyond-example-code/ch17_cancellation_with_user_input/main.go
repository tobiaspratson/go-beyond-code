package main

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strings"
    "time"
)

func dataProcessor(ctx context.Context, id int) {
    for i := 0; i < 100; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Processor %d stopped: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Processor %d: processing item %d\n", id, i)
            time.Sleep(100 * time.Millisecond)
        }
    }
    fmt.Printf("Processor %d completed all items\n", id)
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // Start data processors
    for i := 1; i <= 3; i++ {
        go dataProcessor(ctx, i)
    }
    
    // Wait for user input to cancel
    fmt.Println("Press Enter to cancel all processors...")
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    
    fmt.Println("Cancelling all processors...")
    cancel()
    
    // Give processors time to clean up
    time.Sleep(500 * time.Millisecond)
    fmt.Println("All processors stopped")
}