package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (circ *Circle) Area() float64 {
	return math.Pi * circ.Radius * circ.Radius
}

func (rect *Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Shape interface {
	Area() float64
}

// *Circle and *Rectangle implicitly implement Shape
// because they provide an Area() float64 method

func main() {
	shapes := [...]Shape{
		&Circle{Radius: 2},
		&Rectangle{Width: 16, Height: 9},
	}

	for _, shape := range shapes {
		fmt.Printf("%#v -> %f\n", shape, shape.Area())
	}
}
