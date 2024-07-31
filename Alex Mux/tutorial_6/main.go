package main

import "fmt"

// Struct defining type (similarly class in java)
// can hold mixed types in form of field
type gasEngine struct {
	brand     string
	mpg       uint8
	gallons   uint8
	ornerInfo owner
	int
}

type owner struct {
	name string
}
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// struct  or func -> func tied to a struct and have access to the struct instance itself
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

//takes only gasEngine
func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

//takes only electric
func canMakeIt2(e electricEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

// make a function to  make it general and take any engine type -> solution interfaces
/* Any type that implements these methods is said to satisfy the interface, which allows for flexible and reusable code.  */
type engine interface {
	milesLeft() uint8
}

func canMakeItGeneral(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

func main() {
	// struct literal syntax
	// var myEngine gasEngine = gasEngine{brand: "Toyota", mpg: 25, gallons: 15, ownerInfo:owner{name:"John"}}
	// ommit the fields
	var myEngine gasEngine = gasEngine{"Toyota", 25, 15, owner{"John"}, 10}
	// set values
	myEngine.brand = "Ford"
	fmt.Println(myEngine.brand, myEngine.gallons, myEngine.mpg, myEngine.ornerInfo.name, myEngine.int)

	// Anonymous structs, not defining a name type like with struct.
	// Define and initialized it in the same location
	// Not reusable! So create another one!
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	var myEngine3 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.gallons, myEngine2.mpg)
	fmt.Println(myEngine3.gallons, myEngine3.mpg)

	// instantiating a class and calling one of its methods
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 7)
	var myEV electricEngine = electricEngine{100, 10}
	fmt.Printf("Total miles left in tank: %v\n", myEV.milesLeft())
	canMakeIt2(myEV, 7)

	// call an interface function
	canMakeItGeneral(myEngine, 7)
	canMakeItGeneral(myEV, 7)
}
