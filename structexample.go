package main

import "fmt"

func main() {

	type User struct {
		Name  string
		Email string
		Age   int
		score float32
	}

	asraful := User{"Asraful", "asrafbd@gmail.com", 40, 67.8}

	fmt.Printf("%v\n", asraful)
	fmt.Printf("%+v", asraful)
}
