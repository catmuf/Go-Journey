package main

import (
	"fmt"
)

func main() {
	/* Array, slices, maps, loops */
	/*	Array, fixed length (cannot be change after is initialized),
		same type, indexable and contiguous in memory
	*/
	// 32bits / 8 btis = 4 bytes. 3 elements which allocates 12 bytes
	// var intArr [3]int32 // [0,0,0]

	// initialize array
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// var intArr = [3]int32{1, 2, 3}
	// inferred by the compiler. still an array with fixed size
	// intArr := [...]int32{1, 2, 3}
	intArr := [3]int32{1, 2, 3}

	// change index value
	intArr[1] = 123
	// access array
	fmt.Println(intArr[0])
	// access element 1 and 2
	fmt.Println(intArr[1:3])

	// print memory location using apersand pointer
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	/*	slice is a dynamic array, grow or shrink in size dynamically
		and point the data stored in an array. Omitting the size, we have a slice
	*/
	var intSlice = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v, memory %p", len(intSlice), cap(intSlice), &intSlice)
	// add value
	intSlice = append(intSlice, 7)
	// array holds 3 values. when append, it checks if it has enough room for 4 values
	// in this case it does not. Creates a new array with new capacity and copied the array, stored in different memory location
	fmt.Printf("\nThe length is %v with capacity %v, memory %p", len(intSlice), cap(intSlice), &intSlice)

	//Append mutliple values to the slice, using spread operator ...
	var intSlice2 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\nThe length is %v with capacity %v, memory %p", len(intSlice), cap(intSlice), &intSlice)

	// Another way to make a slice, using make
	var intSlice3 = make([]int32, 3, 8)
	fmt.Printf("\nThe length is %v with capacity %v, memory %p", len(intSlice3), cap(intSlice3), &intSlice3)

	// map is a dictionary formed by k and v pairs
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Printf("\n%v\n", myMap)

	// initialize the map with values
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 40}
	fmt.Println(myMap2["Adam"])
	// key noexist -> return 0 uint8
	fmt.Println(myMap2["Josh"])

	// optional second value that checks is the key is in the map, ok is a bool
	var age, ok = myMap2["kim"]
	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}
	// number of k, v
	fmt.Println("Length of myMap2:", len(myMap2))

	// dete by reference, no return value is given
	delete(myMap2, "Adam")

	// iterate the map using for loop
	for k, v := range myMap2 {
		fmt.Printf("Index: %v, Value:%v\n", k, v)
	}

	// while loop
	var i, b int = 0, 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	// optionally, omit the condition in for loop
	for {
		if b >= 10 {
			break
		}
		fmt.Println(b)
		b += 1
	}

	// alternative for loop
	for e := 0; e < 10; e++ {
		fmt.Println(e)
	}

}
