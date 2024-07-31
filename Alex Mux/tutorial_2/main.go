package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// declare a variable
	// name and type, meaning that stores a value type
	// int8 to 64, memory or bits to store number
	/* if int exceeds (overflow), warning from go.
	But adding it,  gives a strange output */
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	// flotat must be specified, either 32 or 64!!!
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// cannot mixed different type. Must do common types -> casting
	var result float32 = floatNum32 + float32(intNum32)

	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// single line
	var myString string = "Hello world"
	// format string using back quotes
	/* 	var myString string = `Hello
	world` */
	fmt.Println(myString)
	fmt.Println(len("test")) // number of bytes, go uses utf-8 encoding! not to be cnonfused number of characters

	//use import unicode/utf8 to count character, length of a string
	fmt.Println(utf8.RuneCountInString("Hello world"))

	var myRune rune = 'a' // ASCII table in decimal
	fmt.Println(myRune)

	// default value false
	var myBoolean bool
	fmt.Println(myBoolean)

	// inferred, meaning myVar is string type
	var myVar = "text"
	fmt.Println(myVar)

	// alternatively
	myVar2 := "text"
	fmt.Println(myVar2)

	// initialize multiple variables
	var var1, var2 int = 1, 2
	// var var1, var2 = 1, 2
	// var1, var2 := 1, 2

	fmt.Println(var1, var2)

	// Always initialize with a value!!!
	const myConst string = "const string"
	fmt.Println(myConst)
	// Cannot change its value since its a constant
	// myConst = "Hello"
	// Cannot declare constant without a value
	// const myConst string -> THIS IS WRONG!
}
