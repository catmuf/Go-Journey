package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5)
	go process(c)
	// deadlock will appear again bc after finishing printing, from 0 to 4, main function will go back and wait at the top of the loop for another value!
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

}

func process(c chan int) {
	// close the channel to avoid deadlock before exiting the process, notifies any other process using the channel that is finsihed
	// main funciton will break out the loop and exit
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}

	fmt.Println("Exiting process")
}
