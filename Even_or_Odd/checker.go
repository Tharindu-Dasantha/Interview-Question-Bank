// This program takes a input from a user and say wheather it's a even number or a odd number

package main


import (
	"fmt"
	"strconv"
)

func main() {
	// getting the user input
	var numberString string

	for {
		
		fmt.Println("Please enter a number: ")
		fmt.Scanln(&numberString)
		
		// checking if the number is a number
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Invalid input, Enter a number.")
		} else {
			// checking if the number is even or odd
			if number % 2 == 0 {
				fmt.Printf("%d is an even number\n", number)
			} else {
				fmt.Printf("%d is an odd number\n", number)
			}
			break
		}
	}
}