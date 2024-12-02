package main

import (
	"fmt"
	"strings"
)

func main() {

	var builder strings.Builder

	for i := 0; i <= 10; i++ {

		_, err := builder.WriteString("Hello Golang\n")

		if err != nil {

			panic(err)
		}

	}

	str := builder.String()

	fmt.Println(str)

}
