package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags"`
}

func main() {

	fmt.Println("Welcome to Json api")
	http.HandleFunc("/", Mycourse)
	http.ListenAndServe(":8080", nil)

}

var courselist = []course{

	{"React JS", 300, "courses.com", "abc1234", []string{"web-app"}},
	{"GoLang", 345, "courses.com", "vbdg435345", []string{"go-app"}},
}

func Mycourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courselist)

}
