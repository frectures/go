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

// Methods have an additional receiver argument
// and can be overloaded by their receiver

func (rect *Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func main() {
	c := Circle{Radius: 2}
	r := Rectangle{Width: 16, Height: 9}

	fmt.Printf("%#v -> %f\n", c, c.Area())
	fmt.Printf("%#v -> %f\n", r, r.Area())
}
