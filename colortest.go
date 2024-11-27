package main

import (
	"fmt"
	"image/color"
)

func main() {

	var (
		RED = color.new(color.FgRed).SprintFunc()
	)

	fmt.Println(RED("Red"))
}
