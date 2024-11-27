package main

import (
	"fmt"
)

func calc(x, y int) (int, int) {

	out1 := x + y

	out2 := x - y

	return out1, out2
}

func main() {

	a := 40
	b := 20
	result1, result2 := calc(a, b)

	fmt.Println(result1, result2)
}
