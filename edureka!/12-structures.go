package main

import "fmt"

func main() {
	// structures allows to combine data, items of different kinds. To keep records

	//create rectangle type data
	//rect1 = Rectangle{height: 10, width: 5}

	// if we know the order of the struct rectangle
	rect1 := Rectangle{10, 5}
	fmt.Println(rect1.height)
	fmt.Println(rect1.width)

	fmt.Println("Area:", rect1.area())
}

// create struct and defines the data type, which is destruct
type Rectangle struct {
	height float64
	width  float64
}

// pass in rect type of Rectangle, has a func area and return float64
func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
