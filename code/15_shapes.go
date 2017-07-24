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

func areaCircle(circ *Circle) float64 {
	return math.Pi * circ.Radius * circ.Radius
}

// You need distinct function names for areas
// because Go does not support overloading

func areaRectangle(rect *Rectangle) float64 {
	return rect.Width * rect.Height
}

func main() {
	c := Circle{Radius: 2}
	r := Rectangle{Width: 16, Height: 9}

	fmt.Printf("%#v -> %f\n", c, areaCircle(&c))
	fmt.Printf("%#v -> %f\n", r, areaRectangle(&r))
}
