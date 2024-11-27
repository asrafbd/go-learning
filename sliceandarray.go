package main

import "fmt"

func main() {
	var numberlist = []int{1, 3, 0, 5, 2, 0, 4, 0}

	fmt.Println("Original list:", numberlist)

	var result []int
	zeroCount := 0

	// Loop through the numberlist
	for _, num := range numberlist {
		if num != 0 {
			result = append(result, num)
		} else {
			zeroCount++
		}
	}

	// Append zeros at the end
	for i := 0; i < zeroCount; i++ {
		result = append(result, 0)
	}

	fmt.Println("Updated list:", result)
}
