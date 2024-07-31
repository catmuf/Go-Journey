package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0

	var result, remainder, err = intDivision(numerator, denominator)
	// error encountered
	// if statement
	/* 	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	} */

	// switch statement
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the division is %v", result)
	default:
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	}

	// conditional switch statement
	switch remainder {
	case 0:
		fmt.Println("The division was exact.")
	case 1, 2:
		fmt.Println("The division was close.")
	default:
		fmt.Println("The division was not close!")
	}
}

// parameters of the func, name type. therefore, the parameter is enforced!
// inside the func is the logic of the func
func printMe(printvalue string) {
	fmt.Println(printvalue)
}

// function returns a type int
// it can algo reutrn multiple values at the same time
func intDivision(numerator int, denominator int) (int, int, error) {
	// default value is nil. built-in func in go
	var err error
	if denominator == 0 {
		err = errors.New("cannot devide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
