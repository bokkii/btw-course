package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/mozart", mozartHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func mozartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Mozart!")
}
