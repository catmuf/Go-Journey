package main

import "fmt"

func main() {

	fmt.Println(addup(10, 20, 30, 40, 50))

}

func addup(args ...int) int {
	sum := 0
	fmt.Println(args)
	for _, value := range args {
		sum += value
	}
	return sum
}
