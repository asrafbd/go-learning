package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hello-a"
	output := convertToAlphabetPosition(input)
	fmt.Println(output)
}

func convertToAlphabetPosition(s string) string {
	var result strings.Builder

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			// Convert lowercase letters to their alphabet position
			position := int(r - 'a' + 1)                    // 97-97+1= 1
			result.WriteString(fmt.Sprintf("%d", position)) // return string value
		} else if r >= 'A' && r <= 'Z' {
			// Convert uppercase letters to their alphabet position
			position := int(r - 'A' + 1)
			result.WriteString(fmt.Sprintf("%d", position))
		} else {
			// Preserve non-alphabetic characters
			result.WriteRune(r)
		}
	}

	return result.String()
}
