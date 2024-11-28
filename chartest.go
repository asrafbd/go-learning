package main

import "fmt"

func main() {
	var r rune = 'A'
	fmt.Println(r)        // Prints the Unicode code point of 'A'
	fmt.Printf("%c\n", r) // Prints the character 'A'
}
