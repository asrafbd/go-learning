package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")
	content := "the is a context of new file"
	file, err := os.Create("./mytestfile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("file is created:", length)
	defer file.Close()
	readfile("./mytestfile.txt")
}

// read file

func readfile(filename string) {

	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("the file context is: \n", string(databyte))
}
