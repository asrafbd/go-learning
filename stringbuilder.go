package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	// Append strings to the builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world")

	// Convert the builder to a string
	result := builder.String()

	fmt.Println(result) // Output: Hello, world
}
