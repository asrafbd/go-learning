package main

import (
	"errors"
	"fmt"
)

func main() {

	message := "Hello"
	myerror := errors.New("Wrong")
	if message == "Hello" {

		fmt.Println("You are correct")
	} else {

		fmt.Println(myerror)
	}
}
