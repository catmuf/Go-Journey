package main

import "fmt"

func main() {
	// declare two variables
	x, y := 14, 15

	// arithmetic
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)

	// var isbool bool = true
	isbool := true
	notbool, a := false, "sdasd"

	fmt.Println("isbool is", isbool, "and notbool is", notbool)

	// logical
	fmt.Println(isbool && notbool)
	fmt.Println(notbool || isbool || notbool)
	fmt.Println(!isbool, a)
}
