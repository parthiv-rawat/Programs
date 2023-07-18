package main

import (
	"fmt"
)

func main() {
	// Prompt the user for input
	var input rune
	fmt.Print("Enter a character (a-z): ")
	fmt.Scanf("%c", &input)

	// Create channels for each step in the pipeline
	ch1 := make(chan rune)
	ch2 := make(chan rune)
	ch3 := make(chan rune)

	// Start the pipeline by passing the user input to the first function
	go firstFunction(input, ch1)

	// Each subsequent function in the pipeline reads from the previous channel and writes to the next channel
	go secondFunction(ch1, ch2)
	go thirdFunction(ch2, ch3)

	// Read the final result from the last channel in the pipeline
	result := <-ch3

	// Print the final result
	fmt.Printf("Final result: %c\n", result)
}

// First function in the pipeline
func firstFunction(input rune, ch chan<- rune) {
	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("First Function Output: %c\n", output)

	// Write the result to the channel
	ch <- output
}

// Second function in the pipeline
func secondFunction(in <-chan rune, out chan<- rune) {
	// Read from the input channel
	input := <-in

	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("Second Function Output: %c\n", output)

	// Write the result to the output channel
	out <- output
}

// Third function in the pipeline
func thirdFunction(in <-chan rune, out chan<- rune) {
	// Read from the input channel
	input := <-in

	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("Third Function Output: %c\n", output)

	// Write the result to the output channel
	out <- output
}
