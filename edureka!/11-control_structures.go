package main

/* func main() {
	// in a lifo (stack), it will execute until all surrounding fucntions are finished
	defer FirstRun()
	SecondRun()
}

func FirstRun() {
	fmt.Println("First")
}

func SecondRun() {
	fmt.Println("Second")
}
*/

/* func main() {
	// return error
	// fmt.Println("3 divided by 0 =", div(3, 0), "Impossible to divide zero hahaha")
	fmt.Println(div(36, 3), "Divisible")
	demPanic()
}

func div(num1, num2 int) int {
	// After everything is done
	// recover() = Regains the normal flow of execution afer a panic (like java exception) and capture the panic value  IN THE FIRST PLACE.
	// nil means no panic occured
	// Zeros

	defer func() {
		fmt.Println("Error message:", recover())
	}()

	solution := num1 / num2

	return solution
}

func demPanic() {
	// captures panic value after eveythong is done
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic, stops normal execution")
}
*/
