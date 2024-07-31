package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Understanding about strings, runes and bytes

func main() {
	// non-ASCII characters
	var myString = "résumé"

	fmt.Println(myString)

	// we can index the string like an array, but not really, pritns 144
	fmt.Println(myString[1])

	var indexed = myString[1]
	// check type of index
	fmt.Printf("Index %v, type of value %T\n", indexed, indexed)

	// iterate myString
	for i, v := range myString {
		fmt.Println(i, v)
	}
	// takes the numbers of bytes, not the characters!!!!!
	fmt.Printf("\nThe length of myString is %v\n", len(myString))

	// Solution use rune unicode code point, working with non-ASCII characters and multilingual text, that represent the character

	var myString2 = []rune("résumé") // rune is a type int32
	fmt.Println(myString2[1])
	var indexed2 = myString2[1]
	fmt.Printf("Index %v, type of value %T\n", indexed2, indexed2)
	// iterate myString
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	// count characters
	fmt.Println("Count:", utf8.RuneCountInString(myString))

	// declare a rune
	var myRune = 'a'
	fmt.Println("myRune", myRune)

	// String building (concatenate)
	var strSlice = []string{"h", "o", "m", "e", "p", "a", "g", "e"}

	var concatStr = ""

	for i := range strSlice {
		// creates a new string each time which is inefficient!!!
		concatStr += strSlice[i]
	}

	fmt.Println(concatStr)

	// use built-in go string package
	var strBuilder strings.Builder

	for i := range strSlice {
		// this appends using writestring method
		strBuilder.WriteString(strSlice[i])
	}
	// get the new string created from appended values
	var newConcatStr = strBuilder.String()
	fmt.Printf("%v", newConcatStr)
}
