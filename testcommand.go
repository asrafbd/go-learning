package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "dir")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("An error occurred:", err) // Corrected the typo
		return
	}

	fmt.Println(string(output)) // Convert bytes to string
}
