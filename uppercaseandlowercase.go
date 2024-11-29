package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string = "Hello"

	upperstr := strings.ToUpper(name)
	lowerstr := strings.ToLower(name)
	titlestr := strings.Title(name)

	fmt.Println("Uppercase version is: ", upperstr)
	fmt.Println("Lowercase version is: ", lowerstr)
	fmt.Println("Title version is: ", titlestr)
}
