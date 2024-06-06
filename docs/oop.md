# OOP in Go

1. Encapsulation
Encapsulation in Go is achieved using structs and methods. Visibility is controlled through capitalized names for exported identifiers and lowercase names for unexported identifiers.

```go
package main

import (
    "fmt"
)

// Struct definition
type Person struct {
    firstName string
    lastName  string
}

// Method to access firstName (getter)
func (p *Person) FirstName() string {
    return p.firstName
}

// Method to set firstName (setter)
func (p *Person) SetFirstName(firstName string) {
    p.firstName = firstName
}

func main() {
    p := Person{firstName: "John", lastName: "Doe"}
    fmt.Println(p.FirstName()) // Accessing firstName using getter
    p.SetFirstName("Jane")     // Modifying firstName using setter
    fmt.Println(p.FirstName())
}
```

2. Abstraction
Abstraction in Go is achieved using interfaces. Interfaces define methods that must be implemented by the types.

```go
package main

import "fmt"

// Define an interface
type Animal interface {
    Speak() string
}

// Implement the interface in a struct
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}}

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

3. Polymorphism
Polymorphism in Go is achieved through interfaces. Different types can implement the same interface, and the interface can be used to refer to any of these types.

```go
package main

import "fmt"

// Interface definition
type Shape interface {
    Area() float64
}

// Struct for Rectangle
type Rectangle struct {
    Width, Height float64
}

// Method to calculate area of Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Struct for Circle
type Circle struct {
    Radius float64
}

// Method to calculate area of Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 5, Height: 10},
        Circle{Radius: 7},
    }

    for _, shape := range shapes {
        fmt.Printf("Area: %f\n", shape.Area())
    }
}
```

4. Composition (Inheritance)
Go favors composition over inheritance. Instead of inheriting from a base class, you compose structs to build complex types.

```go
package main

import "fmt"

// Basic struct
type Address struct {
    City, State string
}

// Another struct that includes Address (composition)
type Person struct {
    FirstName, LastName string
    Address             // Embedded struct
}

func main() {
    p := Person{
        FirstName: "John",
        LastName:  "Doe",
        Address: Address{
            City:  "New York",
            State: "NY",
        },
    }

    fmt.Printf("Name: %s %s\n", p.FirstName, p.LastName)
    fmt.Printf("Address: %s, %s\n", p.City, p.State) // Direct access to Address fields
}
```

- Encapsulation: Achieved using structs and methods.
- Abstraction: Achieved using interfaces.
- Polymorphism: Achieved through interfaces, allowing different types to be treated uniformly.
- Composition: Preferred over inheritance; achieved by embedding structs within other structs.