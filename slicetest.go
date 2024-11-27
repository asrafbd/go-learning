package main

import "fmt"

func main() {

	numberlist := []int{1, 4, 0, 2, 5, 0, 8}
	fmt.Println("Original list:", numberlist)
	zerocount := 0
	var result []int
	for _, num := range numberlist {

		if num != 0 {

			result = append(result, num)
		} else {

			zerocount++
		}

	}

	for i := 0; i < zerocount; i++ {

		result = append(result, 0)
	}

	fmt.Println("final result is :", result)
}
