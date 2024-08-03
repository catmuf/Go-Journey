package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEgine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEgine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	// instantiate two types of car
	var myGasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	// var myElectricCar = car[electricEgine]{
	// 	carMake:  "Tesla",
	// 	carModel: "Model 3",
	// 	engine: electricEgine{
	// 		kwh:   57.5,
	// 		mpkwh: 4.17,
	// 	},
	// }
	// %+v verb to display the field names and their values
	fmt.Printf("%+v\n", myGasCar)
}
