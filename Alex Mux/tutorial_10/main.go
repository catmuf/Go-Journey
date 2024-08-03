package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sunSlice[float32](float32Slice))

}

// Generic function that allows int | float32 | float64
func sunSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// allow all variable types - any type, this func doesnt allow bool to do operations
/* func sunSlice[T any](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
} */

// but can be used any for checking and other cases.
// check if the slice is empty, which returns a bool
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
