package main

import "fmt"

func main() {

	var fruits = []string{"Mango", "Banana", "Apple", "Orange"}
	for i := 0; i < len(fruits); i++ {

		fmt.Println(fruits[i])
	}

	// 2nd method for loop

	fmt.Println("2nd Method")
	for i := range fruits {
		fmt.Println(fruits[i])
	}

}
