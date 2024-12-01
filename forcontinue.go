package main

import "fmt"

func main() {

	var i int = 0

	fmt.Println("Even number are:")

	for i = 0; i <= 20; i++ {

		if i%2 == 0 {

			//continue

			fmt.Println(i)
		}

		//fmt.Println(i)

	}

	fmt.Println("Odd number are:")

	for i = 0; i <= 20; i++ {

		if i%2 != 0 {

			//continue
			fmt.Println(i)
		}

		//fmt.Println(i)

	}
}
