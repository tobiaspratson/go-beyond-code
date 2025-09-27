package main

import "fmt"

// HTTP client interface - represents an external HTTP service
type HTTPClient interface {
    Get(url string) (string, error)
    Post(url string, data string) error
}

// Mock HTTP client with comprehensive call tracking
type MockHTTPClient struct {
    getCalls    []GetCall    // Track all GET requests
    postCalls   []PostCall   // Track all POST requests
    getResponse  string       // Configurable GET response
    getError    error        // Configurable GET error
    postError   error        // Configurable POST error
}

// Data structures to capture call information
type GetCall struct {
    URL string
}

type PostCall struct {
    URL  string
    Data string
}

// Constructor for mock HTTP client
func NewMockHTTPClient() *MockHTTPClient {
    return &MockHTTPClient{
        getCalls:  make([]GetCall, 0),
        postCalls: make([]PostCall, 0),
    }
}

// Get method with call tracking
func (m *MockHTTPClient) Get(url string) (string, error) {
    // Record the call with timestamp and parameters
    m.getCalls = append(m.getCalls, GetCall{URL: url})
    
    // Return configured error if set
    if m.getError != nil {
        return "", m.getError
    }
    
    // Return configured response
    return m.getResponse, nil
}

// Post method with call tracking
func (m *MockHTTPClient) Post(url string, data string) error {
    // Record the call with all parameters
    m.postCalls = append(m.postCalls, PostCall{URL: url, Data: data})
    
    // Return configured error
    return m.postError
}

// Configuration methods for controlling mock behavior
func (m *MockHTTPClient) SetGetResponse(response string) {
    m.getResponse = response
}

func (m *MockHTTPClient) SetGetError(err error) {
    m.getError = err
}

func (m *MockHTTPClient) SetPostError(err error) {
    m.postError = err
}

// Verification methods - these are crucial for testing
func (m *MockHTTPClient) GetCalls() []GetCall {
    return m.getCalls
}

func (m *MockHTTPClient) PostCalls() []PostCall {
    return m.postCalls
}

// Advanced verification methods
func (m *MockHTTPClient) WasGetCalledWith(url string) bool {
    for _, call := range m.getCalls {
        if call.URL == url {
            return true
        }
    }
    return false
}

func (m *MockHTTPClient) WasPostCalledWith(url, data string) bool {
    for _, call := range m.postCalls {
        if call.URL == url && call.Data == data {
            return true
        }
    }
    return false
}

func (m *MockHTTPClient) GetCallCount() int {
    return len(m.getCalls)
}

func (m *MockHTTPClient) PostCallCount() int {
    return len(m.postCalls)
}

func (m *MockHTTPClient) Clear() {
    m.getCalls = make([]GetCall, 0)
    m.postCalls = make([]PostCall, 0)
    m.getResponse = ""
    m.getError = nil
    m.postError = nil
}

// API service that uses HTTP client - this is what we want to test
type APIService struct {
    client HTTPClient
}

func NewAPIService(client HTTPClient) *APIService {
    return &APIService{client: client}
}

func (s *APIService) FetchUserData(userID string) (string, error) {
    url := fmt.Sprintf("https://api.example.com/users/%s", userID)
    return s.client.Get(url)
}

func (s *APIService) UpdateUser(userID, data string) error {
    url := fmt.Sprintf("https://api.example.com/users/%s", userID)
    return s.client.Post(url, data)
}

func main() {
    // Create mock HTTP client
    mockClient := NewMockHTTPClient()
    apiService := NewAPIService(mockClient)
    
    fmt.Println("=== Testing API Service with Call Tracking ===")
    
    // Set up mock responses
    mockClient.SetGetResponse(`{"id": "123", "name": "Alice"}`)
    
    // Test GET request
    data, err := apiService.FetchUserData("123")
    if err != nil {
        fmt.Printf("Error fetching user: %v\n", err)
    } else {
        fmt.Printf("✓ User data: %s\n", data)
    }
    
    // Test POST request
    err = apiService.UpdateUser("123", `{"name": "Alice Smith"}`)
    if err != nil {
        fmt.Printf("Error updating user: %v\n", err)
    } else {
        fmt.Println("✓ User updated successfully")
    }
    
    // Verify calls were made correctly
    fmt.Println("\n=== Call Verification ===")
    
    getCalls := mockClient.GetCalls()
    postCalls := mockClient.PostCalls()
    
    fmt.Printf("GET calls made: %d\n", len(getCalls))
    for i, call := range getCalls {
        fmt.Printf("  %d. GET %s\n", i+1, call.URL)
    }
    
    fmt.Printf("POST calls made: %d\n", len(postCalls))
    for i, call := range postCalls {
        fmt.Printf("  %d. POST %s with data: %s\n", i+1, call.URL, call.Data)
    }
    
    // Advanced verification
    fmt.Println("\n=== Advanced Verification ===")
    
    if mockClient.WasGetCalledWith("https://api.example.com/users/123") {
        fmt.Println("✓ GET was called with correct URL")
    }
    
    if mockClient.WasPostCalledWith("https://api.example.com/users/123", `{"name": "Alice Smith"}`) {
        fmt.Println("✓ POST was called with correct URL and data")
    }
    
    fmt.Printf("Total API calls made: %d\n", mockClient.GetCallCount()+mockClient.PostCallCount())
}