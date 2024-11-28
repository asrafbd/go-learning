package main

import (
	"fmt"
	"unicode"
)

func CodelandUsernameValidation(str string) string {
	fs := rune(str[0])
	ls := str[len(str)-1]
	if len(str) > 4 && len(str) < 25 && unicode.IsLetter(fs) && ls != '_' {
		return "true"
	} else {
		return "false"
	}
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	var username string
	fmt.Printf("UserName: ")

	fmt.Scanf("%s", &username)
	fmt.Println(CodelandUsernameValidation(username))

}
