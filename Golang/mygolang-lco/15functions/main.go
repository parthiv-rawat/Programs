package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in golang")

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(2, 5, 7, 5, 3)

	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", message)
}

func adder(valOne, valTwo int) (val int) {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namaste from golang")
}
