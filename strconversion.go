package main

import (
	"fmt"
	"strconv"
)

func main() {

	bstr := "true"

	b, err := strconv.ParseBool(bstr)

	if err != nil {

		fmt.Println("Error found", err)
	}

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
