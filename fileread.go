package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Read from a file:")

	data, err := os.ReadFile("./mytestfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File content is:", string(data))

	// Open file for writing
	file, err := os.OpenFile("./mytestfile.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write to the file
	_, err = file.WriteString("\nThis is new string")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File write success")
}
