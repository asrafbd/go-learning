package main

import "fmt"

// start the function main ()
func main() {

	// declare and initialize the string
	str := "RADAR"
	fmt.Println("Golang program to check palindrome,\n Given Word =", str)

	// calling the function
	Palindrome(str)

	// print the result using fmt.Printf () function
	fmt.Printf("'%s' is palindrome\n", str)
}
func Palindrome(str string) bool {
	lastIdx := len(str) - 1
	// using for loop
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}
	return true
}
