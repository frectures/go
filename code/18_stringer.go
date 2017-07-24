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

/*
package fmt

type Stringer interface {
	String() string
}
*/
type Shape interface {
	fmt.Stringer
	Area() float64
}

func (circ *Circle) String() string {
	return fmt.Sprintf("(%f)", circ.Radius)
}

func (rect *Rectangle) String() string {
	return fmt.Sprintf("[%f x %f]", rect.Width, rect.Height)
}

func main() {
	shapes := [...]Shape{
		&Circle{Radius: 2},
		&Rectangle{Width: 16, Height: 9},
	}

	for _, shape := range shapes {
		fmt.Printf("%s -> %f\n", shape.String(), shape.Area())
	}
}
