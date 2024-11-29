package main

import "fmt"

func main() {

	var name string = "Hello"

	switch name {

	case "hello":

		fmt.Println("You are not correct")

	case "Hello":

		fmt.Println("You are correct")

	default:
		fmt.Println("Try again")
	}
}
