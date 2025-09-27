package main

import "fmt"

// Base struct - contains common animal properties
type Animal struct {
    Name string
    Age  int
}

// Base methods that all animals have
func (a Animal) Speak() {
    fmt.Printf("%s makes a sound\n", a.Name)
}

func (a Animal) GetInfo() string {
    return fmt.Sprintf("%s is %d years old", a.Name, a.Age)
}

func (a Animal) Move() string {
    return fmt.Sprintf("%s moves around", a.Name)
}

// Embedded struct - Dog includes Animal
type Dog struct {
    Animal  // Embedding - Dog "has-a" Animal
    Breed   string
}

// Override the Speak method for dogs
func (d Dog) Speak() {
    fmt.Printf("%s barks: Woof! Woof!\n", d.Name)
}

// Add new method specific to dogs
func (d Dog) GetBreed() string {
    return d.Breed
}

// Add another method specific to dogs
func (d Dog) Fetch() string {
    return fmt.Sprintf("%s fetches the ball", d.Name)
}

// Another embedded struct - Cat includes Animal
type Cat struct {
    Animal  // Embedding - Cat "has-a" Animal
    Color   string
}

// Override the Speak method for cats
func (c Cat) Speak() {
    fmt.Printf("%s meows: Meow!\n", c.Name)
}

// Add new method specific to cats
func (c Cat) GetColor() string {
    return c.Color
}

// Add another method specific to cats
func (c Cat) Purr() string {
    return fmt.Sprintf("%s purrs softly", c.Name)
}

func main() {
    // Create a dog
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }
    
    // Create a cat
    cat := Cat{
        Animal: Animal{Name: "Whiskers", Age: 2},
        Color:  "Orange",
    }
    
    fmt.Println("=== DOG ===")
    // Access embedded fields directly (promoted fields)
    fmt.Printf("Name: %s\n", dog.Name)        // From Animal
    fmt.Printf("Age: %d\n", dog.Age)          // From Animal
    fmt.Printf("Breed: %s\n", dog.Breed)      // From Dog
    
    // Call methods
    dog.Speak()                               // Dog's Speak method (overridden)
    fmt.Println(dog.GetInfo())                // Animal's GetInfo method (inherited)
    fmt.Println(dog.Move())                   // Animal's Move method (inherited)
    fmt.Println(dog.GetBreed())               // Dog's GetBreed method (new)
    fmt.Println(dog.Fetch())                  // Dog's Fetch method (new)
    
    fmt.Println("\n=== CAT ===")
    // Access embedded fields directly
    fmt.Printf("Name: %s\n", cat.Name)        // From Animal
    fmt.Printf("Age: %d\n", cat.Age)             // From Animal
    fmt.Printf("Color: %s\n", cat.Color)       // From Cat
    
    // Call methods
    cat.Speak()                                // Cat's Speak method (overridden)
    fmt.Println(cat.GetInfo())                 // Animal's GetInfo method (inherited)
    fmt.Println(cat.Move())                    // Animal's Move method (inherited)
    fmt.Println(cat.GetColor())                // Cat's GetColor method (new)
    fmt.Println(cat.Purr())                    // Cat's Purr method (new)
}