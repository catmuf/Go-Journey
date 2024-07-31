package main

import "fmt"

func main() {

	// fixed value which can't be changed
	const pi float64 = 3.14159265359

	// Declare multiple variables
	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	// String contain characters between " or '
	var Name string = "I'm new to programming."

	// get string length
	fmt.Println(len(Name))

	// var a int = 2
	a := 10
	var b float32 = 4.32
	// declare two integers
	x, y := 14, 15

	fmt.Println("Value of a is", a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x, ",", y)
}
