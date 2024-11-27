package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com/"

func main() {

	fmt.Println("Example web request")

	response, err := http.Get(url)

	if err != nil {

		panic(err)
	}
	fmt.Printf("Response of a type: %T\n", response)
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println("web contain is:\n", string(databyte))

}
