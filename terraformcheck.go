package main

import (
	"fmt"
	"os/exec"
)

func main() {

	_, err := exec.LookPath("go")

	if err != nil {

		fmt.Println("Go is not Installed")
	} else {

		fmt.Println("Go is Installed")
	}
}
