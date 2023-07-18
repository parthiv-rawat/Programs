package main

import (
	"fmt"
)

func main() {
	// Prompt the user for input
	var input rune
	fmt.Print("Enter a character (a-z): ")
	fmt.Scanf("%c", &input)

	// Start the pipeline with the user input
	result := firstFunction(input)

	result = secondFunction(result)

	result = thirdFunction(result)

	// Print the final result
	fmt.Printf("Final result: %c\n", result)
}

// First function in the pipeline
func firstFunction(input rune) rune {
	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("First Function Output: %c\n", output)
	return output
}

// Second function in the pipeline
func secondFunction(input rune) rune {
	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("Second Function Output: %c\n", output)
	return output
}

// Third function in the pipeline
func thirdFunction(input rune) rune {
	// Perform any necessary operations on the input
	// In this example, we simply increment the input by 1
	output := input + 1
	fmt.Printf("Third Function Output: %c\n", output)
	return output
}
