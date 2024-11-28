package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Example of file creation")

	file, err := os.Create("file1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("hello")
	fmt.Println("File creation done")

}
