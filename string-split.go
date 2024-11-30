package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Define the word list
	wordlist := "Hello World! Bangladesh"
	words := strings.Split(wordlist, " ")

	// Regular expression to match words containing only alphabet characters
	re := regexp.MustCompile(`^[a-zA-Z]+$`)

	var largestWord string
	for _, word := range words {
		// Find words matching the regular expression
		if re.MatchString(word) {
			// Check if the current word is larger than the largest word found so far
			if len(word) > len(largestWord) {
				largestWord = word
			}
		}
	}

	fmt.Println("Largest word is:", largestWord)
}
