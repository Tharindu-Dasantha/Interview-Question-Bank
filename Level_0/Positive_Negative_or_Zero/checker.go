package main

import ( 
	"fmt"
)


func main() {
	// getting the user input numner
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number) 

	if number == 0 {
		fmt.Println("The number is a zero")
	}	else if number < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is positive")
	}
}