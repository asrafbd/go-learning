package main

import "fmt"

func main() {

	myword := "MADAM"

	lastIdx := len(myword) - 1

	for i := 0; i < lastIdx/2; i++ {

		if myword[i] != myword[lastIdx-i] {

			fmt.Println(false)
		}
	}
	fmt.Println(true)
	fmt.Println(lastIdx)

}
