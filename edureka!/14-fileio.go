package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// creates and can overwrite if file exists
	file, err := os.Create("sample.txt")
	// unique identifier or reference that the operating system assigns to a file when it is opened
	fmt.Println(file)
	//If capture error (panic)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hello guys!")

	file.Close()

	//display file as stream

	stream, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// file content to string
	s1 := string(stream)

	fmt.Println(s1)

}
