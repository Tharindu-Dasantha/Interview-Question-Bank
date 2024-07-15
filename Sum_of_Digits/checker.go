package main

import (
	"fmt"
)


func main() {
	var number int
	fmt.Print("Enter a number (enter any number large or small): ")
	fmt.Scan(&number)

	total := 0

	for {
		if number % 10 != 0 {
			total += number % 10
			number = number / 10
		} else {
			break
		}
	} 
	fmt.Println(total)
}