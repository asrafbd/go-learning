package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "hello-Hello-6767"
	var mystr strings.Builder

	for _, j := range str {

		if j >= 'a' && j <= 'z' {

			ps := int(j - 'a' + 1)

			mystr.WriteString(fmt.Sprintf("%d", ps))

		} else if j >= 'A' && j <= 'Z' {

			ps := int(j - 'A' + 1)

			mystr.WriteString(fmt.Sprintf("%d", ps))
		} else {

			mystr.WriteRune(j)
		}
	}

	fmt.Print(mystr.String())

}
