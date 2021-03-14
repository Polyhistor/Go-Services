package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %S", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
