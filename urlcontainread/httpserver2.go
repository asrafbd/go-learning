package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Golang Server is running at localhost:8080")
	http.HandleFunc("/", MyServer)
	http.ListenAndServe(":8080", nil)

}

func MyServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World !!")

}
