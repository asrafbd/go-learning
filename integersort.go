package main

import "fmt"

func main() {

	fmt.Println("Integer sort list:")

	list := []int{1, 5, 4, 6, 3, 2}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {

			if list[j] > list[j+1] {

				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp

			}
		}
	}

	fmt.Println("New sorted list is:", list)
}
