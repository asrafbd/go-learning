package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &num)
	hrs := (num / 60)

	min := (num % 60)

	fmt.Println("Output:", hrs, ":", min)
}
