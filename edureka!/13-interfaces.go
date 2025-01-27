package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{50, 60}
	circ := Circle{7}
	fmt.Println("Ärea of rectangle", getArea(rect))
	fmt.Println("Area of circle", getArea(circ))
}

/*
	Interface -> collections of method signatures.

defines and describes the exact methods that some other type must have.
*/
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
