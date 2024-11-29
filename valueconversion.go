package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	// Get user input
	fmt.Print("Enter something: ")
	fmt.Scanln(&input)

	// Determine the type of the input
	intValue, err := strconv.Atoi(input)
	if err == nil {
		fmt.Printf("You entered an integer: %d\n", intValue)
		return
	}

	floatValue, err := strconv.ParseFloat(input, 64)
	if err == nil {
		fmt.Printf("You entered a float: %.2f\n", floatValue)
		return
	}

	bvalue, err := strconv.ParseBool(input)
	if err == nil {

		fmt.Printf("You entered a boolean :%t\n", bvalue)
		return
	}

	// Use reflection to show it's a string
	fmt.Printf("You entered a string: %s\n", input)
}
