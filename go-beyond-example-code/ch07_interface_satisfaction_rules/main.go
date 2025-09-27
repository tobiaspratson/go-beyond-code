package main

import "fmt"

// Interface with multiple methods
type Animal interface {
    Speak() string
    Move() string
    Eat() string
}

// Dog implements Animal
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running"
}

func (d Dog) Eat() string {
    return "Dog food"
}

// Cat implements Animal
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Sneaking"
}

func (c Cat) Eat() string {
    return "Cat food"
}

// Function that works with any Animal
func animalInfo(a Animal) {
    fmt.Printf("Animal says: %s\n", a.Speak())
    fmt.Printf("Animal moves: %s\n", a.Move())
    fmt.Printf("Animal eats: %s\n", a.Eat())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    // Both implement Animal interface
    animalInfo(dog)
    fmt.Println()
    animalInfo(cat)
}