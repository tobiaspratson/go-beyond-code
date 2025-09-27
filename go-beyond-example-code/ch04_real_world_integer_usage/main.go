package main

import "fmt"

func main() {
    // E-commerce system example
    type Product struct {
        ID       uint32    // Product ID (up to 4 billion)
        Price    int64     // Price in cents (large numbers)
        Stock    uint16    // Stock quantity (up to 65,535)
        Category uint8     // Category ID (0-255)
        Rating   int8      // Rating from -5 to +5
    }
    
    // User system example
    type User struct {
        ID       uint64    // User ID (very large)
        Age      uint8     // Age (0-255)
        Score    int32     // User score (can be negative)
        Flags    uint8     // Bit flags for user status
    }
    
    // Network protocol example
    type Packet struct {
        Length   uint16    // Packet length (0-65,535 bytes)
        Type     uint8     // Packet type (0-255)
        Sequence uint32    // Sequence number
        Checksum uint16    // Checksum value
    }
    
    // Create examples
    product := Product{
        ID:       12345,
        Price:    1999,  // $19.99 in cents
        Stock:    100,
        Category: 1,
        Rating:   5,
    }
    
    user := User{
        ID:    987654321,
        Age:   25,
        Score: 1500,
        Flags: 0b1010,  // Binary flags
    }
    
    fmt.Printf("Product: ID=%d, Price=$%.2f, Stock=%d\n", 
        product.ID, float64(product.Price)/100, product.Stock)
    fmt.Printf("User: ID=%d, Age=%d, Score=%d\n", 
        user.ID, user.Age, user.Score)
}