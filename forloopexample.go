package main

import "fmt"

func main() {

	var fruits = []string{"Mango", "Banana", "Apple", "Orange"}

	for i := 0; i < len(fruits); i++ {

		fmt.Println(fruits[i])
	}

}
