package main

import "fmt"

func main() {

	score := map[string]int{
		"student-1": 76,
		"student-2": 85,
		"student-3": 66,
		"student-4": 60}

	fmt.Println("Score of Student-3:", score["student-3"])
	fmt.Println(score)
	//delete(score, "student-4")
	for k, v := range score {

		fmt.Printf("score of %v is %v\n", k, v)
	}
	//fmt.Println(score)

}
