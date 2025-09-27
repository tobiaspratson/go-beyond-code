package main

import "fmt"

// Cache interface - represents a caching system
type Cache interface {
    Get(key string) (string, bool)
    Set(key, value string)
    Delete(key string)
}

// Advanced mock cache with comprehensive behavior verification
type MockCache struct {
    data        map[string]string  // Cache storage
    getCalls    []GetCall         // Detailed get call tracking
    setCalls    []SetCall         // Detailed set call tracking
    deleteCalls []DeleteCall      // Detailed delete call tracking
    hitCount    int               // Cache hit counter
    missCount   int               // Cache miss counter
}

// Detailed call tracking structures
type GetCall struct {
    Key      string
    Timestamp int64  // Could be time.Now().UnixNano()
    Result   string
    Hit      bool
}

type SetCall struct {
    Key      string
    Value    string
    Timestamp int64
}

type DeleteCall struct {
    Key      string
    Timestamp int64
}

func NewMockCache() *MockCache {
    return &MockCache{
        data:        make(map[string]string),
        getCalls:    make([]GetCall, 0),
        setCalls:    make([]SetCall, 0),
        deleteCalls: make([]DeleteCall, 0),
    }
}

// Get method with detailed tracking
func (m *MockCache) Get(key string) (string, bool) {
    // Track the call
    call := GetCall{
        Key:       key,
        Timestamp: 1234567890, // In real code, use time.Now().UnixNano()
    }
    
    value, exists := m.data[key]
    call.Result = value
    call.Hit = exists
    
    m.getCalls = append(m.getCalls, call)
    
    if exists {
        m.hitCount++
    } else {
        m.missCount++
    }
    
    return value, exists
}

// Set method with detailed tracking
func (m *MockCache) Set(key, value string) {
    call := SetCall{
        Key:       key,
        Value:     value,
        Timestamp: 1234567890,
    }
    
    m.setCalls = append(m.setCalls, call)
    m.data[key] = value
}

// Delete method with detailed tracking
func (m *MockCache) Delete(key string) {
    call := DeleteCall{
        Key:       key,
        Timestamp: 1234567890,
    }
    
    m.deleteCalls = append(m.deleteCalls, call)
    delete(m.data, key)
}

// Comprehensive verification methods
func (m *MockCache) GetCalls() []GetCall {
    return m.getCalls
}

func (m *MockCache) SetCalls() []SetCall {
    return m.setCalls
}

func (m *MockCache) DeleteCalls() []DeleteCall {
    return m.deleteCalls
}

// Advanced verification methods
func (m *MockCache) WasGetCalled(key string) bool {
    for _, call := range m.getCalls {
        if call.Key == key {
            return true
        }
    }
    return false
}

func (m *MockCache) WasSetCalled(key, value string) bool {
    for _, call := range m.setCalls {
        if call.Key == key && call.Value == value {
            return true
        }
    }
    return false
}

func (m *MockCache) WasDeleteCalled(key string) bool {
    for _, call := range m.deleteCalls {
        if call.Key == key {
            return true
        }
    }
    return false
}

// Performance metrics
func (m *MockCache) GetHitCount() int {
    return m.hitCount
}

func (m *MockCache) GetMissCount() int {
    return m.missCount
}

func (m *MockCache) GetHitRate() float64 {
    total := m.hitCount + m.missCount
    if total == 0 {
        return 0.0
    }
    return float64(m.hitCount) / float64(total)
}

// Call frequency analysis
func (m *MockCache) GetCallFrequency(key string) int {
    count := 0
    for _, call := range m.getCalls {
        if call.Key == key {
            count++
        }
    }
    return count
}

func (m *MockCache) GetSetCallFrequency(key string) int {
    count := 0
    for _, call := range m.setCalls {
        if call.Key == key {
            count++
        }
    }
    return count
}

// Call order verification
func (m *MockCache) WasGetCalledBefore(key1, key2 string) bool {
    var key1Index, key2Index int = -1, -1
    
    for i, call := range m.getCalls {
        if call.Key == key1 && key1Index == -1 {
            key1Index = i
        }
        if call.Key == key2 && key2Index == -1 {
            key2Index = i
        }
    }
    
    return key1Index != -1 && key2Index != -1 && key1Index < key2Index
}

// Call count methods
func (m *MockCache) GetCallCount() int {
    return len(m.getCalls)
}

func (m *MockCache) SetCallCount() int {
    return len(m.setCalls)
}

func (m *MockCache) DeleteCallCount() int {
    return len(m.deleteCalls)
}

func (m *MockCache) TotalCallCount() int {
    return len(m.getCalls) + len(m.setCalls) + len(m.deleteCalls)
}

// Cache state inspection
func (m *MockCache) GetCacheSize() int {
    return len(m.data)
}

func (m *MockCache) GetCacheKeys() []string {
    keys := make([]string, 0, len(m.data))
    for key := range m.data {
        keys = append(keys, key)
    }
    return keys
}

func (m *MockCache) IsKeyInCache(key string) bool {
    _, exists := m.data[key]
    return exists
}

func (m *MockCache) Clear() {
    m.data = make(map[string]string)
    m.getCalls = make([]GetCall, 0)
    m.setCalls = make([]SetCall, 0)
    m.deleteCalls = make([]DeleteCall, 0)
    m.hitCount = 0
    m.missCount = 0
}

// Service that uses cache - demonstrates cache usage patterns
type DataService struct {
    cache Cache
}

func NewDataService(cache Cache) *DataService {
    return &DataService{cache: cache}
}

func (s *DataService) GetData(key string) (string, error) {
    // Try cache first (cache-aside pattern)
    if value, exists := s.cache.Get(key); exists {
        return value, nil
    }
    
    // Simulate expensive operation (database query, API call, etc.)
    value := fmt.Sprintf("expensive-data-for-%s", key)
    
    // Cache the result for future use
    s.cache.Set(key, value)
    
    return value, nil
}

func (s *DataService) InvalidateData(key string) {
    s.cache.Delete(key)
}

func (s *DataService) GetDataWithFallback(key string) (string, error) {
    // Try cache first
    if value, exists := s.cache.Get(key); exists {
        return value, nil
    }
    
    // Try fallback data source
    fallbackValue := fmt.Sprintf("fallback-data-for-%s", key)
    
    // Cache the fallback result
    s.cache.Set(key, fallbackValue)
    
    return fallbackValue, nil
}

func main() {
    // Create advanced mock cache
    mockCache := NewMockCache()
    dataService := NewDataService(mockCache)
    
    fmt.Println("=== Testing Cache Behavior ===")
    
    // Test cache miss scenario
    fmt.Println("\n--- Cache Miss Test ---")
    data, err := dataService.GetData("user:123")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("✓ Retrieved data: %s\n", data)
    }
    
    // Test cache hit scenario
    fmt.Println("\n--- Cache Hit Test ---")
    data, err = dataService.GetData("user:123")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("✓ Retrieved cached data: %s\n", data)
    }
    
    // Test cache invalidation
    fmt.Println("\n--- Cache Invalidation Test ---")
    dataService.InvalidateData("user:123")
    data, err = dataService.GetData("user:123")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("✓ Retrieved fresh data after invalidation: %s\n", data)
    }
    
    // Test multiple keys
    fmt.Println("\n--- Multiple Keys Test ---")
    dataService.GetData("user:456")
    dataService.GetData("user:789")
    dataService.GetData("user:456") // Should hit cache
    
    // Comprehensive behavior verification
    fmt.Println("\n=== Behavior Verification ===")
    
    // Basic call counts
    fmt.Printf("Total GET calls: %d\n", mockCache.GetCallCount())
    fmt.Printf("Total SET calls: %d\n", mockCache.SetCallCount())
    fmt.Printf("Total DELETE calls: %d\n", mockCache.DeleteCallCount())
    fmt.Printf("Total operations: %d\n", mockCache.TotalCallCount())
    
    // Cache performance metrics
    fmt.Printf("Cache hits: %d\n", mockCache.GetHitCount())
    fmt.Printf("Cache misses: %d\n", mockCache.GetMissCount())
    fmt.Printf("Hit rate: %.2f%%\n", mockCache.GetHitRate()*100)
    
    // Specific key analysis
    fmt.Printf("user:123 GET calls: %d\n", mockCache.GetCallFrequency("user:123"))
    fmt.Printf("user:456 GET calls: %d\n", mockCache.GetCallFrequency("user:456"))
    
    // Cache state
    fmt.Printf("Cache size: %d\n", mockCache.GetCacheSize())
    fmt.Printf("Cached keys: %v\n", mockCache.GetCacheKeys())
    
    // Advanced verifications
    if mockCache.WasGetCalled("user:123") {
        fmt.Println("✓ user:123 was queried")
    }
    
    if mockCache.WasSetCalled("user:123", "expensive-data-for-user:123") {
        fmt.Println("✓ user:123 was cached with correct value")
    }
    
    if mockCache.WasDeleteCalled("user:123") {
        fmt.Println("✓ user:123 was invalidated")
    }
    
    if mockCache.WasGetCalledBefore("user:123", "user:456") {
        fmt.Println("✓ user:123 was queried before user:456")
    }
    
    // Cache efficiency analysis
    fmt.Println("\n=== Cache Efficiency Analysis ===")
    for _, key := range mockCache.GetCacheKeys() {
        frequency := mockCache.GetCallFrequency(key)
        setFrequency := mockCache.GetSetCallFrequency(key)
        fmt.Printf("Key %s: %d gets, %d sets\n", key, frequency, setFrequency)
    }
}