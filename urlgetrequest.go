package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Example of URL Get Request:")
	response, err := http.Get("https://example.com")

	if err != nil {

		fmt.Println("Http request is failed: %s", err)
	}

	data, _ := io.ReadAll(response.Body)

	fmt.Println(string(data))
	defer response.Body.Close()

}
