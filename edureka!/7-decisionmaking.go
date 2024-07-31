package main

import (
	"fmt"
)

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Can drink.")
	} else {
		fmt.Println("Can't drink")
	}

	switch {

	case age < 16:
		fmt.Println("Go out!")
	case age == 17:
		fmt.Println("Still can't drink")
	default:
		fmt.Println("Get a beer.")
	}
}
