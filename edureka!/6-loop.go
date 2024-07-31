package main

import (
	"fmt"
)

/* func main() {
	start := time.Now()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	elapse := time.Since(start)
	fmt.Printf("Took %s", elapse)
} */

/* func main() {
	start := time.Now()
	// while loop but using for
	i := 0
	for i < 1000000000 {
		i++
	}
	fmt.Println(i)
	elapse := time.Since(start)
	fmt.Printf("Took %s", elapse)
} */

func main() {
	// nested loop
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
