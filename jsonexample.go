package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {

	fmt.Println("Welcome to Json practice")
	EncodeJson()

}

func EncodeJson() {
	courselist := []course{

		{"React JS", 300, "courses.com", "abc1234", []string{"web-app"}},
		{"GoLang", 345, "courses.com", "vbdg435345", []string{"go-app"}},
	}

	//finalJson, err := json.Marshal(courselist := []course{)
	finalJson, err := json.MarshalIndent(courselist, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
