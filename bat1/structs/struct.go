package main

import "fmt"

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	height, width float32
}

func (r Rectangle) Area() float32 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.height + r.width)
}

const PI = 3.14159

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return PI * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * PI * c.radius
}

func main() {
	var one Shape = Rectangle{height: 5, width: 3}
	var two Shape = Circle{radius: 4}
	fmt.Println("area of shape one: ", one.Area(), " perimeter of shape one: ", one.Perimeter())
	fmt.Println("area of shape two: ", two.Area(), " perimeter of shape two: ", two.Perimeter())
}

// END OMIT
