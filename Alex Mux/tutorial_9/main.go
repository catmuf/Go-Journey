package main

import "fmt"

func main() {
	// channels enables go routines to pass around information

	/*
		Holds data, slices or anything!
		Thread safe, this is good bc it avoids data to be erased when reading and writting from memory
		Listen for data, listen the data when added or removed from a channel.
		It has the possibility to block code execution one of these events happen
	*/
	// to create a channel use make function followed by the chan keyword and type of value the channel holds
	var c = make(chan int)
	// special syntax for channel to add value to 1
	// but we get a deadlock error and it gets block until something else reads from it -> waiting forever and unable to go the next step
	c <- 1
	// retrieve the value in a variable
	var i = <-c
	fmt.Println(i)
}
