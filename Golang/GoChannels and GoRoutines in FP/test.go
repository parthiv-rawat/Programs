package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 500 and 1000 for the number of alphabets
	numAlphabets := rand.Intn(501) + 5

	// Generate random alphabets
	StrAlp := generateRandomAlphabets(numAlphabets)
	alphabets := []string{StrAlp}
	// Print the generated alphabets
	fmt.Println(alphabets)
}

// Function to generate random English alphabets
func generateRandomAlphabets(num int) string {
	rand.Seed(time.Now().UnixNano())

	var result string
	for i := 0; i < num; i++ {
		// Generate a random number between 0 and 25 for the alphabet index
		index := rand.Intn(26)
		// Convert the index to ASCII value of English alphabets (A: 65, B: 66, ..., Z: 90)
		ascii := index + 65
		// Convert the ASCII value to alphabet character
		letter := string(ascii)
		// Concatenate the letter to the result string
		result += " " + letter
	}

	return result
}
