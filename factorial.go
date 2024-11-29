package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Print("Eneter your number: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Factorila value is:", factorial(n))
}

func factorial(n int) int {

	if n == 0 || n == 1 {

		return 1
	} else {
		result := 1
		for i := 2; i <= n; i++ {
			result = result * i

		}

		return result
	}
}
