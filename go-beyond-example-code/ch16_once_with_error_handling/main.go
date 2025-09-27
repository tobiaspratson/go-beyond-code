package main

import (
    "fmt"
    "sync"
    "time"
)

type Database struct {
    connected bool
    url       string
}

var (
    db   *Database
    once sync.Once
    err  error
)

func GetDatabase() (*Database, error) {
    once.Do(func() {
        fmt.Println("Initializing database...")
        time.Sleep(200 * time.Millisecond)
        
        // Simulate connection
        db = &Database{
            connected: true,
            url:       "localhost:5432",
        }
        fmt.Println("Database connected!")
    })
    
    return db, err
}

func main() {
    var wg sync.WaitGroup
    
    // Multiple goroutines trying to get database
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            db, err := GetDatabase()
            if err != nil {
                fmt.Printf("Goroutine %d: Error: %v\n", id, err)
            } else {
                fmt.Printf("Goroutine %d: Database connected to %s\n", id, db.url)
            }
        }(i)
    }
    
    wg.Wait()
}