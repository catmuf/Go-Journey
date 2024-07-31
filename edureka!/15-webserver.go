package main

import (
	"fmt"
	"net/http"
)

/* create webserver simple and fast */
func main() {
	// like URL mapping in django
	// home page
	http.HandleFunc("/", home)
	http.HandleFunc("/page_1", page_1)
	// set port will listen on
	http.ListenAndServe(":8080", nil)
}

// create handler, takes a write and read response
func home(w http.ResponseWriter, r *http.Request) {
	// response request to page
	fmt.Fprint(w, "Welcome to main page!")
}

func page_1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
