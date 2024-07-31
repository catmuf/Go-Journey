package main

import "fmt"

// custom print format library
func main() {
	const pi float64 = 3.16412732
	x := 5
	isbool := true

	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", isbool)
	fmt.Printf("%t \n", isbool)
	fmt.Printf("%d \n", x)
	// format to binary
	fmt.Printf("%b \n", x)
	// the character represented by the corresponding Unicode code point
	fmt.Printf("%c \n", x)
	// HEX
	fmt.Printf("%x \n", x)
}
