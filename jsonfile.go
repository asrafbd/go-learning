package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {

	fmt.Println("Welcome to Json practice")
	user := User{

		Name: "asraful",
		Role: "DevOps",
	}
	jsonbytes, _ := json.Marshal(user)

	file, _ := os.Create("myjson.json")

	defer file.Close()

	bytes, _ := file.Write(jsonbytes)

	fmt.Println(bytes)

}
