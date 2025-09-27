package main

import (
    "context"
    "fmt"
    "time"
)

type UserService struct {
    name string
}

func (us *UserService) GetUser(ctx context.Context, userID string) (string, error) {
    fmt.Printf("%s: Getting user %s\n", us.name, userID)
    
    select {
    case <-time.After(300 * time.Millisecond):
        return fmt.Sprintf("User %s from %s", userID, us.name), nil
    case <-ctx.Done():
        return "", fmt.Errorf("%s cancelled: %v", us.name, ctx.Err())
    }
}

func (us *UserService) UpdateUser(ctx context.Context, userID string, data string) error {
    fmt.Printf("%s: Updating user %s with %s\n", us.name, userID, data)
    
    select {
    case <-time.After(200 * time.Millisecond):
        fmt.Printf("%s: User %s updated\n", us.name, userID)
        return nil
    case <-ctx.Done():
        return fmt.Errorf("%s update cancelled: %v", us.name, ctx.Err())
    }
}

func orchestrateUserUpdate(ctx context.Context, userID string, data string) error {
    fmt.Println("Orchestrating user update")
    
    // Create user service
    userService := &UserService{name: "UserService"}
    
    // Get user first
    user, err := userService.GetUser(ctx, userID)
    if err != nil {
        return fmt.Errorf("failed to get user: %v", err)
    }
    fmt.Printf("Retrieved: %s\n", user)
    
    // Update user
    err = userService.UpdateUser(ctx, userID, data)
    if err != nil {
        return fmt.Errorf("failed to update user: %v", err)
    }
    
    return nil
}

func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    // Orchestrate user update
    err := orchestrateUserUpdate(ctx, "12345", "new profile data")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Println("User update completed successfully")
}