package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path, err := filepath.Abs("./main.go")

	if err != nil {

		fmt.Println("Not found:", err)
	}

	fmt.Println(path)

	ext := filepath.Ext(path)

	fmt.Println("File extension is :", ext)

}
