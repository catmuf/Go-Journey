package main

import (
	"fmt"
	"math/rand"
	"time"
)

var max_chicken_price float32 = 5
var max_tofu_price float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart", "costco", "wholefoods"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Println("Checking prices for tofu")
		if chickenPrice < max_tofu_price {
			tofuChannel <- website
			fmt.Println("Found a deal for tofu!")
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Println("Checking prices for chicken")
		if chickenPrice < max_chicken_price {
			chickenChannel <- website
			fmt.Println("Found a deal for chicken!")
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\n Found a deal chicken at %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\n Found a deal tofu at %s", website)
	}

}
