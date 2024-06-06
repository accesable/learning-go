package main

import "fmt"

// Define an interface with some methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct that implements the Shape interface
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method (value receiver)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method (value receiver)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Implement the Resize method (without pointer receiver)
func (r Rectangle) ResizeWithoutPointer(newWidth, newHeight float64) {
	r.Width = newWidth
	r.Height = newHeight
}

// Implement the Resize method (without pointer receiver)
func (r *Rectangle) ResizeWithPointer(newWidth, newHeight float64) {
	r.Width = newWidth
	r.Height = newHeight
}

// Function that takes a Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	// Pass the rectangle by value to the function
	printShapeInfo(rect)

	// Try to modify the rectangle (this won't work because Resize has a pointer receiver)
	rect.ResizeWithPointer(20, 10)
	printShapeInfo(rect)

}
