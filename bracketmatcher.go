package main

import (
	"fmt"
	"strings"
)

func BracketMatcher(str string) string {
	// code goes here
	if strings.Count(str, "(") != strings.Count(str, ")") {
		return "0"
	}
	return "1"
}
func main() {

	message := "((hello(world))"
	fmt.Println(BracketMatcher(message))

}
