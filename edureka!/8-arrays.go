package main

import (
	"fmt"
)

func main() {
	// array of the same type and fixed size
	// var studentCount := [10]int{1,2,3,4,5,6,7,8,9,10}
	// studentCount := [10]int{1,2,3,4,5,6,7,8,9,10}
	// Alternative studentCount := [10]int{}
	var studentCount [10]int

	for i := 0; i < 9; i++ {
		studentCount[i] = i + 1
		fmt.Println(studentCount[i])
	}

	studentCount[9] = 10

	fmt.Printf("Recent added item: %d \n", studentCount[9])

	testNum := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, value := range testNum {
		fmt.Println(value)
		if value == 5 {
			println("Found", value, "at index", i)
			break
		}
	}

	testNumSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(len(testNum))
	// var sliceNum []int = testNum[1:4]
	sliceNum := testNumSlice[1:5]

	fmt.Println(sliceNum)

	sliceNum2 := testNumSlice[1:]
	fmt.Println(sliceNum2)

	sliceNum3 := make([]int, 5, 10)

	copy(sliceNum3, testNumSlice)
	fmt.Println(sliceNum3)
	sliceNum3 = append(sliceNum3, 6, 7, 8, 9, 10, 11)
	fmt.Println(cap(sliceNum3))
}
