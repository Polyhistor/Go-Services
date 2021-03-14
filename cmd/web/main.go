package main

import (
	"fmt"
	"github.com/polyhistor/go-service/pkg/handlers"
	"net/http"
)

var portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %S", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
