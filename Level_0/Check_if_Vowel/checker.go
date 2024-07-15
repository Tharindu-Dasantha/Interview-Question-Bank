package main

import (
	"fmt"
)

func main() {
	vowels := [5]byte{'a', 'e', 'i', 'o', 'u'}

	// getting the word form the user
	var userInput byte
	fmt.Print("Enter a letter: ")
	fmt.Scan(&userInput)

	// checking the letter
	for _, letter := range vowels {
		if letter == userInput {
			fmt.Println("The letter is a vowel.")
			return
		} else {
			fmt.Println("The letter is not a vowel.")
			return
		}
	}
}