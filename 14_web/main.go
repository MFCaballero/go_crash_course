package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

type Course struct {
	Title  string `json:"title"`
	Videos int    `json:"videos"`
}

type Courses []Course

func handleJson(w http.ResponseWriter, r *http.Request) {
	courses := Courses{
		Course{"Go", 100},
		Course{"Python", 90},
		Course{"NodeJS", 95},
	}
	json.NewEncoder(w).Encode(courses)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/json", handleJson)
	fmt.Println("Server Starting...")

	http.ListenAndServe(":3000", nil)
}
