package main

import "fmt"

func main() {
	// map like a dict, relies key value structure. used for hash maps
	// map(key) value. keys are unique!!!
	// create a map
	Student := make(map[string]int)

	// add element k : v
	Student["Aryya"] = 42
	//print value of Aryya
	fmt.Println(Student["Aryya"])
	// length of map
	fmt.Println("Length of Student map is", len(Student))

	Student["Pepe"] = 20
	fmt.Println(len(Student))
	// prints entire map
	fmt.Println(Student)
	// delete item from key
	delete(Student, "Pepe")
	fmt.Println(len(Student))

	// nested maps, dual mapping
	superHero := map[string]map[string]string{
		"Superman": map[string]string{
			"Name": "Clark Kent",
			"City": "Metropolis",
		},
		"Batman": map[string]string{
			"Name": "Bruce Wayne",
			"City": "Gotham City",
		},
	}
	// output data where key matches
	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["Name"], temp["City"])
	}
	// prints k and v
	for hero, details := range superHero {
		fmt.Println(hero)
		for name, city := range details {
			fmt.Println(name, city)
		}
	}

	for hero, details := range superHero {
		fmt.Println(hero)
		// print values
		for _, value := range details {
			fmt.Println(value)
		}
	}
	// create a map
	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	// for-range loop to iterate through each key-value of map
	for number, squared := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}
}
