package main

import (
	"fmt"
	"net/http"
)

const url = "https://example.com/"

func main() {

	fmt.Println("Checking example.com")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Response Status Code:", response.StatusCode)
}
