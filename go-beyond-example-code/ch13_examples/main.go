package main

import (
    "fmt"
    "sync"
    "time"
)

type Scraper struct {
    urls []string
    results chan string
    wg sync.WaitGroup
}

func NewScraper(urls []string) *Scraper {
    return &Scraper{
        urls: urls,
        results: make(chan string, len(urls)),
    }
}

func (s *Scraper) scrapeURL(url string) {
    defer s.wg.Done()
    
    // Simulate web scraping
    time.Sleep(500 * time.Millisecond)
    
    result := fmt.Sprintf("Scraped: %s", url)
    s.results <- result
}

func (s *Scraper) ScrapeAll() []string {
    // Start scraping all URLs concurrently
    for _, url := range s.urls {
        s.wg.Add(1)
        go s.scrapeURL(url)
    }
    
    // Wait for all scraping to complete
    s.wg.Wait()
    close(s.results)
    
    // Collect results
    var results []string
    for result := range s.results {
        results = append(results, result)
    }
    
    return results
}

func main() {
    urls := []string{
        "https://example.com",
        "https://google.com",
        "https://github.com",
        "https://stackoverflow.com",
    }
    
    scraper := NewScraper(urls)
    
    start := time.Now()
    results := scraper.ScrapeAll()
    elapsed := time.Since(start)
    
    fmt.Printf("Scraped %d URLs in %v\n", len(results), elapsed)
    for _, result := range results {
        fmt.Println(result)
    }
}