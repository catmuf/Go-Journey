package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("The value of a: %v\n", a)
	fmt.Printf("The type of a: %T\n\n", a)

	b := &a
	fmt.Printf("The value of b: %v\n", b)
	fmt.Printf("The type of b: %T\n\n", b)

	var c *int
	c = b
	fmt.Printf("The value of c: %v\n", c)
	fmt.Printf("The type of c: %T\n\n", c)

	*c++
	fmt.Println(a)
}
