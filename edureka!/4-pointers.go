package main

import "fmt"

func main() {
	x := 5
	// show value of x
	fmt.Println("Before", x)
	// every variable is a memory location, memory location has its own address using &
	fmt.Println(&x)

	changeValue(&x)
	fmt.Println("After", x)
}

func changeValue(x *int) {
	*x = 10
}
