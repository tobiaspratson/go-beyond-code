package main

import "fmt"

// Base struct
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return fmt.Sprintf("%s makes a sound", a.Name)
}

func (a Animal) Move() string {
    return fmt.Sprintf("%s moves", a.Name)
}

// Embedded struct with method override
type Dog struct {
    Animal
    Breed string
}

// Override the Speak method
func (d Dog) Speak() string {
    return fmt.Sprintf("%s barks (Breed: %s)", d.Name, d.Breed)
}

// Add new method
func (d Dog) Fetch() string {
    return fmt.Sprintf("%s fetches the ball", d.Name)
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    
    // Call overridden method
    fmt.Println(dog.Speak())
    
    // Call inherited method
    fmt.Println(dog.Move())
    
    // Call new method
    fmt.Println(dog.Fetch())
}