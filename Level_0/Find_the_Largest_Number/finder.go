package main

import (
	"fmt"
)

func main() {
	// getting three user inputs
	var num1, num2, num3 int

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter the third number: ")
	fmt.Scanln(&num3)

	var largest int
	largest = 0

	if num1 > largest {
		largest = num1
	}
	if num2 > largest {
		largest = num2
	}
	if num3 > largest {
		largest = num3
	}

	fmt.Printf("The largest number is : %d", largest)
	fmt.Println()

}
