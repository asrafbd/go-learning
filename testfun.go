package main

import (
	"fmt"
)

func add(x, y int) int {

	z := x + y
	return z
}

func main() {

	a := 10
	b := 20
	c := add(a, b)

	fmt.Println(c)
}
