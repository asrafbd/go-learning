package main

import "fmt"

func main() {

	mymap := map[int]string{
		1: "DevOPs",
		2: "Developer",
		3: "Cloud Engineer"}

	fmt.Println(mymap)
	fmt.Println(mymap[1])
}
