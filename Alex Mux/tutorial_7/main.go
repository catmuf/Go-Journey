package main

import "fmt"

func main() {
	// pointer are used to store memory location
	// doesnt have a value assigned to it and also memory address
	// to give a pointer an address
	// initialize a pointer
	var p *int32 = new(int32)
	// value as 0
	var i int32

	// to get the value of the pointer -> deferencing the pointer using *
	fmt.Printf("The value stored p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// to assign a value of a pointer
	*p = 10
	fmt.Printf("The value stored p points to is: %v\n", *p)
	// create a pointer from the address of another variable using &
	// get memory address of the variable, not its value
	// reference the int32 in memory
	a := &i
	fmt.Printf("The value stored a points to is: %v\n", *a)
	// change the value of a, it changes for var i
	*a = 1
	fmt.Printf("The value stored a points to is: %v\n", *a)
	fmt.Printf("The value of i is: %v\n", i)

	// regular variable changing value of i
	var k int32 = 2
	i = k
	fmt.Printf("The value of i is: %v\n", i)

	// slices cointain pointer to an underline array

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// pointer and functions
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Memory location of thing1: %p\n", &thing1)
	// this prevents from affecting the values of thing2 and thing1, but it double the memmory usage which is not good due to its creating a copy
	var result [5]float64 = square(thing1)
	fmt.Printf("\n The result of thing2 is %v:", result)
	var resultpointer [5]float64 = squarePointer(&thing1)
	fmt.Printf("\n Using squarePointer function the result of thing1 is %v:", resultpointer)
}

// this prevents from affecting the values of thing2 and thing1, but it double the memmory usage which is not good due to its creating a copy
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("Memory location of thing1: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// To solve the problem, use pointer
func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("Memory location of thing1: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
