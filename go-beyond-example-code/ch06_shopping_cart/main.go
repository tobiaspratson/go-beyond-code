package main

import (
    "fmt"
    "sort"
)

type Item struct {
    Name     string
    Price    float64
    Quantity int
    Category string
}

type ShoppingCart struct {
    Items []Item
}

func (c *ShoppingCart) AddItem(item Item) {
    // Check if item already exists
    for i, existingItem := range c.Items {
        if existingItem.Name == item.Name {
            c.Items[i].Quantity += item.Quantity
            return
        }
    }
    c.Items = append(c.Items, item)
}

func (c *ShoppingCart) RemoveItem(name string) bool {
    for i, item := range c.Items {
        if item.Name == name {
            c.Items = append(c.Items[:i], c.Items[i+1:]...)
            return true
        }
    }
    return false
}

func (c *ShoppingCart) UpdateQuantity(name string, quantity int) bool {
    for i, item := range c.Items {
        if item.Name == name {
            if quantity <= 0 {
                c.RemoveItem(name)
            } else {
                c.Items[i].Quantity = quantity
            }
            return true
        }
    }
    return false
}

func (c *ShoppingCart) GetTotal() float64 {
    total := 0.0
    for _, item := range c.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

func (c *ShoppingCart) GetItemCount() int {
    count := 0
    for _, item := range c.Items {
        count += item.Quantity
    }
    return count
}

func (c *ShoppingCart) GetItemsByCategory(category string) []Item {
    var filtered []Item
    for _, item := range c.Items {
        if item.Category == category {
            filtered = append(filtered, item)
        }
    }
    return filtered
}

func (c *ShoppingCart) SortByPrice() {
    sort.Slice(c.Items, func(i, j int) bool {
        return c.Items[i].Price < c.Items[j].Price
    })
}

func (c *ShoppingCart) SortByName() {
    sort.Slice(c.Items, func(i, j int) bool {
        return c.Items[i].Name < c.Items[j].Name
    })
}

func (c *ShoppingCart) Display() {
    if len(c.Items) == 0 {
        fmt.Println("Cart is empty!")
        return
    }
    
    fmt.Println("=== Shopping Cart ===")
    total := 0.0
    for i, item := range c.Items {
        itemTotal := item.Price * float64(item.Quantity)
        fmt.Printf("%d. %s (%s) - $%.2f x %d = $%.2f\n", 
            i+1, item.Name, item.Category, item.Price, item.Quantity, itemTotal)
        total += itemTotal
    }
    fmt.Printf("Total: $%.2f (%d items)\n", total, c.GetItemCount())
}

func main() {
    cart := ShoppingCart{}
    
    // Add items to cart
    cart.AddItem(Item{"Apple", 1.50, 3, "Fruits"})
    cart.AddItem(Item{"Banana", 0.75, 5, "Fruits"})
    cart.AddItem(Item{"Orange", 1.25, 2, "Fruits"})
    cart.AddItem(Item{"Milk", 3.50, 1, "Dairy"})
    cart.AddItem(Item{"Bread", 2.00, 2, "Bakery"})
    
    // Display cart
    cart.Display()
    
    // Add more of existing item
    fmt.Println("\nAdding more apples...")
    cart.AddItem(Item{"Apple", 1.50, 2, "Fruits"})
    cart.Display()
    
    // Update quantity
    fmt.Println("\nUpdating banana quantity...")
    cart.UpdateQuantity("Banana", 3)
    cart.Display()
    
    // Remove an item
    fmt.Println("\nRemoving milk...")
    cart.RemoveItem("Milk")
    cart.Display()
    
    // Sort by price
    fmt.Println("\nSorting by price...")
    cart.SortByPrice()
    cart.Display()
    
    // Get items by category
    fmt.Println("\nFruits in cart:")
    fruits := cart.GetItemsByCategory("Fruits")
    for _, fruit := range fruits {
        fmt.Printf("- %s: $%.2f x %d\n", fruit.Name, fruit.Price, fruit.Quantity)
    }
}