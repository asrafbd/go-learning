package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	path, err := filepath.Abs("test.txt")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(path)

	ext := filepath.Ext("test.txt")
	fmt.Println(ext)

	var fruits = []string{"Apple", "Banana", "Mango"}

	fmt.Println(fruits)
	fruits = append(fruits, "Coconut")
	fmt.Println(fruits)

	for i, j := range fruits {

		fmt.Println(i, j)
	}
}
