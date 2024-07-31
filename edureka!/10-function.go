package main

import "fmt"

func main() {
	num := -5
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Iterative factorial function
func factorialIterative(n int) int {
	if n < 0 {
		return -1 // Factorial of a negative number is undefined
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
